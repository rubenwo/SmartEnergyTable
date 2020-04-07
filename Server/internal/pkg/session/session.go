package session

import (
	"github.com/google/uuid"
	"github.com/rubenwo/SmartEnergyTable/Server/internal/protocol"
	"log"
)

type session struct {
	state []protocol.Entity
}

func (s *session) join(user string) {

}

func (s *session) save() error {
	log.Println("Saving session:", s)
	return nil
}

func (s *session) update(state []protocol.Entity) {
	s.state = state
}

func (s *session) destroy() {

}

type owner struct {
	id   string
	user string
}

type Manager struct {
	owners   map[string]owner
	sessions map[owner]*session
}

func NewManager() *Manager {
	return &Manager{owners: make(map[string]owner), sessions: make(map[owner]*session)}
}

func (m *Manager) CreateSession(user string) (id string) {
	id = uuid.New().String()
	m.sessions[owner{id: id, user: user}] = &session{}
	return id
}

func (m *Manager) DestroySession(id string, user string) (ok bool, err error) {
	m.sessions[owner{id: id, user: user}].destroy()
	delete(m.sessions, owner{id: id, user: user})
	return true, nil
}

func (m *Manager) SaveSession(id string, user string) error {
	s := m.sessions[owner{id: id, user: user}]
	if err := s.save(); err != nil {
		return err
	}
	return nil
}

func (m *Manager) JoinSession(id string, user string) error {
	o := m.owners[id]
	s := m.sessions[o]
	s.join(user)

	return nil
}

func (m *Manager) Update(id string, user string, entities []protocol.Entity) {
	m.sessions[owner{id: id, user: user}].update(entities)
}

func (m *Manager) Refresh(id string) []protocol.Entity {
	o := m.owners[id]
	s := m.sessions[o]
	return s.state
}
