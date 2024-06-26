package controllers

import (
	"mini-zanzibar/dtos"
	errs "mini-zanzibar/errors"
	"mini-zanzibar/services"

	"github.com/gin-gonic/gin"
	"github.com/hashicorp/consul/api"
)

type ConsulDBController struct {
	service    services.IConsulDBService
	logService *services.LogService
}

func NewConsulDBController(db *api.Client, logService *services.LogService) ConsulDBController {
	return ConsulDBController{service: services.NewConsulDBService(db), logService: logService}
}

func (cc *ConsulDBController) Get(c *gin.Context) {
	cc.logService.Info("Processing Get request for namespaces")
	data, err := cc.service.GetAll()
	if err != nil {
		cc.logService.Error("Failed to get all data for namespaces: " + err.Error())
		errs.InternalServerError(c, err)
		return
	}
	cc.logService.Info("Successfully retrieved all namespaces data")
	c.JSON(200, data)
}

func (cc *ConsulDBController) AddNamespace(c *gin.Context) {
	cc.logService.Info("Processing AddNamespace request")
	var namespaces dtos.Namespaces
	if err := c.BindJSON(&namespaces); err != nil {
		cc.logService.Error("Bad input data for namespace: " + err.Error())
		errs.BadRequestError(c, err)
		return
	}
	err := cc.service.AddNamespace(namespaces)
	if err != nil {
		cc.logService.Error("Failed to add namespace: " + err.Error())
		errs.ReturnError(c, err)
		return
	}
	cc.logService.Info("Successfully added namespace")
	c.JSON(200, gin.H{"message": "Key-Value pair saved"})
}

func (cc *ConsulDBController) GetByNamespace(c *gin.Context) {
	key := c.Param("key")
	cc.logService.Info("Processing GetByNamespace request for key: " + key)
	namespace, err := cc.service.GetByNamespace(key)
	if err != nil {
		cc.logService.Error("Failed to get namespace for key " + key + ": " + err.Error())
		errs.ReturnError(c, err)
		return
	}
	cc.logService.Info("Successfully retrieved namespace for key: " + key)
	c.JSON(200, namespace)
}

func (cc *ConsulDBController) Delete(c *gin.Context) {
	key := c.Param("key")
	cc.logService.Info("Processing Delete namespace request for key: " + key)
	err := cc.service.DeleteNamespace(key)
	if err != nil {
		cc.logService.Error("Failed to delete namespace key " + key + ": " + err.Error())
		errs.NotFoundError(c)
		return
	}
	cc.logService.Info("Successfully deleted namespace key: " + key)
	c.JSON(200, gin.H{"message": "Key deleted"})
}
