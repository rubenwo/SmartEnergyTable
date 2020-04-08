package main

import (
	"github.com/go-chi/chi"
	"log"
	"net/http"
)

func main() {

	router := chi.NewRouter()

	log.Println("SmartEnergyTable API is running!")
	if err := http.ListenAndServe(":80", router); err != nil {
		log.Fatal(err)
	}
	log.Println("SmartEnergyTable API exited.")
}
