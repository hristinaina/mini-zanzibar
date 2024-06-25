package controllers

import (
	"io/ioutil"
	"net/http"

	"back/dtos"
	"back/services"

	"github.com/gin-gonic/gin"
)

type DataController struct {
	service services.DataService
}

func NewDataController() DataController {
	return DataController{
		service: services.NewDataService(),
	}
}

func (dc DataController) GetAll(c *gin.Context) {
	resp, err := dc.service.GetAll()
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

func (dc DataController) GetByKey(c *gin.Context) {
	key := c.Param("key")

	resp, err := dc.service.GetByKey(key)
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

func (dc DataController) Add(c *gin.Context) {
	var kv dtos.KeyValue
	if err := c.ShouldBindJSON(&kv); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Bad input data"})
		return
	}

	resp, err := dc.service.Add(kv.Key, kv.Value)
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

func (dc DataController) Delete(c *gin.Context) {
	key := c.Param("key")

	resp, err := dc.service.Delete(key)
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
