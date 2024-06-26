package services

import (
	"github.com/hashicorp/consul/api"
	"github.com/syndtr/goleveldb/leveldb"
	"mini-zanzibar/dtos"
	errs "mini-zanzibar/errors"
	"strings"
)

type IACLService interface {
	AddACL(relation dtos.Relation) error
	CheckACL(relation dtos.Relation) (bool, error)
}

type ACLService struct {
	levelDBService  ILevelDBService
	consulDBService IConsulDBService
	logService      *LogService
}

func NewACLService(levelDB *leveldb.DB, consuldDB *api.Client, logService *LogService) IACLService {
	return &ACLService{levelDBService: NewLevelDBService(levelDB), consulDBService: NewConsulDBService(consuldDB), logService: logService}
}

func (acls *ACLService) AddACL(relation dtos.Relation) error {
	acls.logService.Info("Adding ACL for relation: " + relation.Relation)

	namespace := strings.Split(relation.Object, ":")[0]
	namespaceObj, err := acls.consulDBService.GetByNamespace(namespace)
	if err != nil {
		acls.logService.Error("Namespace not found: " + namespace)
		return errs.CustomError{Code: 404, Message: "Namespace not found"}
	}

	if !Contains(acls.consulDBService.GetRelationsByNamespace(namespaceObj), relation.Relation) {
		acls.logService.Error("Invalid relation: " + relation.Relation)
		return errs.CustomError{Code: 400, Message: "Invalid relation"}
	}

	key, value := acls.createKeyValue(relation)
	err = acls.levelDBService.Add(key, value)
	if err != nil {
		acls.logService.Error("Failed to add key-value pair: " + err.Error())
		return err
	}

	acls.logService.Info("Successfully added ACL for relation: " + relation.Relation)
	return nil
}

func (acls *ACLService) CheckACL(relation dtos.Relation) (bool, error) {
	acls.logService.Info("Checking ACL for relation: " + relation.Relation)

	key, _ := acls.createKeyValue(relation)
	actualRelation, err := acls.levelDBService.GetByKey(key)
	if err != nil {
		acls.logService.Error("Failed to get value for key: " + key + " with error: " + err.Error())
		return false, err
	}

	if relation.Relation == actualRelation || acls.isRelationSubset(relation, actualRelation) {
		acls.logService.Info("Relation is authorized: " + relation.Relation)
		return true, nil
	}

	acls.logService.Info("Relation is not authorized: " + relation.Relation)
	return false, nil
}

func (acls *ACLService) createKeyValue(relation dtos.Relation) (string, string) {
	return relation.Object + "+" + relation.User, relation.Relation
}

// isRelationSubset checks hierarchy of relations
func (acls *ACLService) isRelationSubset(relation dtos.Relation, actualRelation string) bool {
	relationConfig, err := acls.consulDBService.GetByNamespace(acls.extractNamespace(relation.Object))
	if err != nil || len(relationConfig.Relations[relation.Relation]) == 0 {
		acls.logService.Error("Failed to get relation config or no relations found for: " + relation.Relation)
		return false
	}

	var currentRelation = relationConfig.Relations[relation.Relation][0]
	if currentRelation == actualRelation {
		return true
	}

	for {
		if len(relationConfig.Relations[currentRelation]) == 0 {
			return false
		}
		if Contains(relationConfig.Relations[currentRelation], actualRelation) {
			return true
		}
		currentRelation = relationConfig.Relations[currentRelation][0]
	}
}

func (acls *ACLService) extractNamespace(object string) string {
	return strings.Split(object, ":")[0]
}
