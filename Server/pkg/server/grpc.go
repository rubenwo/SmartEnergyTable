package server

import (
	"context"
	"errors"
	"fmt"
	v1 "github.com/rubenwo/SmartEnergyTable/Server/pkg/api/v1"
	"github.com/rubenwo/SmartEnergyTable/Server/pkg/room"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"log"
	"net"
	"time"
)

type server struct {
	manager    *room.Manager
	energyData *struct {
		energyUser         []EnergyUser
		energyDemandHourly []EnergyDemandHourly
	}
}

func (s *server) Run() error {
	//TODO: Rewrite RoomUser message to JWT which is sent with every request
	if s.manager == nil {
		return errors.New("room manager is nil")
	}
	// Start the gRPC server in the main goroutine to avoid exiting.
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", 8443))
	if err != nil {
		log.Fatal("can't listen on port: 8443", err)
	}

	creds, err := credentials.NewServerTLSFromFile("/certs/server.pem", "/certs/server.key")
	if err != nil {
		log.Fatal("can't load certificates:", err)
	}

	grpcServer := grpc.NewServer(grpc.Creds(creds))
	v1.RegisterSmartEnergyTableServiceServer(grpcServer, s)
	log.Println("grpc server started listening on port:", lis.Addr())
	if err := grpcServer.Serve(lis); err != nil {
		return err
	}
	log.Println("grpc server stopped listening")
	return nil
}

// CreateRoom is a implementation of the gRPC services interface created by the .proto file.
func (s *server) CreateRoom(ctx context.Context, empty *v1.Empty) (*v1.RoomUser, error) {
	id := s.manager.CreateRoom()

	log.Println("CreateRoom() => instantiated room:", id)

	return &v1.RoomUser{Id: id}, nil
}

// JoinRoom is a implementation of the gRPC services interface created by the .proto file.
func (s *server) JoinRoom(roomID *v1.RoomUser, stream v1.SmartEnergyTableService_JoinRoomServer) error {
	log.Println("JoinRoom() => A client:", roomID.UserId, "has joined room:", roomID.Id)

	patches := make(chan room.Patch)
	if err := s.manager.JoinRoom(roomID.Id, roomID.UserId, patches); err != nil {
		return err
	}
	for {
		// No select statement as we're only listening from 1 channel. This also insures we never send multiple message
		// concurrently over the same 'stream' channel.
		patch, ok := <-patches //  Start listening for patches on the callback channel. This is blocking, now we have an
		//  event-based update loop.
		if !ok {
			break // Break the loop when the channel is closed. This happens when LeaveRoom is called.
			//  We also need to break in order to allow the client to close down. Especially as the Unity Editor will
			// crash if we don't.
		}
		start := time.Now() // Timer for debugging

		// Copy the diffs from our internal architecture to the protocol messages.
		diffs := make([]*v1.Diff, len(patch.Diffs))
		for i, diff := range patch.Diffs {
			diffs[i] = &v1.Diff{
				Action: 0,
				Token:  diff.Token,
			}

			switch diff.Action {
			case room.MOVE:
				diffs[i].Action = v1.Diff_MOVE
			case room.ADD:
				diffs[i].Action = v1.Diff_ADD
			case room.DELETE:
				diffs[i].Action = v1.Diff_DELETE
			}
		}

		history := make([]*v1.Diff, len(patch.History)) // make instead of append for marginal performance benefits.
		//  ('make' allocates once)
		for i, diff := range patch.History {
			history[i] = &v1.Diff{
				Action: 0,
				Token:  diff.Token,
			}
			switch diff.Action {
			case room.ADD:
				history[i].Action = v1.Diff_ADD
			default:
				log.Println(
					"JoinRoom() => ERROR: cleaning of the history didn't work correctly, there was a:",
					diff.Action,
					"left")
			}
		}

		// Finally send the update message to the client
		if err := stream.Send(&v1.Patch{
			RoomId:       patch.RoomID,
			SceneId:      int32(patch.SceneID),
			UserPosition: &patch.UserPosition,
			IsMaster:     patch.IsMaster,
			Diffs:        diffs,
			History:      history,
			Energy:       &v1.GeneratedEnergy{Data: patch.GenEnergyData},
		}); err != nil {
			log.Println(err)
			return err
		}
		log.Println("JoinRoom() => Update took:",
			time.Since(start).Microseconds(),
			"microseconds to client",
			roomID.UserId)
	}
	log.Println("JoinRoom() => Connection to client:", roomID.UserId, "closed.")
	return nil
}

// SaveRoom is a implementation of the gRPC services interface created by the .proto file.
func (s *server) SaveRoom(ctx context.Context, roomUser *v1.RoomUser) (*v1.Empty, error) {
	if err := s.manager.SaveRoom(roomUser.Id); err != nil {
		return &v1.Empty{}, fmt.Errorf("error occurred when saving room: %s error: %w", roomUser.Id, err)
	}

	log.Println("SaveRoom() => Room:", roomUser.Id, " is saved to the database.", "by:")
	return &v1.Empty{}, nil
}

// AddToken is a implementation of the gRPC services interface created by the .proto file.
func (s *server) AddToken(ctx context.Context, token *v1.Token) (*v1.Empty, error) {
	if err := s.manager.AddToken(token.RoomUser.Id, token.RoomUser.UserId, token); err != nil {
		return &v1.Empty{}, fmt.Errorf("error occurred when adding Token: %d from the scene: %w",
			token.ObjectIndex, err)
	}
	log.Println("AddToken() => Token with objectLibrary index:", token.ObjectIndex, "has been added to room:",
		token.RoomUser.Id, "by:", token.RoomUser.UserId)

	return &v1.Empty{}, nil
}

// RemoveToken is a implementation of the gRPC services interface created by the .proto file.
func (s *server) RemoveToken(ctx context.Context, token *v1.Token) (*v1.Empty, error) {
	if err := s.manager.RemoveToken(token.RoomUser.Id, token.RoomUser.UserId, token); err != nil {
		return &v1.Empty{}, fmt.Errorf("error occurred when removing Token: %s from the scene: %w",
			token.ObjectId, err)
	}
	log.Println("RemoveToken() => Token with uuid:", token.ObjectId, "has been removed from room:",
		token.RoomUser.Id, "by:", token.RoomUser.UserId)

	return &v1.Empty{}, nil
}

// MoveToken is a implementation of the gRPC services interface created by the .proto file.
func (s *server) MoveToken(ctx context.Context, token *v1.Token) (*v1.Empty, error) {
	if err := s.manager.MoveToken(token.RoomUser.Id, token.RoomUser.UserId, token); err != nil {
		return &v1.Empty{}, fmt.Errorf("error occurred when moving Token: %s in the scene: %w",
			token.ObjectId, err)
	}
	log.Println("MoveToken() => Token with uuid:", token.ObjectId, "has been moved in room:", token.RoomUser.Id,
		"by:", token.RoomUser.UserId)

	return &v1.Empty{}, nil
}

func (s *server) ClearRoom(ctx context.Context, roomUser *v1.RoomUser) (*v1.Empty, error) {
	if err := s.manager.ClearRoom(roomUser.Id, roomUser.UserId); err != nil {
		return &v1.Empty{}, fmt.Errorf("error occurred when clearing room: %s, %w", roomUser.Id, err)
	}
	log.Println("ClearRoom() => room:", roomUser.Id, "was cleared by:", roomUser.UserId)
	return &v1.Empty{}, nil
}

// ChangeScene is a implementation of the gRPC services interface created by the .proto file.
func (s *server) ChangeScene(ctx context.Context, scene *v1.Scene) (*v1.Empty, error) {
	if err := s.manager.ChangeScene(scene.RoomUser.Id, scene.RoomUser.UserId, int(scene.SceneId)); err != nil {
		return &v1.Empty{}, err
	}
	log.Println("ChangeScene() => Scene in room:", scene.RoomUser.Id, "has been changed to:", scene.SceneId, "by:",
		scene.RoomUser.UserId)
	return &v1.Empty{}, nil
}

// MoveUsers is a implementation of the gRPC services interface created by the .proto file.
func (s *server) MoveUsers(ctx context.Context, position *v1.UserPosition) (*v1.Empty, error) {
	if err := s.manager.MoveUsers(position.RoomUser.Id, position.RoomUser.UserId, *position.NewPosition); err != nil {
		return &v1.Empty{}, err
	}
	log.Println("MoveUsers() => Users in room:", position.RoomUser.Id, "have been moved to:", position.NewPosition,
		"by:", position.RoomUser.UserId)
	return &v1.Empty{}, nil
}

// LeaveRoom is a implementation of the gRPC services interface created by the .proto file.
func (s *server) LeaveRoom(ctx context.Context, roomID *v1.RoomUser) (*v1.Empty, error) {
	if err := s.manager.RemoveClient(roomID.Id, roomID.UserId); err != nil {
		log.Println(err)
		return &v1.Empty{}, nil
	}
	log.Println("LeaveRoom() => User:", roomID.UserId, "left room:", roomID.Id)
	return &v1.Empty{}, nil
}

// ChangeMaster is a implementation of the gRPC services interface created by the .proto file.
func (s *server) ChangeMaster(ctx context.Context, user *v1.MasterSwitch) (*v1.Empty, error) {
	if err := s.manager.ChangeMaster(user.Id, user.MasterId, user.NewMasterId); err != nil {
		log.Println(err)
		return &v1.Empty{}, err
	}
	log.Println("ChangeMaster() =>", user.NewMasterId, "is now the new master of room:", user.Id, "by:", user.MasterId)

	return &v1.Empty{}, nil
}

func (s *server) GetEnergyData(ctx context.Context, roomID *v1.RoomUser) (*v1.EnergyData, error) {
	energyUser := make([]*v1.EnergyUser, len(s.energyData.energyUser))
	energyDataHourly := make([]*v1.EnergyDemandHourly, len(s.energyData.energyDemandHourly))
	for index, data := range s.energyData.energyUser {
		energyUser[index] = &v1.EnergyUser{
			Time:        data.Time,
			Label:       data.Label,
			Name:        data.Name,
			SourceId:    data.SourceId,
			TotalDemand: data.TotalDemand,
			Lighting:    data.Lighting,
			Hvac:        data.HVAC,
			Appliances:  data.Appliances,
			Lab:         data.Lab,
			Pv:          data.PV,
			Unit:        data.Unit,
		}
	}
	for index, data := range s.energyData.energyDemandHourly {
		energyDataHourly[index] = &v1.EnergyDemandHourly{
			Id:               data.Id,
			Date:             data.Date,
			Year:             data.Year,
			Month:            data.Month,
			Day:              data.Day,
			Hour:             data.Hour,
			Minutes:          data.Minutes,
			SourceId:         data.SourceId,
			ChannelId:        data.ChannelId,
			Unit:             data.Unit,
			TotalDemand:      data.TotalDemand,
			DeltaValue:       data.DeltaValue,
			SourceTag:        data.SourceTag,
			ChannelTag:       data.ChannelTag,
			Label:            data.Label,
			Name:             data.Name,
			Height:           data.Height,
			Area:             data.Area,
			WindSpeed:        data.WindSpeed,
			Temperature:      data.Temperature,
			SolarRad:         data.SolarRad,
			ElectricityPrice: data.ElectricityPrice,
			Supply:           data.Supply,
			Renewables:       data.Renewables,
		}
	}
	log.Println("GetEnergyData() => User:", roomID.UserId, "requested the energy data.")

	return &v1.EnergyData{
		EnergyUsers:        energyUser,
		EnergyDemandHourly: energyDataHourly,
	}, nil
}
