package session

import "github.com/google/uuid"

type session struct {
}

func (s *session) destroy() {

}

type owner struct {
	id   string
	user string
}

type Manager struct {
	sessions map[owner]*session
}

func NewManager() *Manager {
	return &Manager{sessions: make(map[owner]*session)}
}

func (m *Manager) CreateSession(user string) (id string) {
	id = uuid.New().String()
	m.sessions[owner{id: id, user: user}] = &session{}
	return id
}

func (m *Manager) DestroySession(id string, user string) (ok bool, err error) {
	m.sessions[owner{id: id, user: user}].destroy()

	return true, nil
}
