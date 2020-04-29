package room

import (
	"encoding/json"
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
	RoomID  string
	History []Diff

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

	changes []Diff //This is a slice of the pending changes.

	history []Diff //This is a slice that contains every action that has taken place during the session. When the
	//changes slice is processed those diffs are appended to the history.

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
		History:      r.history, //This history should be clean
	}
	if len(patch.History) == 0 { //The patch history is empty when a master just joined.
		patch.History = append(r.history, r.changes...)
	}

	r.history = append(r.history, r.changes...) //Append the now processed changes to the history.
	r.changes = []Diff{}                        //Clear the pending changes
	r.Lock.Unlock()
	go r.gcHistory() //Run a garbage collection on the history slice concurrently.

	for user, c := range r.clients {
		patch.IsMaster = user == r.master
		c <- patch //Push the patch to the client
	}

}

//gcHistory can be used to reduce the size of the history slice in the room by performing the 'Action' operations from a
//Diff. First adding only the non-deleted diffs to the history, then setting the positions of those diffs if they have
//been moved. This results in a slice where only 'ADD' actions remain.
func (r *Room) gcHistory() {
	r.Lock.Lock()

	var add []Diff
	var move []Diff
	var del []Diff

	//Push the different diffs to their respective slices.
	for _, diff := range r.history {
		switch diff.Action {
		case ADD:
			add = append(add, diff)
			break
		case MOVE:
			move = append(move, diff)
			break
		case DELETE:
			del = append(del, diff)
			break
		}
	}

	r.history = []Diff{} //Clear the history so we can fill it refill it

	for _, diff := range add {
		deleted := false
		for _, d := range del {
			if d.Token.ObjectId == diff.Token.ObjectId {
				deleted = true
				break
			}
		}
		if !deleted {
			r.history = append(r.history, diff) //Only add the diffs that haven't been deleted
		}
	}

	if len(move) > 0 { //Speed optimization
		for _, diff := range r.history {
			for _, d := range move { //Move the diffs that are remaining
				if diff.Token.ObjectId == d.Token.ObjectId {
					diff.Token.Position = d.Token.Position
					diff.Token.Rotation = d.Token.Rotation
				}
			}
		}
	}

	r.Lock.Unlock()
}

//MarshalBinary is an implementation of the encoding.BinaryMarshaller interface.
func (r *Room) MarshalBinary() ([]byte, error) {
	return json.Marshal(r)
}

//UnmarshalBinary is an implementation of the encoding.BinaryMarshaller interface.
func (r *Room) UnmarshalBinary(data []byte) error {
	if err := json.Unmarshal(data, &r); err != nil {
		return err
	}
	return nil
}
