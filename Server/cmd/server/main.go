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
	for {
		time.Sleep(time.Second * 2)
		if err := stream.Send(&v1.Update{}); err != nil {
			log.Fatal(err)
		}
	}
	return nil
}

func (s *server) CreateRoom(ctx context.Context, empty *v1.Empty) (*v1.Room, error) {
	return &v1.Room{}, nil
}

func (s *server) AddGameObject(ctx context.Context, gameObject *v1.GameObject) (*v1.Empty, error) {
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
		roomManager.JoinRoom("")
		writer.Write([]byte("Hello world"))
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
	log.Println("grpc server started listening on port 8080")
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatal(err)
	}
	log.Println("grpc server stopped listening")
}
