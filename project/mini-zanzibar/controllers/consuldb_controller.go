package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/hashicorp/consul/api"
	"mini-zanzibar/dtos"
	errs "mini-zanzibar/errors"
	"mini-zanzibar/services"
)

type ConsulDBController struct {
	db      *api.Client
	service services.IConsulDBService
}

func NewConsulDBController(db *api.Client) ConsulDBController {
	return ConsulDBController{db: db, service: services.NewConsulDBService(db)}
}

func (cc *ConsulDBController) Get(c *gin.Context) {
	data, err := cc.service.GetAll()
	if err != nil {
		errs.InternalServerError(c, err)
		return
	}
	c.JSON(200, data)
}

func (cc *ConsulDBController) AddNamespace(c *gin.Context) {
	var namespaces dtos.Namespaces
	if err := c.BindJSON(&namespaces); err != nil {
		errs.BadRequestError(c, err)
		return
	}
	err := cc.service.AddNamespace(namespaces)
	if err != nil {
		errs.InternalServerError(c, err)
	}

	c.JSON(200, gin.H{"message": "Key-Value pair saved"})
}

func (cc *ConsulDBController) GetByNamespace(c *gin.Context) {
	key := c.Param("key")
	namespace, err := cc.service.GetByNamespace(key)
	if err != nil {
		errs.KeyNotFoundError(c)
		return
	}

	c.JSON(200, namespace)
}

func (cc *ConsulDBController) Delete(c *gin.Context) {
	key := c.Param("key")
	err := cc.service.DeleteNamespace(key)
	if err != nil {
		errs.KeyNotFoundError(c)
		return
	}

	c.JSON(200, gin.H{"message": "Key deleted"})
}
