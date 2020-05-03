package server

import (
	"github.com/rubenwo/SmartEnergyTable/Server/pkg/room"
)

func Run() error {
	manager, err := room.NewManager()
	if err != nil {
		return err
	}
	rest := &api{manager: manager}
	gRPC := &server{manager: manager}
	go rest.Run()
	if err := gRPC.Run(); err != nil {
		return err
	}
	return nil
}
