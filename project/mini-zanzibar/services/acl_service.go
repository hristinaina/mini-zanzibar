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
}

func NewACLService(levelDB *leveldb.DB, consuldDB *api.Client) IACLService {
	return &ACLService{levelDBService: NewLevelDBService(levelDB), consulDBService: NewConsulDBService(consuldDB)}
}

func (acls *ACLService) AddACL(relation dtos.Relation) error {
	namespace := strings.Split(relation.Object, ":")[0]
	namespaceObj, err := acls.consulDBService.GetByNamespace(namespace)
	if err != nil {
		return errs.CustomError{Code: 404, Message: "Namespace not found"}
	}
	if !Contains(acls.consulDBService.GetRelationsByNamespace(namespaceObj), relation.Relation) {
		return errs.CustomError{Code: 400, Message: "Invalid relation"}
	}
	key, value := acls.createKeyValue(relation)
	return acls.levelDBService.Add(key, value)
}

func (acls *ACLService) CheckACL(relation dtos.Relation) (bool, error) {
	key, _ := acls.createKeyValue(relation)
	actualRelation, err := acls.levelDBService.GetByKey(key)
	if err != nil {
		return false, err
	}
	if (relation.Relation == actualRelation) || acls.isRelationSubset(relation, actualRelation) {
		return true, nil
	}
	return false, nil
}

func (acls *ACLService) createKeyValue(relation dtos.Relation) (string, string) {
	return relation.Object + "+" + relation.User, relation.Relation
}

// isRelationSubset checks hierarchy of relations
func (acls *ACLService) isRelationSubset(relation dtos.Relation, actualRelation string) bool {
	relationConfig, err := acls.consulDBService.GetByNamespace(acls.extractNamespace(relation.Object))
	if err != nil || len(relationConfig.Relations[relation.Relation]) == 0 {
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
