package main

import (
	"context"
	"fmt"
	"github.com/go-chi/chi"
	"github.com/google/uuid"
	v1 "github.com/rubenwo/SmartEnergyTable/Server/pkg/api/v1"
	"github.com/rubenwo/SmartEnergyTable/Server/pkg/room"
	"google.golang.org/grpc"
	"log"
	"net"
	"net/http"
	"time"
)

type server struct {
	manager *room.Manager
}

func (s *server) CreateRoom(ctx context.Context, empty *v1.Empty) (*v1.Room, error) {
	id := s.manager.CreateRoom()
	return &v1.Room{Id: id}, nil
}

func (s *server) JoinRoom(roomId *v1.RoomUser, stream v1.SmartEnergyTableService_JoinRoomServer) error {
	log.Println("Someone joined room:", roomId.Id)
	cb := make(chan room.Data)
	if err := s.manager.JoinRoom(roomId.Id, roomId.UserId, cb); err != nil {
		return err
	}

	for {
		data, ok := <-cb
		if !ok {
			break
		}
		log.Println("player still connected")
		start := time.Now()
		objs := make([]*v1.GameObject, len(data.Objects))
		for index, objData := range data.Objects {
			objs[index] = &v1.GameObject{
				ObjectName: objData.Name,
				Position: &v1.Vector3{
					X: objData.PosX,
					Y: objData.PosY,
					Z: objData.PosZ,
				},
			}
		}
		if err := stream.Send(&v1.Update{Id: data.ID, Room: &v1.Room{Id: data.ID, SceneId: int32(data.SceneID), Objects: objs}}); err != nil {
			return err
		}
		log.Println("Sending update message took:", time.Since(start).Microseconds(), "microseconds.")

	}
	log.Println("Closed")
	return nil
}

func (s *server) SaveRoom(ctx context.Context, room *v1.Room) (*v1.Empty, error) {
	panic("implement me")
}

func (s *server) AddGameObject(ctx context.Context, gameObject *v1.GameObject) (*v1.Empty, error) {
	log.Println("Adding gameobject")
	_ = s.manager.AddGameObject(gameObject.RoomUser.Id, gameObject.RoomUser.UserId, gameObject)
	return &v1.Empty{}, nil
}

func (s *server) RemoveGameObject(ctx context.Context, gameObject *v1.GameObject) (*v1.Empty, error) {
	panic("implement me")
}

func (s *server) MoveGameObject(ctx context.Context, gameObject *v1.GameObject) (*v1.Empty, error) {
	panic("implement me")
}

func (s *server) ChangeScene(ctx context.Context, scene *v1.Scene) (*v1.Empty, error) {
	if err := s.manager.UpdateRoom(scene.RoomUser.Id, int(scene.SceneId), nil); err != nil {
		return &v1.Empty{}, err
	}
	return &v1.Empty{}, nil
}

func (s *server) MoveUsers(ctx context.Context, position *v1.UserPosition) (*v1.Empty, error) {
	panic("implement me")
}

func (s *server) LeaveRoom(ctx context.Context, roomId *v1.RoomUser) (*v1.Empty, error) {
	log.Println("LeaveRoom() =>", roomId.UserId)
	if err := s.manager.RemoveClient(roomId.Id, roomId.UserId); err != nil {
		log.Println(err)
		return &v1.Empty{}, nil
	}
	return &v1.Empty{}, nil
}

func main() {
	roomManager, err := room.NewManager()
	if err != nil {
		log.Fatal(err)
	}

	roomManager.CreateRoom()
	router := chi.NewRouter()
	router.Get("/healthz", func(writer http.ResponseWriter, request *http.Request) {
		_, _ = writer.Write([]byte("Hello world"))
	})
	router.Get("/join", func(writer http.ResponseWriter, request *http.Request) {
		id := request.URL.Query().Get("id")
		log.Println(id)
		cb := make(chan room.Data)
		if err := roomManager.JoinRoom(id, uuid.New().String(), cb); err != nil {
			log.Fatal(err)
		}
		for {
			data, ok := <-cb
			if !ok {
				break
			}
			log.Println(data)
		}
		writer.Write([]byte("Bye"))

	})
	go func() {
		log.Println("SmartEnergyTable API is running!")
		if err := http.ListenAndServe(":80", router); err != nil {
			log.Fatal(err)
		}
		log.Println("SmartEnergyTable API exited.")
	}()

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", 8080))
	if err != nil {
		log.Fatal("can't listen on port 8080")
	}
	grpcServer := grpc.NewServer()
	v1.RegisterSmartEnergyTableServiceServer(grpcServer, &server{manager: roomManager})
	log.Println("grpc server started listening on port:", lis.Addr())
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatal(err)
	}
	log.Println("grpc server stopped listening")
}
