package processor

import (
	"fmt"
	"strings"
	"sync"
	"volta-monitor/utils/volta"
)

var locationIds = []string{
	"WyJzaXRlcyIsIjdkNDgyYWU0LTJjMjgtNGE0Ni1iNTQyLTJkNGNlYzA0NzA1ZSJd",
}

type ChargerStation struct {
	ID           string
	Name         string
	TimeLimit    int64
	Availability struct {
		State      string
		ID         string
		Connectors []string
	}
}

func monitor(locationId string) {
	client := volta.NewClient()

	info, err := client.GetInformation(locationId)

	if err != nil {
		panic("ERROR: " + err.Error())
	}

	var stationList []ChargerStation

	for _, location := range info.Data.LocationByNodeID.StationsByLocationID.Edges {
		stationName := location.Node.Name
		for _, station := range location.Node.Evses.Edges {

			connectors := []string{}
			for _, connector := range station.Node.Connectors.Nodes {
				connectors = append(connectors, connector.Type)
			}
			a := ChargerStation{
				Name:      stationName,
				ID:        location.Node.ID,
				TimeLimit: int64(location.Node.ChargeDurationSeconds),
			}
			a.Availability.State = station.Node.State
			a.Availability.ID = station.Node.ID
			a.Availability.Connectors = connectors
			stationList = append(stationList, a)
		}
	}

	for _, station := range stationList {
		fmt.Printf("ID: %s\n", station.ID)
		fmt.Printf("Name: %s\n", station.Name)
		fmt.Printf("State: %s\n", station.Availability.State)
		fmt.Printf("Time Limit: %d\n", station.TimeLimit)
		fmt.Printf("Connectors: %s\n", strings.Join(station.Availability.Connectors, ", "))
		fmt.Printf("-----------------------------------------\n")
	}
}

// func Receiver(c chan updateStruct){
// 	//for c := range
// }

func Start() {
	// channel := make()
	wg := sync.WaitGroup{}
	for _, id := range locationIds {
		wg.Add(1)
		go func(id string) {
			defer wg.Done()
			monitor(id)
		}(id)
	}
	wg.Wait()
}
