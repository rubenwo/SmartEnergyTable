package room

type Room struct {
	id      string
	sceneId int
	objects []SceneObject

	master  chan int
	clients []chan int
}

type SceneObject struct {
}
