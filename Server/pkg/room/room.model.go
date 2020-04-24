package room

//Data is a struct containing all room data the server needs to know and distribute.
//This is the 'source of truth' for every client.
type Data struct {
	ID       string
	SceneID  int
	Objects  map[string]SceneObject //A Map of all the tokens.
	IsMaster bool
}

type Room struct {
	Data Data

	master  string
	clients map[string]chan Data //clients are just a map with key:userId and value:callback channel
}

//Notify should be called after every altering of the Data struct inside the room.
func (r *Room) Notify() {

	//Loop over all the clients and send the updated Data concurrently.
	for key, client := range r.clients {
		go func(c chan Data) {
			if key == r.master {
				r.Data.IsMaster = true
			} else {
				r.Data.IsMaster = false
			}
			c <- r.Data
		}(client)
	}
}

//Vector3 is a 3-dimensional vector.
type Vector3 struct {
	X float32
	Y float32
	Z float32
}

//SceneObject has an index which should correspond to the index in the objectLibrary and a position.
type SceneObject struct {
	Index    int32
	Position Vector3
}
