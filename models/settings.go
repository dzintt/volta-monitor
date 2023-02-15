package models

type Settings struct {
	Webhook     string   `json:"webhook"`
	LocationIDs []string `json:"locationIDs"`
}
