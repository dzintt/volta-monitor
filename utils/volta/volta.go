package volta

import (
	"bytes"
	"encoding/json"
	"net/http"
	"volta-monitor/models"
)

type ChargeStationInfo struct {
	StationId     string
	Name          string
	StreetAddress string
	Description   string
	Image         string
	State         string
	ConnectorType string
	StationNumber int
}

func (station *VoltaClient) GetInformation(locationId string) (models.VoltaResponse, error) {
	var voltaResponse models.VoltaResponse
	payload := models.CreatePayload(locationId)
	marshaled, err := json.Marshal(payload)
	if err != nil {
		return voltaResponse, err
	}
	req, err := http.NewRequest(
		http.MethodPost,
		"https://api.voltaapi.com/v1/pg-graphql",
		bytes.NewBuffer(marshaled),
	)
	if err != nil {
		return voltaResponse, err
	}
	resp, err := station.ReqDo(*req, map[string]string{})
	if err != nil {
		return voltaResponse, err
	}
	defer resp.Body.Close()

	if err = json.NewDecoder(resp.Body).Decode(&voltaResponse); err != nil {
		return voltaResponse, err
	}
	return voltaResponse, nil
}
