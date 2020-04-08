package main

import (
	"context"
	"fmt"
	v1 "github.com/rubenwo/SmartEnergyTable/Server/pkg/api/v1"
	"github.com/rubenwo/SmartEnergyTable/Server/pkg/room"
	"google.golang.org/grpc"
	"log"
	"net"
	"time"
)

type server struct {
	manager *room.Manager
}

func (s *server) JoinRoom(roomId *v1.RoomId, stream v1.SmartEnergyTableService_JoinRoomServer) error {
	log.Println("Someone joined the room")
	cb := make(chan room.Data)
	_ = s.manager.JoinRoom(roomId.Id, cb)
Loop:
	for {
		select {
		case data := <-cb:
			time.Sleep(time.Second * 2)
			log.Println("player still connected")
			start := time.Now()
			if err := stream.Send(&v1.Update{Id: data.ID, Room: &v1.Room{Id: data.ID, SceneId: int32(data.SceneID)}}); err != nil {
				log.Println(err)
				break Loop
			}
			log.Println("Sending update message took:", time.Since(start).Microseconds(), "seconds.")
		}
	}
	return nil
}

func (s *server) CreateRoom(ctx context.Context, empty *v1.Empty) (*v1.Room, error) {
	return &v1.Room{Id: "123"}, nil
}

func (s *server) AddGameObject(ctx context.Context, gameObject *v1.GameObject) (*v1.Empty, error) {
	return &v1.Empty{}, nil
}

func (s *server) ChangeScene(ctx context.Context, scene *v1.Scene) (*v1.Empty, error) {
	if err := s.manager.UpdateRoom(scene.Id, int(scene.SceneId), nil); err != nil {
		return &v1.Empty{}, err
	}
	return &v1.Empty{}, nil
}

func main() {
	//roomManager, err := room.NewManager()
	//if err != nil {
	//	log.Fatal(err)
	//}
	//
	//roomManager.CreateRoom()
	//router := chi.NewRouter()
	//router.Get("/healthz", func(writer http.ResponseWriter, request *http.Request) {
	//	roomManager.JoinRoom("")
	//	writer.Write([]byte("Hello world"))
	//})
	//go func() {
	//	log.Println("SmartEnergyTable API is running!")
	//	if err := http.ListenAndServe(":80", router); err != nil {
	//		log.Fatal(err)
	//	}
	//	log.Println("SmartEnergyTable API exited.")
	//}()

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", 8080))
	if err != nil {
		log.Fatal("can't listen on port 8080")
	}
	grpcServer := grpc.NewServer()
	v1.RegisterSmartEnergyTableServiceServer(grpcServer, &server{})
	log.Println("grpc server started listening on port:", lis.Addr())
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatal(err)
	}
	log.Println("grpc server stopped listening")
}
