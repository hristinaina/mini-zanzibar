package services

import (
	"github.com/hashicorp/consul/api"
	"github.com/syndtr/goleveldb/leveldb"
	"mini-zanzibar/dtos"
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

// TODO: add second type of createKeyValue fn

func (acls *ACLService) isRelationSubset(relation dtos.Relation, actualRelation string) bool {
	relationConfig, err := acls.consulDBService.GetByNamespace(acls.extractNamespace(relation.Object))
	if err != nil {
		return false
	}
	if len(relationConfig.Relations[relation.Relation]) == 0 {
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
		if acls.contains(relationConfig.Relations[currentRelation], actualRelation) {
			return true
		}
		currentRelation = relationConfig.Relations[currentRelation][0]
	}
}

func (acls *ACLService) extractNamespace(object string) string {
	return strings.Split(object, ":")[0]
}

func (acls *ACLService) contains(array []string, item string) bool {
	for _, i := range array {
		if i == item {
			return true
		}
	}
	return false
}
