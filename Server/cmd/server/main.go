package main

import (
	"context"
	"encoding/json"
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

//TODO: Add logs for every call. (Abuse & Bug detection)

//CreateRoom is a implementation of the gRPC services interface created by the .proto file.
func (s *server) CreateRoom(ctx context.Context, empty *v1.Empty) (*v1.Room, error) {
	id := s.manager.CreateRoom()
	return &v1.Room{Id: id}, nil
}

//JoinRoom is a implementation of the gRPC services interface created by the .proto file.
func (s *server) JoinRoom(roomId *v1.RoomUser, stream v1.SmartEnergyTableService_JoinRoomServer) error {
	log.Println("Someone joined room:", roomId.Id)
	cb := make(chan room.Data)
	if err := s.manager.JoinRoom(roomId.Id, roomId.UserId, cb); err != nil {
		return err
	}

	for {
		//No select statement as we're only listening from 1 channel. This also insures we never send multiple message concurrently over the same 'stream.Send' channel.
		data, ok := <-cb // Start listening for the data structs on the callback channel. This is blocking.
		if !ok {
			break //Break the loop when the channel is closed. This happens when LeaveRoom is called.
			// We also need to break in order to allow the client to close down. Especially as the Unity Editor will crash if we don't.
		}
		log.Println("player still connected")
		start := time.Now() //Timer for debugging
		objs := make([]*v1.Token, len(data.Objects))

		//Copy SceneObjects to protocol 'Token' message.
		index := 0
		for key, objData := range data.Objects {
			objs[index] = &v1.Token{
				ObjectIndex: objData.Index,
				ObjectId:    key,
				Position: &v1.Vector3{
					X: objData.Position.X,
					Y: objData.Position.X,
					Z: objData.Position.Z,
				},
			}
			index++
		}
		//Finally send the update message to the client
		if err := stream.Send(&v1.Update{Id: data.ID, Room: &v1.Room{Id: data.ID, SceneId: int32(data.SceneID), Objects: objs}, IsMaster: data.IsMaster}); err != nil {
			return err
		}
		log.Println("Sending update message took:", time.Since(start).Microseconds(), "microseconds.")

	}
	log.Println("Closed")
	return nil
}

//SaveRoom is a implementation of the gRPC services interface created by the .proto file.
func (s *server) SaveRoom(ctx context.Context, room *v1.Room) (*v1.Empty, error) {
	//TODO: Implement saving the room
	panic("implement me")
}

//AddToken is a implementation of the gRPC services interface created by the .proto file.
func (s *server) AddToken(ctx context.Context, Token *v1.Token) (*v1.Empty, error) {
	log.Println("Adding Token")
	if err := s.manager.AddToken(Token.RoomUser.Id, Token.RoomUser.UserId, Token); err != nil {
		return &v1.Empty{}, fmt.Errorf("error occurred when adding Token: %d from the scene: %w", Token.ObjectIndex, err)
	}
	return &v1.Empty{}, nil
}

//RemoveToken is a implementation of the gRPC services interface created by the .proto file.
func (s *server) RemoveToken(ctx context.Context, Token *v1.Token) (*v1.Empty, error) {
	log.Println("Removing Token")
	if err := s.manager.RemoveToken(Token.RoomUser.Id, Token.RoomUser.UserId, Token); err != nil {
		return &v1.Empty{}, fmt.Errorf("error occurred when removing Token: %s from the scene: %w", Token.ObjectId, err)
	}
	return &v1.Empty{}, nil
}

//MoveToken is a implementation of the gRPC services interface created by the .proto file.
func (s *server) MoveToken(ctx context.Context, Token *v1.Token) (*v1.Empty, error) {
	log.Println("Moving Token")
	if err := s.manager.MoveToken(Token.RoomUser.Id, Token.RoomUser.UserId, Token); err != nil {
		return &v1.Empty{}, fmt.Errorf("error occurred when moving Token: %s in the scene: %w", Token.ObjectId, err)
	}
	return &v1.Empty{}, nil
}

//ChangeScene is a implementation of the gRPC services interface created by the .proto file.
func (s *server) ChangeScene(ctx context.Context, scene *v1.Scene) (*v1.Empty, error) {
	if err := s.manager.ChangeScene(scene.RoomUser.Id, scene.RoomUser.UserId, int(scene.SceneId), nil); err != nil {
		return &v1.Empty{}, err
	}
	return &v1.Empty{}, nil
}

//MoveUsers is a implementation of the gRPC services interface created by the .proto file.
func (s *server) MoveUsers(ctx context.Context, position *v1.UserPosition) (*v1.Empty, error) {
	//TODO: Implement moving users in RoomManager
	panic("implement me")
}

//LeaveRoom is a implementation of the gRPC services interface created by the .proto file.
func (s *server) LeaveRoom(ctx context.Context, roomId *v1.RoomUser) (*v1.Empty, error) {
	log.Println("LeaveRoom() =>", roomId.UserId)
	if err := s.manager.RemoveClient(roomId.Id, roomId.UserId); err != nil {
		log.Println(err)
		return &v1.Empty{}, nil
	}
	return &v1.Empty{}, nil
}

//ChangeMaster is a implementation of the gRPC services interface created by the .proto file.
func (s *server) ChangeMaster(ctx context.Context, user *v1.MasterSwitch) (*v1.Empty, error) {
	log.Println(fmt.Sprintf("ChangeMaster() => Old: %s => New: %s", user.MasterId, user.NewMasterId))
	if err := s.manager.ChangeMaster(user.Id, user.MasterId, user.NewMasterId); err != nil {
		log.Println(err)
		return &v1.Empty{}, err
	}
	return &v1.Empty{}, nil
}

func main() {

	//TODO: clean-up main
	roomManager, err := room.NewManager()
	if err != nil {
		log.Fatal(err)
	}
	router := chi.NewRouter()
	//Health check endpoint
	router.Get("/healthz", func(writer http.ResponseWriter, request *http.Request) {
		writer.WriteHeader(http.StatusOK)
	})

	//Debug endpoint
	router.Get("/rooms", func(writer http.ResponseWriter, request *http.Request) {
		var r struct {
			Rooms []string `json:"rooms"`
		}
		r.Rooms = roomManager.RoomIDs()
		if err := json.NewEncoder(writer).Encode(&r); err != nil {
			log.Println(err)
			writer.WriteHeader(http.StatusInternalServerError)
		}
	})

	//Debug endpoint
	router.Post("/rooms", func(writer http.ResponseWriter, request *http.Request) {
		var r struct {
			ID string `json:"id"`
		}
		r.ID = roomManager.CreateRoom()
		if err := json.NewEncoder(writer).Encode(&r); err != nil {
			log.Println(err)
			writer.WriteHeader(http.StatusInternalServerError)
		}
	})

	//Start the HTTP REST server in another goroutine as it contains a blocking call.
	go func() {
		log.Println("SmartEnergyTable API is running!")
		if err := http.ListenAndServe(":80", router); err != nil {
			log.Fatal(err)
		}
		log.Println("SmartEnergyTable API exited.")
	}()

	//Start the gRPC server in the main goroutine to avoid exiting.
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
