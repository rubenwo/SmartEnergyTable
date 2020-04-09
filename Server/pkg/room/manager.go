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
	room, ok := m.rooms[id]
	if !ok {
		return fmt.Errorf("room with id: %s does not exist", id)
	}
	if room.master == nil {
		room.master = callback
		log.Println("Master joined")
	} else {
		room.clients = append(room.clients, callback)
		log.Println("Client joined")
	}
	log.Println("Sending room data:", room.Data)
	go func(cb chan Data) { cb <- room.Data }(callback)

	return nil
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

func (m *Manager) RemoveClient(id string, callback chan Data) error {
	room, ok := m.rooms[id]
	if !ok {
		return fmt.Errorf("room with id: %s does not exist", id)
	}

	if room.master == callback {
		if len(room.clients) > 0 {
			room.master = room.clients[0]
			return nil
		}
		room.master = nil
	}
	for index, cb := range room.clients {
		if cb == callback {
			room.clients = remove(room.clients, index)
			return nil
		}
	}
	return nil
}
func remove(s []chan Data, i int) []chan Data {
	s[i] = s[len(s)-1]
	// We do not need to put s[i] at the end, as it will be discarded anyway
	return s[:len(s)-1]
}
