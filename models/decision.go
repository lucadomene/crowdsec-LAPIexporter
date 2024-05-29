package models

import (
	"sync"
)

type Decision struct {
	Scenario string `json:"scenario"`
	IPAddress string `json:"ip"`
	Latitude float32 `json:"latitude"`
	Longitude float32 `json:"longitude"`
}

var DecisionsMutex struct {
	Mu sync.Mutex
	Decisions []Decision
}
