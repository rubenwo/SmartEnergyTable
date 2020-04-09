package room

import "log"

type Data struct {
	ID      string
	SceneID int
	Objects []SceneObject
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

type SceneObject struct {
	Name string
	PosX float32
	PosY float32
	PosZ float32
}
