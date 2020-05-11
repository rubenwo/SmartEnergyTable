package server

import (
	"encoding/json"
	"errors"
	"github.com/go-chi/chi"
	"github.com/rubenwo/SmartEnergyTable/Server/pkg/room"
	"log"
	"net/http"
)

type api struct {
	manager *room.Manager
}

func (a *api) Run() error {
	if a.manager == nil {
		return errors.New("room manager is nil")
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
		r.Rooms = a.manager.RoomIDs()
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
		r.ID = a.manager.CreateRoom()
		if err := json.NewEncoder(writer).Encode(&r); err != nil {
			log.Println(err)
			writer.WriteHeader(http.StatusInternalServerError)
		}
	})

	//Start the HTTP REST server.
	log.Println("SmartEnergyTable API is running!")
	if err := http.ListenAndServe(":80", router); err != nil {
		return err
	}
	log.Println("SmartEnergyTable API exited.")
	return nil
}
