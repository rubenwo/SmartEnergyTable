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

//CreateRoom is a implementation of the gRPC services interface created by the .proto file.
func (s *server) CreateRoom(ctx context.Context, empty *v1.Empty) (*v1.RoomUser, error) {
	id := s.manager.CreateRoom()

	log.Println("CreateRoom() => instantiated room:", id)

	return &v1.RoomUser{Id: id}, nil
}

//JoinRoom is a implementation of the gRPC services interface created by the .proto file.
func (s *server) JoinRoom(roomId *v1.RoomUser, stream v1.SmartEnergyTableService_JoinRoomServer) error {

	log.Println("JoinRoom() => A client:", roomId.UserId, "has joined room:", roomId.Id)

	patches := make(chan room.Patch)
	if err := s.manager.JoinRoom(roomId.Id, roomId.UserId, patches); err != nil {
		return err
	}

	for {
		//No select statement as we're only listening from 1 channel. This also insures we never send multiple message concurrently over the same 'stream' channel.
		patch, ok := <-patches // Start listening for patches on the callback channel. This is blocking, now we have an event-based update loop.
		if !ok {
			break //Break the loop when the channel is closed. This happens when LeaveRoom is called.
			// We also need to break in order to allow the client to close down. Especially as the Unity Editor will crash if we don't.
		}
		start := time.Now() //Timer for debugging

		//Copy the diffs from our internal architecture to the protocol messages.
		diffs := make([]*v1.Diff, len(patch.Diffs))
		for i, diff := range patch.Diffs {

			diffs[i] = &v1.Diff{
				Action: 0,
				Token:  diff.Token,
			}

			switch diff.Action {
			case room.MOVE:
				diffs[i].Action = v1.Diff_MOVE
				break
			case room.ADD:
				diffs[i].Action = v1.Diff_ADD
				break
			case room.DELETE:
				diffs[i].Action = v1.Diff_DELETE
				break
			}
		}

		history := make([]*v1.Diff, len(patch.History))
		for i, diff := range patch.History {
			history[i] = &v1.Diff{
				Action: 0,
				Token:  diff.Token,
			}
			switch diff.Action {
			case room.ADD:
				history[i].Action = v1.Diff_ADD
				break
			default:
				log.Println("JoinRoom() => ERROR: cleaning of the history didn't work correctly, there was a:", diff.Action, "left")
			}
		}

		//Finally send the update message to the client
		if err := stream.Send(&v1.Patch{
			RoomId:       patch.RoomID,
			SceneId:      int32(patch.SceneID),
			UserPosition: &patch.UserPosition,
			IsMaster:     patch.IsMaster,
			Diffs:        diffs,
			History:      history,
		}); err != nil {
			log.Println(err)
			return err
		}
		log.Println("JoinRoom() => Update took:", time.Since(start).Microseconds(), "microseconds to client", roomId.UserId)

	}
	log.Println("JoinRoom() => Connection to client:", roomId.UserId, "closed.")
	return nil
}

//SaveRoom is a implementation of the gRPC services interface created by the .proto file.
func (s *server) SaveRoom(ctx context.Context, room *v1.RoomUser) (*v1.Empty, error) {
	if err := s.manager.SaveRoom(room.Id); err != nil {
		return &v1.Empty{}, fmt.Errorf("error occurred when saving room: %s error: %w", room.Id, err)
	}

	log.Println("SaveRoom() => Room:", room.Id, " is saved to the database.", "by:")
	return &v1.Empty{}, nil
}

//AddToken is a implementation of the gRPC services interface created by the .proto file.
func (s *server) AddToken(ctx context.Context, Token *v1.Token) (*v1.Empty, error) {
	if err := s.manager.AddToken(Token.RoomUser.Id, Token.RoomUser.UserId, Token); err != nil {
		return &v1.Empty{}, fmt.Errorf("error occurred when adding Token: %d from the scene: %w", Token.ObjectIndex, err)
	}
	log.Println("AddToken() => Token with objectLibrary index:", Token.ObjectIndex, "has been added to room:", Token.RoomUser.Id, "by:", Token.RoomUser.UserId)

	return &v1.Empty{}, nil
}

//RemoveToken is a implementation of the gRPC services interface created by the .proto file.
func (s *server) RemoveToken(ctx context.Context, Token *v1.Token) (*v1.Empty, error) {
	if err := s.manager.RemoveToken(Token.RoomUser.Id, Token.RoomUser.UserId, Token); err != nil {
		return &v1.Empty{}, fmt.Errorf("error occurred when removing Token: %s from the scene: %w", Token.ObjectId, err)
	}
	log.Println("RemoveToken() => Token with uuid:", Token.ObjectId, "has been removed from room:", Token.RoomUser.Id, "by:", Token.RoomUser.UserId)

	return &v1.Empty{}, nil
}

//MoveToken is a implementation of the gRPC services interface created by the .proto file.
func (s *server) MoveToken(ctx context.Context, Token *v1.Token) (*v1.Empty, error) {
	if err := s.manager.MoveToken(Token.RoomUser.Id, Token.RoomUser.UserId, Token); err != nil {
		return &v1.Empty{}, fmt.Errorf("error occurred when moving Token: %s in the scene: %w", Token.ObjectId, err)
	}
	log.Println("MoveToken() => Token with uuid:", Token.ObjectId, "has been moved in room:", Token.RoomUser.Id, "by:", Token.RoomUser.UserId)

	return &v1.Empty{}, nil
}

//ChangeScene is a implementation of the gRPC services interface created by the .proto file.
func (s *server) ChangeScene(ctx context.Context, scene *v1.Scene) (*v1.Empty, error) {
	if err := s.manager.ChangeScene(scene.RoomUser.Id, scene.RoomUser.UserId, int(scene.SceneId)); err != nil {
		return &v1.Empty{}, err
	}
	log.Println("ChangeScene() => Scene in room:", scene.RoomUser.Id, "has been changed to:", scene.SceneId, "by:", scene.RoomUser.UserId)
	return &v1.Empty{}, nil
}

//MoveUsers is a implementation of the gRPC services interface created by the .proto file.
func (s *server) MoveUsers(ctx context.Context, position *v1.UserPosition) (*v1.Empty, error) {
	if err := s.manager.MoveUsers(position.RoomUser.Id, position.RoomUser.UserId, *position.NewPosition); err != nil {
		return &v1.Empty{}, err
	}
	log.Println("MoveUsers() => Users in room:", position.RoomUser.Id, "have been moved to:", position.NewPosition, "by:", position.RoomUser.UserId)
	return &v1.Empty{}, nil
}

//LeaveRoom is a implementation of the gRPC services interface created by the .proto file.
func (s *server) LeaveRoom(ctx context.Context, roomId *v1.RoomUser) (*v1.Empty, error) {
	if err := s.manager.RemoveClient(roomId.Id, roomId.UserId); err != nil {
		log.Println(err)
		return &v1.Empty{}, nil
	}
	log.Println("LeaveRoom() => User:", roomId.UserId, "left room:", roomId.Id)
	return &v1.Empty{}, nil
}

//ChangeMaster is a implementation of the gRPC services interface created by the .proto file.
func (s *server) ChangeMaster(ctx context.Context, user *v1.MasterSwitch) (*v1.Empty, error) {
	if err := s.manager.ChangeMaster(user.Id, user.MasterId, user.NewMasterId); err != nil {
		log.Println(err)
		return &v1.Empty{}, err
	}
	log.Println("ChangeMaster() =>", user.NewMasterId, "is now the new master of room:", user.Id, "by:", user.MasterId)

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
