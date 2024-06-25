package routes

import (
	"mini-zanzibar/controllers"
	"mini-zanzibar/middleware"

	"github.com/gin-gonic/gin"
	"github.com/hashicorp/consul/api"
	"github.com/syndtr/goleveldb/leveldb"
)

func SetupRoutes(r *gin.Engine, levelDB *leveldb.DB, consulDB *api.Client) {
	levelDBRoutes := r.Group("/api/leveldb/")
	{
		levelDBController := controllers.NewLevelDBController(levelDB)
		middleware := middleware.NewMiddleware()
		levelDBRoutes.GET("all", middleware.ApiKeyAuthMiddleware(), levelDBController.Get)
		levelDBRoutes.GET(":key", middleware.ApiKeyAuthMiddleware(), levelDBController.GetByKey)
		levelDBRoutes.POST("", middleware.ApiKeyAuthMiddleware(), levelDBController.Post)
		levelDBRoutes.DELETE(":key", middleware.ApiKeyAuthMiddleware(), levelDBController.Delete)
	}

	consulDBRoutes := r.Group("/api/consuldb/")
	{
		middleware := middleware.NewMiddleware()
		consulDBController := controllers.NewConsulDBController(consulDB)
		consulDBRoutes.GET("all", middleware.ApiKeyAuthMiddleware(), consulDBController.Get)
		consulDBRoutes.GET(":key", middleware.ApiKeyAuthMiddleware(), consulDBController.GetByNamespace)
		consulDBRoutes.POST("", middleware.ApiKeyAuthMiddleware(), consulDBController.AddNamespace)
		consulDBRoutes.DELETE(":key", middleware.ApiKeyAuthMiddleware(), consulDBController.Delete)
	}

	aclRoutes := r.Group("/api/acl/")
	{
		middleware := middleware.NewMiddleware()
		aclController := controllers.NewACLController(levelDB, consulDB)
		aclRoutes.POST("", middleware.ApiKeyAuthMiddleware(), aclController.Add)
		aclRoutes.PUT("", middleware.ApiKeyAuthMiddleware(), aclController.Check)
	}

}
