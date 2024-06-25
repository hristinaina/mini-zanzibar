package services

import (
	"back/dtos"
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

func (acls ACLService) CheckRelation(relation dtos.Relation) (*http.Response, error) {
	return acls.zanzibarService.sendRequest("PUT", "/acl", relation)
}
