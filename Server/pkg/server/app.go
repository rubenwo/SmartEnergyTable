package server

import (
	"github.com/rubenwo/SmartEnergyTable/Server/pkg/room"
	"log"
)

func Run() error {
	manager, err := room.NewManager()
	if err != nil {
		return err
	}
	rest := &api{manager: manager}
	gRPC := &server{manager: manager}
	go func() {
		if err := rest.Run(); err != nil {
			log.Fatal(err)
		}
	}()
	if err := gRPC.Run(); err != nil {
		return err
	}
	return nil
}
