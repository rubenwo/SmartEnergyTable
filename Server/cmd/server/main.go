package main

import (
	"context"
	"fmt"
	"github.com/go-chi/chi"
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

func (s *server) JoinRoom(roomId *v1.RoomId, stream v1.SmartEnergyTableService_JoinRoomServer) error {
	log.Println("Someone joined room:", roomId.Id)
	cb := make(chan room.Data)
	log.Println(s.manager.JoinRoom(roomId.Id, cb))
	timer := time.Tick(time.Second * 2)

Loop:
	for {
		select {
		case data := <-cb:
			log.Println("player still connected")
			start := time.Now()

			if err := stream.Send(&v1.Update{Id: data.ID, Room: &v1.Room{Id: data.ID, SceneId: int32(data.SceneID)}}); err != nil {
				log.Println(err)
				break Loop
			}
			log.Println("Sending update message took:", time.Since(start).Microseconds(), "microseconds.")
		case <-timer:
			if err := stream.Send(&v1.Update{Id: "-1"}); err != nil {
				log.Println(err)
				break Loop
			}
		}
	}
	log.Println("Closed")
	_ = s.manager.RemoveClient(roomId.Id, cb)
	return nil
}

func (s *server) CreateRoom(ctx context.Context, empty *v1.Empty) (*v1.Room, error) {
	id := s.manager.CreateRoom()
	return &v1.Room{Id: id}, nil
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
	roomManager, err := room.NewManager()
	if err != nil {
		log.Fatal(err)
	}

	roomManager.CreateRoom()
	router := chi.NewRouter()
	router.Get("/healthz", func(writer http.ResponseWriter, request *http.Request) {
		_, _ = writer.Write([]byte("Hello world"))
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
