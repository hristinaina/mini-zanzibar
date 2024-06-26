package services

import (
	"back/dtos"
	"encoding/json"
	"errors"
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
		return false, errors.New("Failed to check relation: received non-200 status code")
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return false, errors.New("Failed to read response from Zanzibar")
	}

	var result struct {
		Allowed bool `json:"allowed"`
	}

	if err := json.Unmarshal(body, &result); err != nil {
		return false, errors.New("Failed to decode response from Zanzibar")
	}

	return result.Allowed, nil
}
