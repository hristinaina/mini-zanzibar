package controllers

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/syndtr/goleveldb/leveldb"
	"mini-zanzibar/dtos"
	"mini-zanzibar/services"
	"net/http"
)

type LevelDBController struct {
	db      *leveldb.DB
	service services.ILevelDBService
}

func NewLevelDBController(db *leveldb.DB) LevelDBController {
	return LevelDBController{db: db, service: services.NewLevelDBService(db)}
}

func (lc LevelDBController) Get(c *gin.Context) {
	data, err := lc.service.GetAll()
	if err != nil {
		ErrorJSON(c, err)
		return
	}
	c.JSON(200, data)
}

func (lc LevelDBController) Post(c *gin.Context) {
	var kv dtos.KeyValue

	if err := c.ShouldBindJSON(&kv); err != nil {
		ErrorJSON(c, err)
		return
	}

	err := lc.service.Add(kv.Key, kv.Value)
	if err != nil {
		ErrorJSON(c, err)
		return
	}

	c.JSON(200, gin.H{
		"message": "Key-Value pair saved",
	})
}

func (lc LevelDBController) GetByKey(c *gin.Context) {
	key := c.Param("key")

	value, err := lc.service.GetByKey(key)
	if errors.Is(err, leveldb.ErrNotFound) {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Key not found",
		})
		return
	} else if err != nil {
		ErrorJSON(c, err)
		return
	}

	c.JSON(200, gin.H{
		"key":   key,
		"value": string(value),
	})
}

func (lc LevelDBController) Delete(c *gin.Context) {
	key := c.Param("key")

	err := lc.service.Delete(key)
	if err != nil {
		ErrorJSON(c, err)
		return
	}

	c.JSON(200, gin.H{
		"message": "Key deleted",
	})
}

func ErrorJSON(c *gin.Context, err error) {
	c.JSON(http.StatusInternalServerError, gin.H{
		"error": err.Error(),
	})
}
