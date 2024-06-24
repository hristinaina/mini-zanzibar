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

func (cc *ConsulDBController) Post(c *gin.Context) {
	var kv dtos.KeyValue
	if err := c.ShouldBindJSON(&kv); err != nil {
		errs.InternalServerError(c, err)
		return
	}
	err := cc.service.AddNamespace(kv)
	if err != nil {
		errs.InternalServerError(c, err)
		return
	}

	c.JSON(200, gin.H{"message": "Key-Value pair saved"})

}

func (cc *ConsulDBController) GetByKey(c *gin.Context) {
	key := c.Param("key")
	kvPair, err := cc.service.GetByNamespace(key)
	if kvPair == nil {
		errs.KeyNotFoundError(c)
		return
	}
	if err != nil {
		errs.InternalServerError(c, err)
		return
	}

	c.JSON(200, gin.H{"key": kvPair.Key, "value": string(kvPair.Value)})
}

func (cc *ConsulDBController) Delete(c *gin.Context) {
	key := c.Param("key")
	err := cc.service.DeleteNamespace(key)
	if err != nil {
		errs.InternalServerError(c, err)
		return
	}

	c.JSON(200, gin.H{"message": "Key deleted"})
}
