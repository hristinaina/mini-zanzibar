package controllers

import (
	"net/http"

	"back/dtos" // Assuming dtos package contains your DTOs like Relation
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

	err := aclc.service.AddRelation(relation)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to add relation in Zanzibar"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Relation successfully saved in Zanzibar"})
}

func (aclc ACLController) Check(c *gin.Context) {
	var relation dtos.Relation
	if err := c.ShouldBind(&relation); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Bad input data"})
		return
	}

	authorized, err := aclc.service.CheckRelation(relation)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to check relation in Zanzibar"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"authorized": authorized})
}
