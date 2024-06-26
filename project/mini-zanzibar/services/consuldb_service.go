package services

import (
	"encoding/json"
	"log"
	"mini-zanzibar/dtos"
	"mini-zanzibar/errors"
	"strconv"
	"strings"

	"github.com/hashicorp/consul/api"
	"golang.org/x/exp/maps"
)

type IConsulDBService interface {
	GetAll() (map[string]string, error)
	GetByNamespace(namespace string) (dtos.Namespace, error)
	AddNamespace(namespaces dtos.Namespaces) error
	DeleteNamespace(namespace string) error
	GetRelationsByNamespace(namespace dtos.Namespace) []string
}

type ConsulDBService struct {
	db         *api.Client
	logService *LogService
}

func NewConsulDBService(db *api.Client, logService *LogService) IConsulDBService {
	return &ConsulDBService{db, logService}
}

func (cs *ConsulDBService) GetAll() (map[string]string, error) {
	cs.logService.Info("Fetching all key-value pairs from Consul")

	pairs, _, err := cs.db.KV().List("", nil)
	if err != nil {
		cs.logService.Error("Failed to fetch all key-value pairs: " + err.Error())
		return nil, err
	}

	data := make(map[string]string)
	for _, pair := range pairs {
		data[pair.Key] = string(pair.Value)
	}

	cs.logService.Info("Successfully fetched all key-value pairs from Consul")
	return data, nil
}

func (cs *ConsulDBService) GetByNamespace(namespace string) (dtos.Namespace, error) {
	key, _ := cs.getHighestVersion(namespace)
	cs.logService.Info("Fetching namespace data for: " + namespace)
	kvPair, _, err := cs.db.KV().Get(key, nil)
	if err != nil {
		cs.logService.Error("Failed to fetch namespace data for: " + namespace + ", Error: " + err.Error())
		return dtos.Namespace{}, err
	}
	if kvPair == nil {
		err := errors.CustomError{Code: 404, Message: "Namespace not found."}
		cs.logService.Error("Namespace not found: " + namespace)
		return dtos.Namespace{}, err
	}

	var namespaceData dtos.Namespace
	if err := json.Unmarshal(kvPair.Value, &namespaceData); err != nil {
		cs.logService.Error("Failed to unmarshal namespace data for: " + namespace + ", Error: " + err.Error())
		return dtos.Namespace{}, err
	}

	cs.logService.Info("Successfully fetched namespace data for: " + namespace)
	return namespaceData, nil
}

func (cs *ConsulDBService) AddNamespace(namespaces dtos.Namespaces) error {
	cs.logService.Info("Adding namespaces")

	for _, namespace := range namespaces.Namespaces {
		cs.logService.Info("Adding namespace: " + namespace.Namespace)

		relations := cs.GetRelationsByNamespace(namespace)
		relatedRelations := cs.getRelatedRelationsByNamespace(namespace)
		if HasUniqueElements(relations, relatedRelations) {
			err := errors.CustomError{Code: 400, Message: "Invalid relation"}
			cs.logService.Error("Invalid relation for namespace: " + namespace.Namespace)
			return err
		}

		if cs.isCyclicGraph(namespace) {
			err := errors.CustomError{Code: 400, Message: "Invalid configuration"}
			cs.logService.Error("Cyclic graph detected in namespace: " + namespace.Namespace)
			return err
		}
		_, version := cs.getHighestVersion(namespace.Namespace)
		version += 1
		key := "namespace/v" + strconv.Itoa(version) + "/" + namespace.Namespace
		value, err := json.Marshal(namespace)
		if err != nil {
			cs.logService.Error("Failed to marshal namespace data for: " + namespace.Namespace + ", Error: " + err.Error())
			return err
		}

		kv := cs.db.KV()
		p := &api.KVPair{Key: key, Value: value}
		_, err = kv.Put(p, nil)
		if err != nil {
			cs.logService.Error("Failed to put namespace data for: " + namespace.Namespace + ", Error: " + err.Error())
			return err
		}

		cs.logService.Info("Successfully added namespace: " + namespace.Namespace)
	}

	return nil
}

func (cs *ConsulDBService) DeleteNamespace(namespace string) error {
	cs.logService.Info("Deleting namespace: " + namespace)

	_, err := cs.GetByNamespace(namespace)
	if err != nil {
		cs.logService.Error("Failed to delete namespace: " + namespace + ", Error: Namespace not found")
		return err
	}

	key := "namespace/" + namespace
	_, err = cs.db.KV().Delete(key, nil)
	if err != nil {
		cs.logService.Error("Failed to delete namespace: " + namespace + ", Error: " + err.Error())
		return err
	}

	cs.logService.Info("Successfully deleted namespace: " + namespace)
	return nil
}

func (cs *ConsulDBService) GetRelationsByNamespace(namespace dtos.Namespace) []string {
	cs.logService.Info("Fetching relations for namespace: " + namespace.Namespace)

	return maps.Keys(namespace.Relations)
}

func (cs *ConsulDBService) getRelatedRelationsByNamespace(namespace dtos.Namespace) []string {
	cs.logService.Info("Fetching related relations for namespace: " + namespace.Namespace)

	values := maps.Values(namespace.Relations)
	var relatedRelations []string
	for _, value := range values {
		relatedRelations = append(relatedRelations, value...)
	}
	return relatedRelations
}

func (cs *ConsulDBService) isCyclicGraph(namespace dtos.Namespace) bool {
	cs.logService.Info("Checking cyclic graph for namespace: " + namespace.Namespace)

	for key, value := range namespace.Relations {
		if len(value) == 0 {
			continue
		}

		if Contains(namespace.Relations[value[0]], key) {
			cs.logService.Error("Cyclic graph detected in namespace: " + namespace.Namespace)
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
				cs.logService.Error("Cyclic graph detected in namespace: " + namespace.Namespace)
				return true
			}
			currentRelation = namespace.Relations[currentRelation][0]
		}
	}

	cs.logService.Info("No cyclic graph detected in namespace: " + namespace.Namespace)
	return false
}

func (cs *ConsulDBService) getHighestVersion(namespace string) (string, int) {
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

func (cs *ConsulDBService) findHighestNumber(pairs []string) (string, int) {
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

	return highestKey, highestVersion
}
