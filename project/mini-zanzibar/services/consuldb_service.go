package services

import (
	"encoding/json"
	"github.com/hashicorp/consul/api"
	"mini-zanzibar/dtos"
	"mini-zanzibar/errors"
)

type IConsulDBService interface {
	GetAll() (map[string]string, error)
	GetByNamespace(namespace string) (dtos.Namespace, error)
	AddNamespace(namespaces dtos.Namespaces) error
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

func (cs *ConsulDBService) GetByNamespace(namespace string) (dtos.Namespace, error) {
	key := "namespace/" + namespace
	kvPair, _, err := cs.db.KV().Get(key, nil)
	if err != nil {
		return dtos.Namespace{}, err
	}
	if kvPair == nil {
		return dtos.Namespace{}, errors.CustomError{Code: 404, Message: "Namespace not found."}
	}
	var namespaceData dtos.Namespace
	if err := json.Unmarshal(kvPair.Value, &namespaceData); err != nil {
		return dtos.Namespace{}, err
	}
	return namespaceData, err
}

func (cs *ConsulDBService) AddNamespace(namespaces dtos.Namespaces) error {
	for _, namespace := range namespaces.Namespaces {
		if cs.isCyclicGraph(namespace) {
			return errors.CustomError{Code: 400, Message: "Invalid configuration"}
		}
		key := "namespace/" + namespace.Namespace
		value, err := json.Marshal(namespace)
		if err != nil {
			return err
		}

		kv := cs.db.KV()
		p := &api.KVPair{Key: key, Value: value}
		_, err = kv.Put(p, nil)
		if err != nil {
			return err
		}
	}
	return nil
}

func (cs *ConsulDBService) DeleteNamespace(namespace string) error {
	_, err := cs.GetByNamespace(namespace)
	if err != nil {
		return err
	}
	key := "namespace/" + namespace
	_, err = cs.db.KV().Delete(key, nil)
	return err
}

func (cs *ConsulDBService) isCyclicGraph(namespace dtos.Namespace) bool {
	for key, value := range namespace.Relations {
		if len(value) == 0 {
			continue
		}
		if Contains(namespace.Relations[value[0]], key) {
			return true
		}
		if len(namespace.Relations[value[0]]) == 0 {
			continue
		}
		currentRelation := namespace.Relations[value[0]][0]
		for {
			if len(namespace.Relations[currentRelation]) == 0 {
				break
			}
			if Contains(namespace.Relations[currentRelation], key) {
				return true
			}
			currentRelation = namespace.Relations[currentRelation][0]
		}
	}
	return false
}
