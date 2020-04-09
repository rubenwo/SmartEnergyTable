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
		Objects []SceneObject
	}{ID: id, SceneID: 0, Objects: make([]SceneObject, 0)}, master: "", clients: make(map[string]chan Data, 1)}
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
		log.Println("Client is master joined")
	}
	room.clients[user] = callback
	log.Println("Client joined")

	go func(r *Room) { r.Notify() }(room)

	return nil
}

func (m *Manager) Room(id string) *Room {
	return m.rooms[id]
}

func (m *Manager) AddGameObject(id string, user string, object *v1.GameObject) error {
	room, ok := m.rooms[id]
	if !ok {
		return fmt.Errorf("room with id: %s does not exist", id)
	}
	room.Data.Objects = append(room.Data.Objects, SceneObject{
		Name: object.Name,
		PosX: object.PosX,
		PosY: object.PosY,
		PosZ: object.PosZ,
	})
	room.Notify()
	return nil
}

func (m *Manager) UpdateRoom(id string, sceneId int, objects []SceneObject) error {
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
	close(room.clients[user])
	delete(room.clients, user)

	return nil
}
