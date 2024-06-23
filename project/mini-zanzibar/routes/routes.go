package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/hashicorp/consul/api"
	"github.com/syndtr/goleveldb/leveldb"
	"mini-zanzibar/controllers"
)

func SetupRoutes(r *gin.Engine, levelDB *leveldb.DB, consulDB *api.Client) {
	levelDBRoutes := r.Group("/api/leveldb/")
	{
		levelDBController := controllers.NewLevelDBController(levelDB)
		levelDBRoutes.GET("all", levelDBController.Get)
		levelDBRoutes.GET(":key", levelDBController.GetByKey)
		levelDBRoutes.POST("", levelDBController.Post)
		levelDBRoutes.DELETE(":key", levelDBController.Delete)
	}

	consulDBRoutes := r.Group("/api/consuldb/")
	{
		consulDBController := controllers.NewConsulDBController(consulDB)
		consulDBRoutes.GET("all", consulDBController.Get)
		consulDBRoutes.GET(":key", consulDBController.GetByKey)
		consulDBRoutes.POST("", consulDBController.Post)
		consulDBRoutes.DELETE(":key", consulDBController.Delete)
	}

}
