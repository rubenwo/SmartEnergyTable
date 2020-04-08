package room

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/rubenwo/SmartEnergyTable/Server/pkg/database"
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
	}{ID: id, SceneID: 0, Objects: nil}, master: nil, clients: make([]chan Data, 1)}
	return id
}

func (m *Manager) JoinRoom(id string, callback chan Data) error {
	if callback == nil {
		return fmt.Errorf("callback channel can't be nil")
	}
	if room, ok := m.rooms[id]; ok {
		if room.master == nil {
			room.master = callback
		} else {
			room.clients = append(room.clients, callback)
		}
	}
}

func (m *Manager) Room(id string) *Room {
	return m.rooms[id]
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
	for _, r := range room.clients {
		r <- room.Data
	}
	room.master <- room.Data
	return nil
}
