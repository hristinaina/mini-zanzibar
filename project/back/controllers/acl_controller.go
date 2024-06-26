package controllers

import (
	"io/ioutil"
	"net/http"

	"back/dtos"
	"back/services"

	"github.com/gin-gonic/gin"
)

type ACLController struct {
	service    services.ACLService
	logService *services.LogService
}

func NewACLController(logService *services.LogService) ACLController {
	return ACLController{
		service:    services.NewACLService(),
		logService: logService,
	}
}

func (aclc ACLController) Add(c *gin.Context) {
	var relation dtos.Relation
	if err := c.ShouldBind(&relation); err != nil {
		aclc.logService.Error("Bad input data when Adding an ACL")
		c.JSON(http.StatusBadRequest, gin.H{"error": "Bad input data"})
		return
	}

	aclc.logService.Info("Processing Add request for relation: " + relation.Object)
	resp, err := aclc.service.AddRelation(relation)
	if err != nil {
		aclc.logService.Error("Failed to send request to Zanzibar in Add for relation: " + relation.Object)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to send request to Zanzibar"})
		return
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		aclc.logService.Error("Failed to read response from Zanzibar in Add for relation: " + relation.Object)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to read response from Zanzibar"})
		return
	}

	aclc.logService.Info("Add request processed successfully for relation: " + relation.Object)
	c.Data(resp.StatusCode, resp.Header.Get("Content-Type"), body)
}

func (aclc ACLController) Check(c *gin.Context) {
	var relation dtos.Relation
	if err := c.ShouldBind(&relation); err != nil {
		aclc.logService.Error("Bad input data in Check")
		c.JSON(http.StatusBadRequest, gin.H{"error": "Bad input data"})
		return
	}

	aclc.logService.Info("Processing Check request for relation: " + relation.Object)
	allowed, err := aclc.service.CheckRelation(relation)
	if err != nil {
		aclc.logService.Error("Failed to check relation in Check for relation: " + relation.Object)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	response := gin.H{
		"allowed": allowed,
	}

	aclc.logService.Info("Check request processed successfully for relation: " + relation.Object)
	c.JSON(http.StatusOK, response)
}
