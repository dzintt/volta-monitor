package models

type VoltaResponse struct {
	Data struct {
		LocationByNodeID struct {
			Name              string `json:"name"`
			StreetAddress     string `json:"streetAddress"`
			City              string `json:"city"`
			State             string `json:"state"`
			ZipCode           string `json:"zipCode"`
			Tips              string `json:"tips"`
			ImageURL          string `json:"imageUrl"`
			HoursByLocationID struct {
				Edges []struct {
					Node struct {
						DayOfWeek                 int    `json:"dayOfWeek"`
						ChargerOperationStartTime string `json:"chargerOperationStartTime"`
						ChargerOperationEndTime   string `json:"chargerOperationEndTime"`
					} `json:"node"`
				} `json:"edges"`
			} `json:"hoursByLocationId"`
			StationsByLocationID struct {
				GeolocationCentroid struct {
					Latitude  float64 `json:"latitude"`
					Longitude float64 `json:"longitude"`
				} `json:"geolocationCentroid"`
				Edges []struct {
					Node struct {
						ID                    string `json:"id"`
						Name                  string `json:"name"`
						StationNumber         int    `json:"stationNumber"`
						ChargeDurationSeconds int    `json:"chargeDurationSeconds"`
						Status                string `json:"status"`
						Evses                 struct {
							Edges []struct {
								Node struct {
									ID         string `json:"id"`
									Level      string `json:"level"`
									State      string `json:"state"`
									Connectors struct {
										Nodes []struct {
											Type string `json:"type"`
										} `json:"nodes"`
									} `json:"connectors"`
								} `json:"node"`
							} `json:"edges"`
						} `json:"evses"`
					} `json:"node"`
				} `json:"edges"`
			} `json:"stationsByLocationId"`
		} `json:"locationByNodeId"`
	} `json:"data"`
}
