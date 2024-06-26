package controllers

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/syndtr/goleveldb/leveldb"
	"mini-zanzibar/dtos"
	errs "mini-zanzibar/errors"
	"mini-zanzibar/services"
)

type LevelDBController struct {
	service    services.ILevelDBService
	logService *services.LogService
}

func NewLevelDBController(db *leveldb.DB, logService *services.LogService) LevelDBController {
	return LevelDBController{service: services.NewLevelDBService(db), logService: logService}
}

func (lc LevelDBController) Get(c *gin.Context) {
	lc.logService.Info("Processing Get request for leveldb")
	data, err := lc.service.GetAll()
	if err != nil {
		lc.logService.Error("Failed to get all data from leveldb: " + err.Error())
		errs.InternalServerError(c, err)
		return
	}
	lc.logService.Info("Successfully retrieved all data from leveldb")
	c.JSON(200, data)
}

func (lc LevelDBController) Post(c *gin.Context) {
	lc.logService.Info("Processing Post request for leveldb")
	var kv dtos.KeyValue

	if err := c.ShouldBindJSON(&kv); err != nil {
		lc.logService.Error("Bad input data for leveldb: " + err.Error())
		errs.BadRequestError(c, err)
		return
	}

	err := lc.service.Add(kv.Key, kv.Value)
	if err != nil {
		lc.logService.Error("Failed to add key-value pair to leveldb: " + err.Error())
		errs.InternalServerError(c, err)
		return
	}

	lc.logService.Info("Successfully saved key-value pair to leveldb")
	c.JSON(200, gin.H{
		"message": "Key-Value pair saved",
	})
}

func (lc LevelDBController) GetByKey(c *gin.Context) {
	key := c.Param("key")
	lc.logService.Info("Processing GetByKey request from leveldb for key: " + key)

	value, err := lc.service.GetByKey(key)
	if errors.Is(err, leveldb.ErrNotFound) {
		lc.logService.Error("Key not found: " + key)
		errs.NotFoundError(c)
		return
	} else if err != nil {
		lc.logService.Error("Failed to get value from leveldb for key " + key + ": " + err.Error())
		errs.InternalServerError(c, err)
		return
	}

	lc.logService.Info("Successfully retrieved value from leveldb for key: " + key)
	c.JSON(200, gin.H{
		"key":   key,
		"value": value,
	})
}

func (lc LevelDBController) Delete(c *gin.Context) {
	key := c.Param("key")
	lc.logService.Info("Processing Delete request from leveldb for key: " + key)

	err := lc.service.Delete(key)
	if err != nil {
		lc.logService.Error("Failed to delete key " + key + ": from leveldb" + err.Error())
		errs.NotFoundError(c)
		return
	}

	lc.logService.Info("Successfully deleted key from levekdb: " + key)
	c.JSON(200, gin.H{
		"message": "Key deleted",
	})
}
