package main

import (
	"flag"
	"fmt"
	"github.com/rubenwo/SmartEnergyTable/Server/pkg/server"
	"log"
)

func main() {
	db := flag.String("database", "jsonDB", "possible database type. Valid values are `jsonDB` and `redis`.")
	flag.Parse()

	fmt.Println("using database:", *db)
	if err := server.Run(*db); err != nil {
		log.Fatal(err)
	}
}
