package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/hashicorp/consul/api"
)

type ConsulDBController struct {
	db *api.Client
}

func NewConsulDBController(db *api.Client) ConsulDBController {
	return ConsulDBController{db: db}
}

func (*ConsulDBController) Get(c *gin.Context) {
	c.JSON(200, "OK")
}

func (*ConsulDBController) Post(c *gin.Context) {

}
