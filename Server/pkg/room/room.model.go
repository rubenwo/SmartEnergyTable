package room

type Room struct {
	id      string
	sceneId int
	objects []SceneObject

	master  interface{}
	clients []interface{}
}

type SceneObject struct {
}
