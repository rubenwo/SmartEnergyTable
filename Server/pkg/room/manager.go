package room

import (
	"fmt"
	"github.com/google/uuid"
	v1 "github.com/rubenwo/SmartEnergyTable/Server/pkg/api/v1"
	"github.com/rubenwo/SmartEnergyTable/Server/pkg/database"
	"log"
	"sync"
)

type Manager struct {
	db    database.Database
	rooms map[string]*Room
}

// NewManager creates a manager object and instantiates a connection to the backend database.
// The function returns nil if an error occurred with the database creation.
func NewManager(db database.Database) *Manager {
	return &Manager{
		db:    db,
		rooms: make(map[string]*Room),
	}
}

// CreateRoom creates a new uuid and creates a room. It then returns that ID.
func (m *Manager) CreateRoom() (id string) {
	id = uuid.New().String()

	m.rooms[id] = &Room{
		Lock:               sync.Mutex{},
		RoomID:             id,
		tokens:             map[string]*v1.Token{},
		userPosition:       v1.Vector3_Protocol{},
		currentScene:       1,
		master:             "",
		clients:            map[string]chan Patch{},
		clientsNeedHistory: map[string]bool{},
	}
	return id
}

// JoinRoom uses the id parameter to get the Room. User is the userID and callback is a channel that receives Patch
// structs.
// These Data structs are used for the updates to the client.
func (m *Manager) JoinRoom(id, user string, callback chan Patch) error {
	if callback == nil {
		return fmt.Errorf("callback channel can't be nil")
	}
	var room *Room
	var ok bool
	room, ok = m.rooms[id]
	if !ok {
		// Room does not exist in memory
		raw, err := m.db.Get(id)
		if err != nil {
			log.Println(err)
			return fmt.Errorf("room with id: %s does not exist", id)
		}
		room, ok = raw.(*Room)
		if !ok {
			log.Println("Conversion from interface to Room didn't work")
			return fmt.Errorf("internal error with the casting of interface{} to Room from database")
		}
		// TODO: Implement proper load-balancing
		room.master = user // As the room needs to be loaded from the database, this means the master might not be the
		// same
	}
	room.Lock.Lock()
	defer func() {
		// We notify everyone in the room because we made a change. We need to do this concurrently as the client that
		// calls the JoinRoom function is not actually listening yet. Which would result in a deadlock.
		go func() { room.Notify() }()
	}()
	defer room.Lock.Unlock()
	// A newly created room does not contain a master yet. So the first person to join the room is automatically the
	// master. This is in ~100% of the cases the creator of the room as these functions are called directly after each
	// other.
	if room.master == "" {
		room.master = user
		log.Println("JoinRoom() => Master joined")
	} else {
		log.Println("JoinRoom() => Client join")
	}
	room.clients[user] = callback
	room.clientsNeedHistory[user] = true

	return nil
}

// SaveRoom persists the room with the id (string) to a datastore backend. If everything went well error == nil.
// This function can be called by anyone as there won't be any changes applied to the room.
func (m *Manager) SaveRoom(id string) error {
	room, ok := m.rooms[id]
	if !ok {
		return fmt.Errorf("room with id: %s does not exist", id)
	}
	if err := m.db.Set(id, room); err != nil {
		return fmt.Errorf("error saving room with id: %s, with error: %w", id, err)
	}
	raw, err := m.db.Get(id)
	if err != nil {
		log.Println(err)
	}
	fmt.Println(raw)
	r, ok := raw.(*Room)
	if !ok {
		log.Println("Conversion failed")
	}
	fmt.Println(r)
	return nil
}

// AddToken checks if the user (param) is the master of the room (id). Only the master may alter the room.
// After checking the function adds the new token object to the room and notifies the clients.
// Finally notifying the room as there has been a change.
func (m *Manager) AddToken(id, user string, object *v1.Token) error {
	room, ok := m.rooms[id]
	if !ok {
		return fmt.Errorf("room with id: %s does not exist", id)
	}
	room.Lock.Lock()
	defer room.Notify()
	defer room.Lock.Unlock()
	if room.master != user {
		return fmt.Errorf("user: %s is not the master of room: %s", user, id)
	}
	object.ObjectId = uuid.New().String() // Generate a uuid for the new object.
	// Set the token
	room.tokens[object.ObjectId] = object
	room.changes = append(room.changes, Diff{
		Action: ADD,
		Token:  object,
	})
	return nil
}

// RemoveToken removes the token from the room after checking that the user is the master.
// Finally notify the room as there has been a change.
func (m *Manager) RemoveToken(id, user string, object *v1.Token) error {
	room, ok := m.rooms[id]
	if !ok {
		return fmt.Errorf("room with id: %s does not exist", id)
	}
	room.Lock.Lock()
	defer room.Notify()
	defer room.Lock.Unlock()
	if room.master != user {
		return fmt.Errorf("user: %s is not the master of room: %s", user, id)
	}
	delete(room.tokens, object.ObjectId)
	room.changes = append(room.changes, Diff{
		Action: DELETE,
		Token:  object,
	})
	return nil
}

// MoveToken moves the token after checking the user is master and the token exists.
// Finally notifying all the clients as there has been a change.
func (m *Manager) MoveToken(id, user string, object *v1.Token) error {
	room, ok := m.rooms[id]
	if !ok {
		return fmt.Errorf("room with id: %s does not exist", id)
	}
	room.Lock.Lock()
	defer room.Notify()
	defer room.Lock.Unlock()

	if room.master != user {
		return fmt.Errorf("user: %s is not the master of room: %s", user, id)
	}
	room.tokens[object.ObjectId] = object
	room.changes = append(room.changes, Diff{
		Action: MOVE,
		Token:  object,
	})
	return nil
}

// ClearRoom creates a list of diffs to clear every token from the room
// id is the room ID and user is the user ID. if the room is not found, an error is returned.
// If the user is not the master of the room that will also result in an error.
func (m *Manager) ClearRoom(id, user string) error {
	room, ok := m.rooms[id]
	if !ok {
		return fmt.Errorf("room with id: %s does not exist", id)
	}
	room.Lock.Lock()
	defer room.Notify()
	defer room.Lock.Unlock()

	if room.master != user {
		return fmt.Errorf("user: %s is not the master of room: %s", user, id)
	}

	for _, token := range room.tokens {
		room.changes = append(room.changes, Diff{
			Action: DELETE,
			Token:  token,
		})
	}
	room.tokens = map[string]*v1.Token{}
	return nil
}

// ChangeScene changes the scene from the room after checking that the user is the master.
// Finally notifying all the clients in the room as there has been a change.
func (m *Manager) ChangeScene(id, user string, sceneID int) error {
	room, ok := m.rooms[id]
	if !ok {
		return fmt.Errorf("room with id: %s does not exist", id)
	}
	room.Lock.Lock()
	defer room.Notify()
	defer room.Lock.Unlock()

	if room.master != user {
		return fmt.Errorf("user: %s is not the master of room: %s", user, id)
	}

	room.currentScene = sceneID

	return nil
}

// RemoveClient closes the callback channel (we can do this as the manager is the sender).
// If the user is the master, all the clients are closed and deleted. Then the room is deleted from memory.
// If the user is a client, only this client is deleted from memory after closing its channel.
func (m *Manager) RemoveClient(id, user string) error {
	room, ok := m.rooms[id]
	if !ok {
		return fmt.Errorf("room with id: %s does not exist", id)
	}
	if room.master == user {
		log.Println("RemoveClient() => Master left, all client are closing.")
		for user, client := range room.clients {
			close(client)
			delete(room.clients, user)
		}
		delete(m.rooms, id)
	} else {
		log.Println("RemoveClient() => Client left.")
		close(room.clients[user])
		delete(room.clients, user)
		room.Notify() // Notify only when a single client has left.
	}
	return nil
}

// ChangeMaster does exactly what the name suggests. It changes the old master to the new master in a given room.
func (m *Manager) ChangeMaster(id, master, newMaster string) error {
	room, ok := m.rooms[id]
	if !ok {
		return fmt.Errorf("room with id: %s does not exist", id)
	}

	room.Lock.Lock()
	defer room.Notify()
	defer room.Lock.Unlock()

	if room.master != master {
		return fmt.Errorf("you don't have the permissions to change the master")
	}
	room.master = newMaster

	return nil
}

// ChangeUserPosition changes the position of the users after check that the caller is the master of the room
// and the room exists
func (m *Manager) MoveUsers(id, master string, position v1.Vector3_Protocol) error {
	room, ok := m.rooms[id]
	if !ok {
		return fmt.Errorf("room with id: %s does not exist", id)
	}
	room.Lock.Lock()
	defer room.Notify()
	defer room.Lock.Unlock()
	if room.master != master {
		return fmt.Errorf("you don't have the permissions to change the master")
	}
	room.userPosition = position

	return nil
}

// RoomIDs returns an array of strings containing all the room IDs (keys) from the rooms map.
func (m *Manager) RoomIDs() []string {
	var ids []string
	for key := range m.rooms {
		ids = append(ids, key)
	}
	return ids
}
