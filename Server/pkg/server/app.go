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

	// TODO: Refactor the csv file loading to implement caching etc.
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
