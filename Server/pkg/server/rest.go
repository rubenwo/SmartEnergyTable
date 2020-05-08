package server

import (
	"encoding/json"
	"errors"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/rubenwo/SmartEnergyTable/Server/pkg/room"
	"log"
	"net/http"
	"time"
)

type api struct {
	manager *room.Manager
}

func (a *api) Run() error {
	if a.manager == nil {
		return errors.New("room manager is nil")
	}
	router := chi.NewRouter()

	// A good base middleware stack
	router.Use(middleware.RequestID)
	router.Use(middleware.RealIP)
	router.Use(middleware.Logger)
	router.Use(middleware.Recoverer)
	router.Use(middleware.URLFormat)
	// Set a timeout value on the request context (ctx), that will signal
	// Through ctx.Done() that the request has timed out and further
	// processing should be stopped.
	router.Use(middleware.Timeout(60 * time.Second))

	// Health check endpoint
	router.Get("/healthz", func(writer http.ResponseWriter, request *http.Request) {
		writer.WriteHeader(http.StatusOK)
	})

	// Debug endpoint
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

	// Debug endpoint
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

	server := &http.Server{
		Addr:         ":80",
		Handler:      router,
		ReadTimeout:  time.Second * 60,
		WriteTimeout: time.Second * 60,
		IdleTimeout:  time.Second * 120,
	}

	// Start the HTTP REST server.
	log.Println("SmartEnergyTable API is running on:", server.Addr)
	return server.ListenAndServe()
}
