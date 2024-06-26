package services

import (
	"back/dtos"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
)

type ACLService struct {
	zanzibarService ZanzibarService
}

func NewACLService() ACLService {
	return ACLService{zanzibarService: NewZanzibarService()}
}

func (acls ACLService) AddRelation(relation dtos.Relation) (*http.Response, error) {
	return acls.zanzibarService.sendRequest("POST", "/acl", relation)
}

func (acls ACLService) CheckRelation(relation dtos.Relation) (bool, error) {
	resp, err := acls.zanzibarService.sendRequest("PUT", "/acl", relation)
	if err != nil {
		return false, errors.New("Failed to send request to Zanzibar")
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		// Read the response body
		bodyBytes, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return false, errors.New("Failed to read response body")
		}

		// Convert body to string
		bodyString := string(bodyBytes)

		// Create an error message with the response status and body
		return false, fmt.Errorf("status %d, response %s", resp.StatusCode, bodyString)
	}

	var result struct {
		Allowed bool `json:"allowed"`
	}
	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		return false, errors.New("Failed to decode response body")
	}
	return result.Allowed, nil
}
