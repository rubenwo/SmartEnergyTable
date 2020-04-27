package room

import (
	v1 "github.com/rubenwo/SmartEnergyTable/Server/pkg/api/v1"
	"sync"
)

type Action uint

const (
	ADD Action = iota
	DELETE
	MOVE
)

type Diff struct {
	Action Action
	Token  *v1.Token
}

type Patch struct {
	RoomID string
	Tokens []*v1.Token

	SceneID      int
	Diffs        []Diff
	UserPosition v1.Vector3
	IsMaster     bool
}

type scene struct {
	id           int
	tokens       map[string]*v1.Token
	userPosition v1.Vector3
}

type Room struct {
	Lock   sync.Mutex
	RoomID string

	changes []Diff

	history []Patch

	scenes       []scene
	currentScene int

	master  string
	clients map[string]chan Patch //clients are just a map with key:userId and value:callback channel
}

//Size returns the amount of connected clients in the room.
func (r *Room) Size() int {
	return len(r.clients)
}

//Notify should be called after every altering of the Data struct inside the room.
func (r *Room) Notify() {
	r.Lock.Lock()
	//patch := r.generatePatch()
	patch := Patch{
		RoomID:       r.RoomID,
		SceneID:      r.currentScene,
		Diffs:        r.changes,
		UserPosition: r.scenes[r.currentScene].userPosition,
		IsMaster:     false,
		Tokens:       make([]*v1.Token, len(r.scenes[r.currentScene].tokens)),
	}
	index := 0
	for _, token := range r.scenes[r.currentScene].tokens {
		patch.Tokens[index] = token
		index++
	}

	r.changes = []Diff{}

	r.history = append(r.history, patch)
	for user, c := range r.clients {
		patch.IsMaster = user == r.master
		c <- patch
	}
	r.Lock.Unlock()
}
