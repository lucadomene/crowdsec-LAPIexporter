package core

import (
	"bytes"
	"encoding/json"
	"log"
	"lucadomeneghetti/LAPIexporter/utils"
	"net/http"
	"sync"
	"time"
)

var AuthToken struct {
	mu sync.Mutex
	Expire time.Time
	BearerToken string
}

func InitializeToken() {
	AuthToken.mu.Lock()
	defer AuthToken.mu.Unlock()

	AuthToken.Expire = time.Now()
	AuthToken.BearerToken = ""
}

func CheckAuth() {
	AuthToken.mu.Lock()
	defer AuthToken.mu.Unlock()
	
	if AuthToken.Expire.Before(time.Now()) {
		authenticate()
	}
}

func authenticate() {

	var credentials struct {
		Machine_id string `json:"machine_id"`
		Password string `json:"password"`
	}

	credentials.Machine_id = utils.Config.Authentication.Login
	credentials.Password = utils.Config.Authentication.Password

	credentials_json, err := json.Marshal(credentials)
	if err != nil {
		log.Fatal(err)
	}
	
	base_url := utils.Config.Server.Protocol + "://" + utils.Config.Server.Host + ":" + utils.Config.Server.Port
	req, err := http.NewRequest("POST", base_url + "/v1/watchers/login", bytes.NewBuffer(credentials_json))
	if err != nil {
		log.Fatal(err)
	}

	req.Header.Set("Content-Type", "application/json")

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()

	var TokenResponse struct {
		Token string `json:"token"`
		Expire string `json:"expire"`
	}
	json.NewDecoder(res.Body).Decode(&TokenResponse)
	AuthToken.BearerToken = TokenResponse.Token
	AuthToken.Expire, err = time.Parse(time.RFC3339, TokenResponse.Expire)
	if err != nil {
		log.Fatal(err)
	}
}
