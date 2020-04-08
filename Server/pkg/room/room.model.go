package room

type Data struct {
	ID      string
	SceneID int
	Objects []SceneObject
}

type Room struct {
	Data Data

	master  chan Data
	clients []chan Data
}

type SceneObject struct {
}
