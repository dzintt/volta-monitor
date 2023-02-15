package models

type VoltaRequest struct {
	Query     string    `json:"query"`
	Variables Variables `json:"variables"`
}
type Variables struct {
	LocationNodeID string `json:"locationNodeId"`
}

func CreatePayload(locationID string) VoltaRequest {
	v := VoltaRequest{}
	v.Query = "\n    query getStation($locationNodeId: ID!) {\n      locationByNodeId(nodeId: $locationNodeId) {\n        name\n        streetAddress\n        city\n        state\n        zipCode\n        tips\n        imageUrl\n        hoursByLocationId {\n          edges {\n            node {\n              dayOfWeek\n              chargerOperationStartTime\n              chargerOperationEndTime\n            }\n          }\n        }\n        stationsByLocationId(orderBy: STATION_NUMBER_ASC) {\n          geolocationCentroid {\n            latitude\n            longitude\n          }\n          edges {\n            node {\n              id\n              name\n              stationNumber\n              chargeDurationSeconds\n              status\n              evses {\n                edges {\n                  node {\n                    id\n                    level\n                    state\n                    connectors {\n                      nodes {\n                        type\n                      }\n                    }\n                  }\n                }\n              }\n            }\n          }\n        }\n      }\n    }\n  "
	v.Variables.LocationNodeID = locationID
	return v
}
