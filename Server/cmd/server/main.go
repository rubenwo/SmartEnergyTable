package main

import (
	"github.com/rubenwo/SmartEnergyTable/Server/pkg/server"
	"log"
)

func main() {
	if err := server.Run(); err != nil {
		log.Fatal(err)
	}
}
