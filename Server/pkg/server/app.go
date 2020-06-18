package server

import (
	"fmt"
	"github.com/rubenwo/SmartEnergyTable/Server/pkg/database"
	"github.com/rubenwo/SmartEnergyTable/Server/pkg/room"
	"log"
)

func Run(dbName string) error {
	db, err := database.Factory(dbName)
	if err != nil {
		return fmt.Errorf("error creating database: %s with error: %w", dbName, err)
	}

	manager := room.NewManager(db)

	eUsers, err := readSummary("/data/Feb (full).csv")
	if err != nil {
		return err
	}
	eHoursDemand, err := readHourlies("/data/Demand.csv")
	if err != nil {
		return err
	}

	rest := &api{manager: manager, EnergyData: &struct {
		EnergyUser         []EnergyUser
		EnergyDemandHourly []EnergyDemandHourly
	}{EnergyUser: eUsers, EnergyDemandHourly: eHoursDemand}}

	gRPC := &server{manager: manager, energyData: &struct {
		energyUser         []EnergyUser
		energyDemandHourly []EnergyDemandHourly
	}{energyUser: eUsers, energyDemandHourly: eHoursDemand}}

	go func() { // Run the REST API server concurrently
		if err := rest.Run(); err != nil {
			log.Fatal(err)
		}
	}()
	if err := gRPC.Run(); err != nil { // Block on the gRPC server
		return err
	}
	return nil
}
