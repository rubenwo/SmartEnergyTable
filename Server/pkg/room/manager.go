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

func (m *Manager) CreateRoom() (id string) {
	id = uuid.New().String()
	m.rooms[id] = &Room{Data: struct {
		ID      string
		SceneID int
		Objects map[string]SceneObject
	}{ID: id, SceneID: 1, Objects: make(map[string]SceneObject)}, master: "", clients: make(map[string]chan Data, 1)}
	log.Println("Created room:", id)
	return id
}

func (m *Manager) JoinRoom(id string, user string, callback chan Data) error {
	if callback == nil {
		return fmt.Errorf("callback channel can't be nil")
	}
	room, ok := m.rooms[id]
	if !ok {
		return fmt.Errorf("room with id: %s does not exist", id)
	}
	if room.master == "" {
		room.master = user
		log.Println("Master joined")
	} else {
		log.Println("Client joined")
	}
	room.clients[user] = callback

	go func(r *Room) { r.Notify() }(room)

	return nil
}

func (m *Manager) Room(id string) *Room {
	return m.rooms[id]
}

func (m *Manager) AddToken(id string, user string, object *v1.Token) error {
	room, ok := m.rooms[id]
	if !ok {
		return fmt.Errorf("room with id: %s does not exist", id)
	}
	if room.master != user {
		return fmt.Errorf("user: %s is not the master of room: %s", user, id)
	}
	object.ObjectId = uuid.New().String()
	room.Data.Objects[object.ObjectId] = SceneObject{
		Index: object.ObjectIndex,
		Position: Vector3{
			X: object.Position.X,
			Y: object.Position.Y,
			Z: object.Position.Z,
		},
	}
	log.Println("Added token")
	room.Notify()
	return nil
}

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

func (m *Manager) UpdateRoom(id string, sceneId int, objects map[string]SceneObject) error {
	room, ok := m.rooms[id]
	if !ok {
		return fmt.Errorf("room with id: %s does not exist", id)
	}
	room.Data.SceneID = sceneId
	if objects != nil {
		room.Data.Objects = objects
	}
	room.Notify()
	return nil
}

func (m *Manager) RemoveClient(id string, user string) error {
	room, ok := m.rooms[id]
	if !ok {
		return fmt.Errorf("room with id: %s does not exist", id)
	}
	if room.master == user {
		log.Println("Master left")
		for user, client := range room.clients {
			close(client)
			delete(room.clients, user)
		}
		delete(m.rooms, id)
	} else {
		log.Println("User left")
		close(room.clients[user])
		delete(room.clients, user)
	}
	return nil
}

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

func (m *Manager) RoomIDs() []string {
	var ids []string
	for key, _ := range m.rooms {
		ids = append(ids, key)
	}
	return ids
}
