package processor

import (
	"fmt"
	"strings"
	"sync"
	"time"
	"volta-monitor/models"
	"volta-monitor/utils/discord"
	"volta-monitor/utils/handlers"
	"volta-monitor/utils/volta"
)

func monitor(locationId string, c chan []models.ChargerStation) {
	for {
		client := volta.NewClient()
		info, err := client.GetInformation(locationId)
		if err != nil {
			continue
		}
		var stationList []models.ChargerStation
		for _, location := range info.Data.LocationByNodeID.StationsByLocationID.Edges {
			stationName := location.Node.Name
			if location.Node.Status != "ACTIVE" {
				fmt.Println("Station is not active: " + stationName)
				continue
			}
			for _, station := range location.Node.Evses.Edges {

				connectors := []string{}
				for _, connector := range station.Node.Connectors.Nodes {
					connectors = append(connectors, connector.Type)
				}
				a := models.ChargerStation{
					LocationName: info.Data.LocationByNodeID.Name,
					Name:         stationName,
					ID:           location.Node.ID,
					TimeLimit:    int64(location.Node.ChargeDurationSeconds),
					Image:        info.Data.LocationByNodeID.ImageURL,
					Address:      info.Data.LocationByNodeID.StreetAddress + ", " + info.Data.LocationByNodeID.City + ", " + info.Data.LocationByNodeID.State + " " + string(info.Data.LocationByNodeID.ZipCode),
					Description:  info.Data.LocationByNodeID.Tips,
				}
				a.Availability.State = station.Node.State
				a.Availability.ID = station.Node.ID
				a.Availability.Connectors = connectors
				stationList = append(stationList, a)
			}
		}
		c <- stationList
		//<-time.After(5 * time.Second)
		time.Sleep(1 * time.Second)
	}
}

func Receiver(channel chan []models.ChargerStation) {
	var ChangeList = map[string]models.ChargerStation{}

	for incoming := range channel {
		for _, station := range incoming {
			if _, ok := ChangeList[station.ID]; !ok {
				fmt.Println("Station Being Monitored: " + station.Name)
				ChangeList[station.ID] = station
				continue
			}

			// fmt.Println(ChangeList[station.ID].Name + ": " + ChangeList[station.ID].Availability.State)
			// fmt.Println(station.Name + " " + station.Availability.State)
			// fmt.Println("-----------------------------------------")
			if ChangeList[station.ID].Availability.State != station.Availability.State {
				ChangeList[station.ID] = station

				err := discord.SendToDiscord(handlers.Settings.Webhook, station)
				if err != nil {
					fmt.Println(err)
				}

				fmt.Printf("CHANGE DETECTED\n")
				fmt.Printf("Location: %s\n", station.LocationName)
				fmt.Printf("ID: %s\n", station.ID)
				fmt.Printf("Name: %s\n", station.Name)
				fmt.Printf("State: %s\n", station.Availability.State)
				fmt.Printf("Time Limit: %d\n", station.TimeLimit)
				fmt.Printf("Connectors: %s\n", strings.Join(station.Availability.Connectors, ", "))
				fmt.Printf("-----------------------------------------\n")
			}

		}
	}
}

func Start() {
	recieverChannel := make(chan []models.ChargerStation)
	defer close(recieverChannel)
	wg := sync.WaitGroup{}
	go Receiver(recieverChannel)
	for _, id := range handlers.Settings.LocationIDs {
		wg.Add(1)
		go func(id string, channel chan []models.ChargerStation) {
			defer wg.Done()
			monitor(id, channel)
		}(id, recieverChannel)
	}
	wg.Wait()
}
