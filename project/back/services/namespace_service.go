package services

import (
	"back/dtos"
	"net/http"
)

type NSService struct {
	zanzibarService ZanzibarService
}

func NewNSService() NSService {
	return NSService{zanzibarService: NewZanzibarService()}
}

func (nss NSService) GetAll() (*http.Response, error) {
	return nss.zanzibarService.sendRequest("GET", "/consuldb/all", nil)
}

func (nss NSService) AddNamespace(namespace dtos.Namespace) (*http.Response, error) {
	return nss.zanzibarService.sendRequest("POST", "/consuldb", namespace)
}

func (nss NSService) GetByNamespace(key string) (*http.Response, error) {
	return nss.zanzibarService.sendRequest("GET", "/consuldb/"+key, nil)
}

func (nss NSService) DeleteNamespace(key string) (*http.Response, error) {
	return nss.zanzibarService.sendRequest("DELETE", "/consuldb/"+key, nil)
}
