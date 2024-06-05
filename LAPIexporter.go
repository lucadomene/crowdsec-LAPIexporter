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

	utils.InitializeToken()

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

		var limit int64
		if r.URL.Query().Has("limit") {
			limit, err = strconv.ParseInt(r.URL.Query().Get("limit"), 10, 64)
			if err != nil {
				w.WriteHeader(http.StatusBadRequest)
				json_response.Encode(err)
				log.Printf("request failed for %v: %v", client, err)
				return
			}
		} else {
			limit = 10
		}

		result, err := core.ReturnAlerts(limit)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			json_response.Encode(err)
			log.Printf("request failed for %v: %v", client, err)
			return
		}

		err = json_response.Encode(result)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			json_response.Encode(err)
			log.Printf("parsing result failed: %v", err)
			return
		} else {
			log.Printf("successfully served %v", client)
			return
		}

	})

	http.HandleFunc("/decisions", func(w http.ResponseWriter, r *http.Request) {
		json_response := json.NewEncoder(w)
		client := r.RemoteAddr

		w.Header().Set("Content-Type", "application/json")

		var limit int64
		if r.URL.Query().Has("limit") {
			limit, err = strconv.ParseInt(r.URL.Query().Get("limit"), 10, 64)
			if err != nil {
				w.WriteHeader(http.StatusBadRequest)
				json_response.Encode(err)
				log.Printf("request failed for %v: %v", client, err)
				return
			}
		} else {
			limit = 10
		}

		result, err := core.ReturnDecisions(limit)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			json_response.Encode(err)
			log.Printf("request failed for %v: %v", client, err)
			return
		}

		err = json_response.Encode(result)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			json_response.Encode(err)
			log.Printf("parsing result failed: %v", err)
			return
		} else {
			log.Printf("successfully served %v", client)
			return
		}

	})

	err = http.ListenAndServe(":5001", nil)
	if err != nil {
		log.Fatal(err)
	}
}
