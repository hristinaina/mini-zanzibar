package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/hashicorp/consul/api"
	"github.com/syndtr/goleveldb/leveldb"
	"mini-zanzibar/dtos"
	errs "mini-zanzibar/errors"
	"mini-zanzibar/services"
	"net/http"
)

type ACLController struct {
	service services.IACLService
}

func NewACLController(leveldb *leveldb.DB, consuldb *api.Client) ACLController {
	return ACLController{service: services.NewACLService(leveldb, consuldb)}
}

func (aclc ACLController) Add(c *gin.Context) {
	var relation dtos.Relation
	if err := c.ShouldBind(&relation); err != nil {
		errs.BadRequestError(c, err)
		return
	}

	err := aclc.service.AddACL(relation)
	if err != nil {
		errs.ReturnError(c, err)
		return
	}
	c.JSON(200, gin.H{"message": "Relation successfully saved"})
}

func (aclc ACLController) Check(c *gin.Context) {
	var relation dtos.Relation
	if err := c.ShouldBind(&relation); err != nil {
		errs.BadRequestError(c, err)
		return
	}
	authorized, err := aclc.service.CheckACL(relation)
	if err != nil {
		errs.InternalServerError(c, err)
		return
	}

	response := gin.H{
		"allowed": authorized,
	}

	c.JSON(http.StatusOK, response)
}
