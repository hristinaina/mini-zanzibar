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
	service services.ILevelDBService
}

func NewLevelDBController(db *leveldb.DB) LevelDBController {
	return LevelDBController{service: services.NewLevelDBService(db)}
}

func (lc LevelDBController) Get(c *gin.Context) {
	data, err := lc.service.GetAll()
	if err != nil {
		errs.InternalServerError(c, err)
		return
	}
	c.JSON(200, data)
}

func (lc LevelDBController) Post(c *gin.Context) {
	var kv dtos.KeyValue

	if err := c.ShouldBindJSON(&kv); err != nil {
		errs.BadRequestError(c, err)
		return
	}

	err := lc.service.Add(kv.Key, kv.Value)
	if err != nil {
		errs.InternalServerError(c, err)
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
		errs.NotFoundError(c)
		return
	} else if err != nil {
		errs.InternalServerError(c, err)
		return
	}

	c.JSON(200, gin.H{
		"key":   key,
		"value": value,
	})
}

func (lc LevelDBController) Delete(c *gin.Context) {
	key := c.Param("key")

	err := lc.service.Delete(key)
	if err != nil {
		errs.NotFoundError(c)
		return
	}

	c.JSON(200, gin.H{
		"message": "Key deleted",
	})
}
