package services

import (
	"github.com/hashicorp/consul/api"
	"github.com/syndtr/goleveldb/leveldb"
	"mini-zanzibar/dtos"
)

type IACLService interface {
	AddACL(relation dtos.Relation) error
	CheckACL(relation dtos.Relation) bool
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

func (acls *ACLService) CheckACL(relation dtos.Relation) bool {
	return false
}

func (acls *ACLService) createKeyValue(relation dtos.Relation) (string, string) {
	return relation.Object + "+" + relation.User, relation.Relation
}
