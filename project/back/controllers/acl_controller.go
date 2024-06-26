package controllers

import (
	"io/ioutil"
	"net/http"

	"back/dtos"
	"back/services"

	"github.com/gin-gonic/gin"
)

type ACLController struct {
	service services.ACLService
}

func NewACLController() ACLController {
	return ACLController{
		service: services.NewACLService(),
	}
}

func (aclc ACLController) Add(c *gin.Context) {
	var relation dtos.Relation
	if err := c.ShouldBind(&relation); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Bad input data"})
		return
	}

	resp, err := aclc.service.AddRelation(relation)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to send request to Zanzibar"})
		return
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to read response from Zanzibar"})
		return
	}

	c.Data(resp.StatusCode, resp.Header.Get("Content-Type"), body)
}

func (aclc ACLController) Check(c *gin.Context) {
	var relation dtos.Relation
	if err := c.ShouldBind(&relation); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Bad input data"})
		return
	}

	allowed, err := aclc.service.CheckRelation(relation)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	response := gin.H{
		"allowed": allowed,
	}

	c.JSON(http.StatusOK, response)
}
