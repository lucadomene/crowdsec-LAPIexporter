package core

import (
	// "fmt"
	// "log"
	"lucadomeneghetti/LAPIexporter/models"
)

func ReturnAlerts(limit int8) (models.Alerts, error) {
	alerts, err := QueryAlerts(limit, 5) 
	if err != nil {
		return nil, err
	} else {
		return alerts, nil
	}
}

func ReturnDecisions(limit int8) (models.Decision, error) {
	decision := models.Decision{}
	return decision, nil
}
