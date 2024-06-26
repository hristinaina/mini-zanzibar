package controllers

import (
	"io/ioutil"
	"net/http"

	"back/dtos"
	"back/services"

	"github.com/gin-gonic/gin"
)

type NSController struct {
	service services.NSService
}

func NewNSController() NSController {
	return NSController{
		service: services.NewNSService(),
	}
}

func (cc *NSController) Get(c *gin.Context) {
	resp, err := cc.service.GetAll()
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

func (cc *NSController) AddNamespace(c *gin.Context) {
	var namespaces dtos.Namespaces
	if err := c.BindJSON(&namespaces); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Bad input data"})
		return
	}

	resp, err := cc.service.AddNamespace(namespaces)
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

func (cc *NSController) GetByNamespace(c *gin.Context) {
	key := c.Param("key")

	resp, err := cc.service.GetByNamespace(key)
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

func (cc *NSController) Delete(c *gin.Context) {
	key := c.Param("key")

	resp, err := cc.service.DeleteNamespace(key)
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
