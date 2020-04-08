package room

import (
	"fmt"
	"github.com/google/uuid"
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
	m.rooms[id] = &Room{id: id}
	return id
}

func (m *Manager) JoinRoom(id string) {
	log.Println(m.db.Set("test", "test123"))
	log.Println(m.db.Get("test"))
}

func (m *Manager) Room(id string) *Room {
	return m.rooms[id]
}

func (m *Manager) UpdateRoom(id string, sceneId int, objects []SceneObject) error {
	room, ok := m.rooms[id]
	if !ok {
		return fmt.Errorf("room with id: %s does not exist", id)
	}
	room.sceneId = sceneId
	room.objects = objects
	return nil
}
