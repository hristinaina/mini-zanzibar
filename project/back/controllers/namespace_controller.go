package controllers

import (
	"back/models"
	"io/ioutil"
	"net/http"

	"back/dtos"
	"back/services"

	"github.com/gin-gonic/gin"
)

type NSController struct {
	service    services.NSService
	logService *services.LogService
}

func NewNSController(logService *services.LogService) NSController {
	return NSController{
		service:    services.NewNSService(),
		logService: logService,
	}
}

func (cc *NSController) Get(c *gin.Context) {
	currentUserFromCookie, _ := c.Get("user")
	currentUser := currentUserFromCookie.(*models.User)

	cc.logService.Info("Processing Get All Namespaces request. User: " + currentUser.Email)
	resp, err := cc.service.GetAll()
	if err != nil {
		cc.logService.Error("Failed to send request to Zanzibar in Get All Namespaces")
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to send request to Zanzibar"})
		return
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		cc.logService.Error("Failed to read response from Zanzibar in Get All Namespaces")
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to read response from Zanzibar"})
		return
	}

	cc.logService.Info("Get All Namespaces request processed successfully. User: " + currentUser.Email)
	c.Data(resp.StatusCode, resp.Header.Get("Content-Type"), body)
}

func (cc *NSController) AddNamespace(c *gin.Context) {
	currentUserFromCookie, _ := c.Get("user")
	currentUser := currentUserFromCookie.(*models.User)

	cc.logService.Info("Processing AddNamespace request for namespaces. User: " + currentUser.Email)
	var namespaces dtos.Namespaces
	if err := c.BindJSON(&namespaces); err != nil {
		cc.logService.Error("Bad input data in AddNamespace")
		c.JSON(http.StatusBadRequest, gin.H{"error": "Bad input data"})
		return
	}

	resp, err := cc.service.AddNamespace(namespaces)
	if err != nil {
		cc.logService.Error("Failed to send request to Zanzibar in AddNamespace for namespaces. User: " + currentUser.Email)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to send request to Zanzibar"})
		return
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		cc.logService.Error("Failed to read response from Zanzibar in AddNamespace for namespaces.")
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to read response from Zanzibar"})
		return
	}

	cc.logService.Info("AddNamespace request processed successfully for namespaces. User: " + currentUser.Email)
	c.Data(resp.StatusCode, resp.Header.Get("Content-Type"), body)
}

func (cc *NSController) GetByNamespace(c *gin.Context) {
	key := c.Param("key")

	cc.logService.Info("Processing GetByNamespace request for key: " + key)
	resp, err := cc.service.GetByNamespace(key)
	if err != nil {
		cc.logService.Error("Failed to send request to Zanzibar in GetByNamespace for key: " + key)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to send request to Zanzibar"})
		return
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		cc.logService.Error("Failed to read response from Zanzibar in GetByNamespace for key: " + key)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to read response from Zanzibar"})
		return
	}

	cc.logService.Info("GetByNamespace request processed successfully for key: " + key)
	c.Data(resp.StatusCode, resp.Header.Get("Content-Type"), body)
}

func (cc *NSController) Delete(c *gin.Context) {
	key := c.Param("key")

	cc.logService.Info("Processing DeleteNamespace request for key: " + key)
	resp, err := cc.service.DeleteNamespace(key)
	if err != nil {
		cc.logService.Error("Failed to send request to Zanzibar in DeleteNamespace for key: " + key)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to send request to Zanzibar"})
		return
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		cc.logService.Error("Failed to read response from Zanzibar in DeleteNamespace for key: " + key)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to read response from Zanzibar"})
		return
	}

	cc.logService.Info("DeleteNamespace request processed successfully for key: " + key)
	c.Data(resp.StatusCode, resp.Header.Get("Content-Type"), body)
}
