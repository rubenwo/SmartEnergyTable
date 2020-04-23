package room

import "log"

type Data struct {
	ID      string
	SceneID int
	Objects map[string]SceneObject
}

type Room struct {
	Data Data

	master  string
	clients map[string]chan Data
}

func (r *Room) Notify() {
	log.Println("Sending room data:", r.Data)

	for _, client := range r.clients {
		go func(c chan Data) {
			c <- r.Data
		}(client)
	}
}

type Vector3 struct {
	X float32
	Y float32
	Z float32
}

type SceneObject struct {
	Index    int32
	Position Vector3
}
