package main

import (
	"encoding/json"
	"log"
	"lucadomeneghetti/LAPIexporter/core"
	"lucadomeneghetti/LAPIexporter/utils"
	"net/http"
	"strconv"
)

func initializeServer() error {
	err := utils.ImportConfig("config.yml")
	if err != nil {
		return err
	}

	core.InitializeToken()

	log.Printf("successfully initialized server")
	return nil
}

func main() {
	err := initializeServer()
	if err != nil {
		log.Fatalf("error initializing server: %v", err)
	}
	
	http.HandleFunc("/alerts", func(w http.ResponseWriter, r *http.Request) {
		json_response := json.NewEncoder(w)
		client := r.RemoteAddr

		w.Header().Set("Content-Type", "application/json")
		limit_64, err := strconv.ParseInt(r.URL.Query().Get("limit"), 10, 8)
		limit := int8(limit_64)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			json_response.Encode(err)
			log.Printf("request failed for %v: %v", client, err)
		}
		result, err := core.ReturnAlerts(limit)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			json_response.Encode(err)
			log.Printf("request failed for %v: %v", client, err)
		}

		err = json_response.Encode(result)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			json_response.Encode(err)
			log.Printf("parsing result failed: %v", err)
		} else {
			w.WriteHeader(http.StatusOK)
			json_response.Encode(result)
			log.Printf("successfully served %v", client)
		}

	})

	http.HandleFunc("/decisions", func(w http.ResponseWriter, r *http.Request) {
		json_response := json.NewEncoder(w)
		client := r.RemoteAddr

		w.Header().Set("Content-Type", "application/json")
		limit_64, err := strconv.ParseInt(r.URL.Query().Get("limit"), 10, 8)
		limit := int8(limit_64)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			json_response.Encode(err)
			log.Printf("request failed for %v: %v", client, err)
		}
		result, err := core.ReturnDecisions(limit)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			json_response.Encode(err)
			log.Printf("request failed for %v: %v", client, err)
		}

		err = json_response.Encode(result)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			json_response.Encode(err)
			log.Printf("parsing result failed: %v", err)
		} else {
			w.WriteHeader(http.StatusOK)
			json_response.Encode(result)
			log.Printf("successfully served %v", client)
		}

	})

	err = http.ListenAndServe(":5001", nil)
	if err != nil {
		log.Fatal(err)
	}
}