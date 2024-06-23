package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/syndtr/goleveldb/leveldb"
	"mini-zanzibar/controllers"
)

func SetupRoutes(r *gin.Engine, levelDB *leveldb.DB) {
	routes := r.Group("/api/")
	{
		levelDBController := controllers.NewLevelDBController(levelDB)
		routes.GET("all", levelDBController.Get)
		routes.GET(":key", levelDBController.GetByKey)
		routes.POST("", levelDBController.Post)
		routes.DELETE(":key", levelDBController.Delete)
	}

}
