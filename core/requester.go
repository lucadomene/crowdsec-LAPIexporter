package core

import (
	"encoding/json"
	"fmt"
	// "io"

	// "log"
	"lucadomeneghetti/LAPIexporter/models"
	"lucadomeneghetti/LAPIexporter/utils"
	"net/http"
)

func QueryAlerts(limit int8, retry int) (models.Alerts, error) {

	utils.CheckAuth()
	base_url := utils.GetBaseURL()
	req, err := http.NewRequest("GET", fmt.Sprintf("%v/alerts?include_capi=true&limit=%v", base_url, limit), nil)
	if err != nil {
		return nil, err
	}
	
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer " + utils.GetToken())

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	} else if res.StatusCode > 300 && retry > 0 {
		return QueryAlerts(limit, retry - 1)
	} else if retry <= 0 {
		http_err := fmt.Errorf("%v", res.Status)
		return nil, http_err
	}
	defer res.Body.Close()

	var result = []map[string]interface{}{}
	err = json.NewDecoder(res.Body).Decode(&result)
	if err != nil {
		return nil, err
	}

	var alerts models.Alerts
	for _, v := range result {
		var alert models.Alert

		scenario, ok := v["scenario"].(string)
		if ok {
			alert.Scenario = scenario
		}

		source := v["source"].(map[string]interface{}) // type assertion
		ipaddr, ok := source["ip"].(string)
		if ok {
			alert.IPAddress = ipaddr
		}
		
		datetime, ok := v["created_at"].(string)
		if ok {
			alert.DateTime = datetime
		}

		latitude, ok := source["latitude"].(float64)
		if ok {
			alert.Latitude = latitude	
		}

		longitude, ok := source["longitude"].(float64)
		if ok {
			alert.Longitude = longitude
		}

		countryiso, ok := source["cn"].(string)
		if ok {
			alert.Country = countryiso
		}

		subnet, ok := source["range"].(string)
		if ok {
			alert.Subnet = subnet
		}

		alerts = append(alerts, alert)
	}

	return alerts, nil

}

func QueryDecisions() {
	
}
