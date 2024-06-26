package services

import (
	"encoding/json"
	"fmt"
	"github.com/hashicorp/consul/api"
	"golang.org/x/exp/maps"
	"log"
	"mini-zanzibar/dtos"
	"mini-zanzibar/errors"
	"strconv"
	"strings"
)

type IConsulDBService interface {
	GetAll() (map[string]string, error)
	GetByNamespace(namespace string) (dtos.Namespace, error)
	AddNamespace(namespaces dtos.Namespaces) error
	DeleteNamespace(namespace string) error
	GetRelationsByNamespace(namespace dtos.Namespace) []string
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
	fmt.Println(cs.getHighestVersion(namespace))
	fmt.Println("--------------------------------")
	key := cs.getHighestVersion(namespace)
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
		relations := cs.GetRelationsByNamespace(namespace)
		relatedRelations := cs.getRelatedRelationsByNamespace(namespace)
		if HasUniqueElements(relations, relatedRelations) {
			return errors.CustomError{Code: 400, Message: "Invalid relation"}
		}
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

func (cs *ConsulDBService) GetRelationsByNamespace(namespace dtos.Namespace) []string {
	return maps.Keys(namespace.Relations)
}

func (cs *ConsulDBService) getRelatedRelationsByNamespace(namespace dtos.Namespace) []string {
	values := maps.Values(namespace.Relations)
	var relatedRelations []string
	for _, value := range values {
		relatedRelations = append(relatedRelations, value...)
	}
	return relatedRelations
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

func (cs *ConsulDBService) getHighestVersion(namespace string) string {
	keyPattern := "namespace/" // Prefix for keys

	// Fetch all keys under the specified prefix
	kv := cs.db.KV()
	pairs, _, err := kv.Keys(keyPattern, "", nil)
	if err != nil {
		log.Fatalf("Failed to fetch keys from Consul: %v", err)
	}

	// Filter keys that match the pattern "namespace/v*/" + namespace
	var matchingKeys []string
	for _, key := range pairs {
		if strings.HasSuffix(key, "/"+namespace) {
			matchingKeys = append(matchingKeys, key)
		}
	}
	return cs.findHighestNumber(matchingKeys)
}

func (cs *ConsulDBService) findHighestNumber(pairs []string) string {
	highestVersion := 0
	var highestKey string

	// Iterate through keys to find highest version
	for _, key := range pairs {
		// Extract version from key
		parts := strings.Split(key, "/")
		versionStr := parts[len(parts)-2] // Assuming version is the second last part
		versionNumStr := versionStr[1:]
		version, err := strconv.Atoi(versionNumStr)
		if err != nil {
			continue // Skip keys with invalid version format
		}

		// Compare versions
		if version > highestVersion {
			highestVersion = version
			highestKey = key
		}
	}

	return highestKey
}
