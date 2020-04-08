package main

import (
	"github.com/go-chi/chi"
	v1 "github.com/rubenwo/SmartEnergyTable/Server/pkg/api/v1"
	"github.com/rubenwo/SmartEnergyTable/Server/pkg/room"
	"log"
	"net/http"
	"time"
)

type server struct{}

func (s *server) Update(stream v1.SmartEnergyTableService_UpdateServer) error {
	for {
		time.Sleep(time.Second * 2)
		if err := stream.Send(&v1.Empty{}); err != nil {
			log.Fatal(err)
		}
	}
	return nil
}
func main() {
	roomManager, err := room.NewManager()
	if err != nil {
		log.Fatal(err)
	}
	roomManager.CreateRoom()
	router := chi.NewRouter()
	router.Get("/healthz", func(writer http.ResponseWriter, request *http.Request) {
		writer.WriteHeader(http.StatusOK)
	})
	log.Println("SmartEnergyTable API is running!")
	if err := http.ListenAndServe(":80", router); err != nil {
		log.Fatal(err)
	}
	log.Println("SmartEnergyTable API exited.")
}
