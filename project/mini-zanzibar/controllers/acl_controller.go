package controllers

import (
	"mini-zanzibar/dtos"
	errs "mini-zanzibar/errors"
	"mini-zanzibar/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/hashicorp/consul/api"
	"github.com/syndtr/goleveldb/leveldb"
)

type ACLController struct {
	service    services.IACLService
	logService *services.LogService
}

func NewACLController(leveldb *leveldb.DB, consuldb *api.Client, logService *services.LogService) ACLController {
	return ACLController{service: services.NewACLService(leveldb, consuldb, logService), logService: logService}
}

func (aclc ACLController) Add(c *gin.Context) {
	var relation dtos.Relation
	if err := c.ShouldBind(&relation); err != nil {
		aclc.logService.Error("Failed to bind relation data in Add method: " + err.Error())
		errs.BadRequestError(c, err)
		return
	}

	aclc.logService.Info("Processing Add relation request for relation: " + relation.Relation)
	err := aclc.service.AddACL(relation)
	if err != nil {
		aclc.logService.Error("Failed to add relation in Add method: " + err.Error())
		errs.ReturnError(c, err)
		return
	}
	aclc.logService.Info("Relation successfully saved: " + relation.Relation)
	c.JSON(http.StatusOK, gin.H{"message": "Relation successfully saved"})
}

func (aclc ACLController) Check(c *gin.Context) {
	var relation dtos.Relation
	if err := c.ShouldBind(&relation); err != nil {
		aclc.logService.Error("Failed to bind relation data in Check method: " + err.Error())
		errs.BadRequestError(c, err)
		return
	}

	aclc.logService.Info("Processing Check relation request for relation: " + relation.Relation)
	authorized, err := aclc.service.CheckACL(relation)
	if err != nil {
		aclc.logService.Error("Failed to check relation in Check method: " + err.Error())
		errs.InternalServerError(c, err)
		return
	}

	response := gin.H{
		"allowed": authorized,
	}

	aclc.logService.Info("Relation check result for " + relation.Relation + ": " + strconv.FormatBool(authorized))
	c.JSON(http.StatusOK, response)
}
