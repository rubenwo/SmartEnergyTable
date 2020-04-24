package room

import (
	"fmt"
	"github.com/google/uuid"
	v1 "github.com/rubenwo/SmartEnergyTable/Server/pkg/api/v1"
	"github.com/rubenwo/SmartEnergyTable/Server/pkg/database"
	"log"
)

type Manager struct {
	db    database.Database
	rooms map[string]*Room
}

//NewManager creates a manager object and instantiates a connection to the backend database.
//The function returns nil if an error occurred with the database creation.
func NewManager() (*Manager, error) {
	db, err := database.Factory("redis")
	if err != nil {
		return nil, err
	}
	return &Manager{
		db:    db,
		rooms: make(map[string]*Room),
	}, nil
}

//CreateRoom creates a new uuid and creates a room. It then returns that ID.
func (m *Manager) CreateRoom() (id string) {
	id = uuid.New().String()
	m.rooms[id] = &Room{Data: struct {
		ID       string
		SceneID  int
		Objects  map[string]SceneObject
		IsMaster bool
	}{ID: id, SceneID: 1, Objects: make(map[string]SceneObject), IsMaster: false}, master: "", clients: make(map[string]chan Data, 1)}
	return id
}

//JoinRoom uses the id parameter to get the Room. User is the userID and callback is a channel that receives Data structs.
//These Data structs are used for the updates to the client.
func (m *Manager) JoinRoom(id string, user string, callback chan Data) error {
	if callback == nil {
		return fmt.Errorf("callback channel can't be nil")
	}
	room, ok := m.rooms[id]
	if !ok {
		return fmt.Errorf("room with id: %s does not exist", id)
	}
	//A newly created room does not contain a master yet. So the first person to join the room is automatically the master.
	//This is in ~100% of the cases the creator of the room as these functions are called directly after each other.
	if room.master == "" {
		room.master = user
		log.Println("JoinRoom() => Master joined")
	} else {
		log.Println("JoinRoom() => Client join")
	}
	room.clients[user] = callback

	//We notify everyone in the room because we made a change. We need to do this concurrently as the client that calls
	//the JoinRoom function is not actually listening yet. Which would result in a deadlock.
	go func(r *Room) { r.Notify() }(room)

	return nil
}

//Room returns a pointer to a room. This means the values inside the room can be changed.
func (m *Manager) Room(id string) (*Room, error) {
	room, ok := m.rooms[id]
	if !ok {
		return nil, fmt.Errorf("room with id: %s does not exist", id)
	}
	return room, nil
}

//AddToken checks if the user (param) is the master of the room (id). Only the master may alter the room.
//After checking the function adds the new token object to the room and notifies the clients.
//Finally notifying the room as there has been a change.
func (m *Manager) AddToken(id string, user string, object *v1.Token) error {
	room, ok := m.rooms[id]
	if !ok {
		return fmt.Errorf("room with id: %s does not exist", id)
	}
	if room.master != user {
		return fmt.Errorf("user: %s is not the master of room: %s", user, id)
	}
	object.ObjectId = uuid.New().String() //Generate a uuid for the new object.
	//Change the token data to a new data structure.
	room.Data.Objects[object.ObjectId] = SceneObject{
		Index: object.ObjectIndex,
		Position: Vector3{
			X: object.Position.X,
			Y: object.Position.Y,
			Z: object.Position.Z,
		},
	}

	room.Notify()
	return nil
}

//RemoveToken removes the token from the room after checking that the user is the master.
//Finally notify the room as there has been a change.
func (m *Manager) RemoveToken(id string, user string, object *v1.Token) error {
	room, ok := m.rooms[id]
	if !ok {
		return fmt.Errorf("room with id: %s does not exist", id)
	}
	if room.master != user {
		return fmt.Errorf("user: %s is not the master of room: %s", user, id)
	}
	delete(room.Data.Objects, object.ObjectId)
	room.Notify()
	return nil
}

//MoveToken moves the token after checking the user is master and the token exists.
//Finally notifying all the clients as there has been a change.
func (m *Manager) MoveToken(id string, user string, object *v1.Token) error {
	room, ok := m.rooms[id]
	if !ok {
		return fmt.Errorf("room with id: %s does not exist", id)
	}
	if room.master != user {
		return fmt.Errorf("user: %s is not the master of room: %s", user, id)
	}

	obj, ok := room.Data.Objects[object.ObjectId]
	if !ok {
		return fmt.Errorf("object with id: %s doesn't exist", object.ObjectId)
	}
	obj.Position = Vector3{
		X: object.Position.X,
		Y: object.Position.Y,
		Z: object.Position.Z,
	}

	room.Notify()
	return nil
}

//ChangeScene changes the scene from the room after checking that the user is the master.
//Finally notifying all the clients in the room as there has been a change.
func (m *Manager) ChangeScene(id string, user string, sceneId int, objects map[string]SceneObject) error {
	room, ok := m.rooms[id]
	if !ok {
		return fmt.Errorf("room with id: %s does not exist", id)
	}
	if room.master != user {
		return fmt.Errorf("user: %s is not the master of room: %s", user, id)
	}

	room.Data.SceneID = sceneId
	if objects != nil {
		room.Data.Objects = objects
	}
	room.Notify()
	return nil
}

//RemoveClient closes the callback channel (we can do this as the manager is the sender).
//If the user is the master, all the clients are closed and deleted. Then the room is deleted from memory.
//If the user is a client, only this client is deleted from memory after closing its channel.
func (m *Manager) RemoveClient(id string, user string) error {
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
		room.Notify() //Notify only when a single client has left.
	}
	return nil
}

//ChangeMaster does exactly what the name suggests. It changes the old master to the new master in a given room.
func (m *Manager) ChangeMaster(id string, master string, newMaster string) error {
	room, ok := m.rooms[id]
	if !ok {
		return fmt.Errorf("room with id: %s does not exist", id)
	}

	if room.master == master {
		room.master = newMaster
		return nil
	}

	return fmt.Errorf("you don't have the permissions to change the master")
}

//RoomIDs returns an array of strings containing all the room IDs (keys) from the rooms map.
func (m *Manager) RoomIDs() []string {
	var ids []string
	for key, _ := range m.rooms {
		ids = append(ids, key)
	}
	return ids
}
