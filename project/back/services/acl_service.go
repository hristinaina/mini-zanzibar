package services

import (
	"back/dtos"
	"encoding/json"
	"fmt"
	"net/http"
)

type ACLService struct {
	zanzibarService ZanzibarService
}

func NewACLService() ACLService {
	return ACLService{zanzibarService: NewZanzibarService()}
}

func (acls ACLService) AddRelation(relation dtos.Relation) error {
	resp, err := acls.zanzibarService.sendRequest("POST", "/acl", relation)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("failed to add relation, status code: %d", resp.StatusCode)
	}

	return nil
}

func (acls ACLService) CheckRelation(relation dtos.Relation) (bool, error) {
	resp, err := acls.zanzibarService.sendRequest("PUT", "/acl", relation)
	if err != nil {
		return false, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return false, fmt.Errorf("failed to check relation, status code: %d", resp.StatusCode)
	}

	var authorized bool
	if err := json.NewDecoder(resp.Body).Decode(&authorized); err != nil {
		return false, err
	}

	return authorized, nil
}
