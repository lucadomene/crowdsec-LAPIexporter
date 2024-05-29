package models

type Alert struct {
	Scenario string `json:"scenario"`
	IPAddress string `json:"ip"`
	Subnet string `json:"subnet"`
	CreatedAt string `json:"datetime"`
	Latitude float32 `json:"latitude"`
	Longitude float32 `json:"longitude"`
	Country string `json:"countryISO"`
}

type Alerts []Alert
