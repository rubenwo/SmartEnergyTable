package server

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"io"
	"os"
)

type EnergyUser struct {
	Time        string `json:"time"`
	Label       string `json:"label"`
	Name        string `json:"name"`
	SourceId    string `json:"sourceid"`
	TotalDemand string `json:"totaldemand"`
	Lighting    string `json:"lightning"`
	HVAC        string `json:"hvac"`
	Appliances  string `json:"appliances"`
	Lab         string `json:"lab"`
	PV          string `json:"PV"`
	Unit        string `json:"unit"`
}

type EnergyDemandHourly struct {
	Id               string `json:"id"`
	Date             string `json:"date"`
	Year             string `json:"year"`
	Month            string `json:"month"`
	Day              string `json:"day"`
	Hour             string `json:"hour"`
	Minutes          string `json:"minutes"`
	SourceId         string `json:"sourceid"`
	ChannelId        string `json:"channelid"`
	Unit             string `json:"unit"`
	TotalDemand      string `json:"totaldemand"`
	DeltaValue       string `json:"delta"`
	SourceTag        string `json:"sourcetag"`
	ChannelTag       string `json:"channeltag"`
	Label            string `json:"label"`
	Name             string `json:"name"`
	Height           string `json:"height"`
	Area             string `json:"area"`
	WindSpeed        string `json:"windspeed"`
	Temperature      string `json:"temperature"`
	SolarRad         string `json:"solarrad"`
	ElectricityPrice string `json:"electricityPrice"`
	Supply           string `json:"supply"`
	Renewables       string `json:"renewables"`
}

func readSummary(fileSrc string) ([]EnergyUser, error) {
	csvFile, err := os.Open(fileSrc)
	if err != nil {
		return nil, fmt.Errorf("error opening the file: %w", err)
	}
	reader := csv.NewReader(bufio.NewReader(csvFile))
	var eUsers []EnergyUser
	for {
		line, err := reader.Read()
		if err == io.EOF {
			break
		} else if err != nil {
			return nil, err
		}
		eUsers = append(eUsers, EnergyUser{
			Time:        line[0],
			Label:       line[1],
			Name:        line[2],
			SourceId:    line[3],
			TotalDemand: line[4],
			Lighting:    line[5],
			HVAC:        line[6],
			Appliances:  line[7],
			Lab:         line[8],
			PV:          line[9],
			Unit:        line[10],
		})
	}
	return eUsers, nil
}

func readHourlies(fileSrc string) ([]EnergyDemandHourly, error) {
	csvFile, err := os.Open(fileSrc)
	if err != nil {
		return nil, fmt.Errorf("error opening the file: %w", err)
	}
	reader := csv.NewReader(bufio.NewReader(csvFile))
	reader.FieldsPerRecord = -1

	var eUsers []EnergyDemandHourly
	for {
		line, err := reader.Read()
		if err == io.EOF {
			break
		} else if err != nil {
			return nil, fmt.Errorf("error reading: %w", err)
		}

		if len(line) < 24 {
			continue
		}

		if len(line) == 24 {
			eUsers = append(eUsers, EnergyDemandHourly{
				Id:               line[0],
				Date:             line[1],
				Year:             line[2],
				Month:            line[3],
				Day:              line[4],
				Hour:             line[5],
				Minutes:          line[6],
				SourceId:         line[7],
				ChannelId:        line[8],
				Unit:             line[9],
				TotalDemand:      line[10],
				DeltaValue:       line[11],
				SourceTag:        line[12],
				ChannelTag:       line[13],
				Label:            line[14],
				Name:             line[15],
				Height:           line[16],
				Area:             line[17],
				WindSpeed:        line[18],
				Temperature:      line[19],
				SolarRad:         line[20],
				ElectricityPrice: line[21],
				Supply:           line[22],
				Renewables:       line[23],
			})
		} else if len(line) == 21 {

			eUsers = append(eUsers, EnergyDemandHourly{
				Id:               line[0],
				Date:             line[1],
				Year:             line[2],
				Month:            line[3],
				Day:              line[4],
				Hour:             line[5],
				Minutes:          line[6],
				SourceId:         line[7],
				ChannelId:        line[8],
				Unit:             line[9],
				TotalDemand:      line[10],
				DeltaValue:       line[11],
				SourceTag:        line[12],
				ChannelTag:       line[13],
				Name:             line[14],
				Height:           line[15],
				Area:             line[16],
				WindSpeed:        line[17],
				Temperature:      line[18],
				SolarRad:         line[19],
				ElectricityPrice: line[20],
			})
		}

	}
	return eUsers, nil
}
