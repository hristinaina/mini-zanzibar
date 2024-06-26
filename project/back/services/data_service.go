package services

import (
	"back/dtos"
	"net/http"
)

type DataService struct {
	zanzibarService ZanzibarService
}

func NewDataService() DataService {
	return DataService{zanzibarService: NewZanzibarService()}
}

func (ds DataService) GetAll() (*http.Response, error) {
	return ds.zanzibarService.sendRequest("GET", "/leveldb/all", nil)
}

func (ds DataService) GetByKey(key string) (*http.Response, error) {
	return ds.zanzibarService.sendRequest("GET", "/leveldb/"+key, nil)
}

func (ds DataService) Add(key, value string) (*http.Response, error) {
	kv := dtos.KeyValue{Key: key, Value: value}
	return ds.zanzibarService.sendRequest("POST", "/leveldb", kv)
}

func (ds DataService) Delete(key string) (*http.Response, error) {
	return ds.zanzibarService.sendRequest("DELETE", "/leveldb/"+key, nil)
}
