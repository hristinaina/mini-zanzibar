package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/hashicorp/consul/api"
	"mini-zanzibar/dtos"
)

type ConsulDBController struct {
	db *api.Client
}

func NewConsulDBController(db *api.Client) ConsulDBController {
	return ConsulDBController{db: db}
}

func (cc *ConsulDBController) Get(c *gin.Context) {
	pairs, _, err := cc.db.KV().List("", nil)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	data := make(map[string]string)
	for _, pair := range pairs {
		data[pair.Key] = string(pair.Value)
	}

	c.JSON(200, data)
}

func (cc *ConsulDBController) Post(c *gin.Context) {
	var kv dtos.KeyValue
	if err := c.ShouldBindJSON(&kv); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	kvPair := &api.KVPair{
		Key:   kv.Key,
		Value: []byte(kv.Value),
	}

	_, err := cc.db.KV().Put(kvPair, nil)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"message": "Key-Value pair saved"})

}

func (cc *ConsulDBController) GetByKey(c *gin.Context) {
	key := c.Param("key")
	kvPair, _, err := cc.db.KV().Get(key, nil)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	if kvPair == nil {
		c.JSON(404, gin.H{"error": "Key not found"})
		return
	}

	c.JSON(200, gin.H{"key": kvPair.Key, "value": string(kvPair.Value)})
}

func (cc *ConsulDBController) Delete(c *gin.Context) {
	key := c.Param("key")
	_, err := cc.db.KV().Delete(key, nil)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"message": "Key deleted"})
}
