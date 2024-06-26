package controllers

import (
	"io/ioutil"
	"net/http"

	"back/dtos"
	"back/services"

	"github.com/gin-gonic/gin"
)

type DataController struct {
	service    services.DataService
	logService *services.LogService
}

func NewDataController(logService *services.LogService) DataController {
	return DataController{
		service:    services.NewDataService(),
		logService: logService,
	}
}

func (dc DataController) GetAll(c *gin.Context) {
	dc.logService.Info("Processing GetAll request")
	resp, err := dc.service.GetAll()
	if err != nil {
		dc.logService.Error("Failed to send request to Zanzibar in GetAll")
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to send request to Zanzibar"})
		return
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		dc.logService.Error("Failed to read response from Zanzibar in GetAll")
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to read response from Zanzibar"})
		return
	}

	dc.logService.Info("GetAll request processed successfully")
	c.Data(resp.StatusCode, resp.Header.Get("Content-Type"), body)
}

func (dc DataController) GetByKey(c *gin.Context) {
	key := c.Param("key")
	dc.logService.Info("Processing GetByKey request for key: " + key)

	resp, err := dc.service.GetByKey(key)
	if err != nil {
		dc.logService.Error("Failed to send request to Zanzibar in GetByKey for key: " + key)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to send request to Zanzibar"})
		return
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		dc.logService.Error("Failed to read response from Zanzibar in GetByKey for key: " + key)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to read response from Zanzibar"})
		return
	}

	dc.logService.Info("GetByKey request processed successfully for key: " + key)
	c.Data(resp.StatusCode, resp.Header.Get("Content-Type"), body)
}

func (dc DataController) Add(c *gin.Context) {
	var kv dtos.KeyValue
	if err := c.ShouldBindJSON(&kv); err != nil {
		dc.logService.Error("Bad input data in Add")
		c.JSON(http.StatusBadRequest, gin.H{"error": "Bad input data"})
		return
	}

	dc.logService.Info("Processing Add request for key: " + kv.Key)
	resp, err := dc.service.Add(kv.Key, kv.Value)
	if err != nil {
		dc.logService.Error("Failed to send request to Zanzibar in Add for key: " + kv.Key)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to send request to Zanzibar"})
		return
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		dc.logService.Error("Failed to read response from Zanzibar in Add for key: " + kv.Key)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to read response from Zanzibar"})
		return
	}

	dc.logService.Info("Add request processed successfully for key: " + kv.Key)
	c.Data(resp.StatusCode, resp.Header.Get("Content-Type"), body)
}

func (dc DataController) Delete(c *gin.Context) {
	key := c.Param("key")
	dc.logService.Info("Processing Delete request for key: " + key)

	resp, err := dc.service.Delete(key)
	if err != nil {
		dc.logService.Error("Failed to send request to Zanzibar in Delete for key: " + key)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to send request to Zanzibar"})
		return
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		dc.logService.Error("Failed to read response from Zanzibar in Delete for key: " + key)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to read response from Zanzibar"})
		return
	}

	dc.logService.Info("Delete request processed successfully for key: " + key)
	c.Data(resp.StatusCode, resp.Header.Get("Content-Type"), body)
}
