package models

type ChargerStation struct {
	ID           string
	LocationName string //new
	Name         string
	Image        string //new
	Address      string //new
	Description  string //new
	TimeLimit    int64
	Availability struct {
		State      string
		ID         string
		Connectors []string
	}
}
