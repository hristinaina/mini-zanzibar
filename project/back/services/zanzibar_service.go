package services

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

type ZanzibarService struct {
	apiKey      string
	client      string
	zanzibarURL string
}

func NewZanzibarService() ZanzibarService {
	return ZanzibarService{
		apiKey:      os.Getenv("ZANZIBAR_SECRET_KEY"),
		client:      os.Getenv("ZANZIBAR_CLIENT_NAME"),
		zanzibarURL: "https://localhost:8443/api",
	}
}

func (zs ZanzibarService) sendRequest(method, endpoint string, body interface{}) (*http.Response, error) {
	url := fmt.Sprintf("%s%s", zs.zanzibarURL, endpoint)

	jsonData, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest(method, url, bytes.NewBuffer(jsonData))
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-API-KEY", zs.apiKey)
	req.Header.Set("Client-Name", zs.client)

	// Disable certificate verification
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}

	// Create HTTP client with custom transport
	client := &http.Client{Transport: tr}

	//client := &http.Client{}
	return client.Do(req)
}
