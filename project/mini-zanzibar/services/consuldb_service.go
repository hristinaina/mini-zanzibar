package services

import (
	"github.com/hashicorp/consul/api"
	"mini-zanzibar/dtos"
)

type IConsulDBService interface {
	GetAll() (map[string]string, error)
	GetByNamespace(namespace string) (*api.KVPair, error)
	AddNamespace(json dtos.KeyValue) error
	DeleteNamespace(namespace string) error
}

type ConsulDBService struct {
	db *api.Client
}

func NewConsulDBService(db *api.Client) IConsulDBService {
	return &ConsulDBService{db}
}

func (cs *ConsulDBService) GetAll() (map[string]string, error) {
	pairs, _, err := cs.db.KV().List("", nil)
	if err != nil {
		return nil, err
	}

	data := make(map[string]string)
	for _, pair := range pairs {
		data[pair.Key] = string(pair.Value)
	}
	return data, nil
}

func (cs *ConsulDBService) GetByNamespace(namespace string) (*api.KVPair, error) {
	kvPair, _, err := cs.db.KV().Get(namespace, nil)
	return kvPair, err
}

func (cs *ConsulDBService) AddNamespace(kv dtos.KeyValue) error {
	kvPair := &api.KVPair{
		Key:   kv.Key,
		Value: []byte(kv.Value),
	}

	_, err := cs.db.KV().Put(kvPair, nil)
	return err
}

func (cs *ConsulDBService) DeleteNamespace(namespace string) error {
	_, err := cs.db.KV().Delete(namespace, nil)
	return err
}
