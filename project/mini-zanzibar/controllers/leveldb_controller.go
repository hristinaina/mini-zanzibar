package controllers

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/syndtr/goleveldb/leveldb"
	"mini-zanzibar/dtos"
	"net/http"
)

type LevelDBController struct {
	db *leveldb.DB
}

func NewLevelDBController(db *leveldb.DB) LevelDBController {
	return LevelDBController{db}
}

func (lc LevelDBController) Get(c *gin.Context) {
	iterator := lc.db.NewIterator(nil, nil)
	defer iterator.Release()

	data := make(map[string]string)
	for iterator.Next() {
		key := iterator.Key()
		value := iterator.Value()
		data[string(key)] = string(value)
	}
	if err := iterator.Error(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, data)
}

func (lc LevelDBController) Post(c *gin.Context) {
	var kv dtos.KeyValue

	if err := c.ShouldBindJSON(&kv); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	fmt.Println("key: ", kv.Key)
	fmt.Println("value ", kv.Value)

	err := lc.db.Put([]byte(kv.Key), []byte(kv.Value), nil)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, gin.H{
		"message": "Key-Value pair saved",
	})
}

func (lc LevelDBController) GetByKey(c *gin.Context) {
	key := c.Param("key")

	value, err := lc.db.Get([]byte(key), nil)
	if errors.Is(err, leveldb.ErrNotFound) {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Key not found",
		})
		return
	} else if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, gin.H{
		"key":   key,
		"value": string(value),
	})
}

func (lc LevelDBController) Delete(c *gin.Context) {
	key := c.Param("key")

	err := lc.db.Delete([]byte(key), nil)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, gin.H{
		"message": "Key deleted",
	})
}
