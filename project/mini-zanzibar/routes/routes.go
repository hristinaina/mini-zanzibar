package routes

import (
	"log"
	"mini-zanzibar/controllers"
	"mini-zanzibar/middleware"
	"mini-zanzibar/services"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/hashicorp/consul/api"
	"github.com/syndtr/goleveldb/leveldb"
)

func SetupRoutes(r *gin.Engine, levelDB *leveldb.DB, consulDB *api.Client) {
	logService, err := services.NewLogService(os.Getenv("LOGS_FILE"))
	if err != nil {
		log.Fatalf("Failed to initialize LogService: %v", err)
	}

	levelDBRoutes := r.Group("/api/leveldb/")
	{
		levelDBController := controllers.NewLevelDBController(levelDB, logService)
		middleware, _ := middleware.NewMiddleware(logService)
		levelDBRoutes.GET("all", middleware.ApiKeyAuthMiddleware(), levelDBController.Get)
		levelDBRoutes.GET(":key", middleware.ApiKeyAuthMiddleware(), levelDBController.GetByKey)
		levelDBRoutes.POST("", middleware.ApiKeyAuthMiddleware(), levelDBController.Post)
		levelDBRoutes.DELETE(":key", middleware.ApiKeyAuthMiddleware(), levelDBController.Delete)
	}

	consulDBRoutes := r.Group("/api/consuldb/")
	{
		middleware, _ := middleware.NewMiddleware(logService)
		consulDBController := controllers.NewConsulDBController(consulDB, logService)
		consulDBRoutes.GET("all", middleware.ApiKeyAuthMiddleware(), consulDBController.Get)
		consulDBRoutes.GET(":key", middleware.ApiKeyAuthMiddleware(), consulDBController.GetByNamespace)
		consulDBRoutes.POST("", middleware.ApiKeyAuthMiddleware(), consulDBController.AddNamespace)
		consulDBRoutes.DELETE(":key", middleware.ApiKeyAuthMiddleware(), consulDBController.Delete)
	}

	aclRoutes := r.Group("/api/acl/")
	{
		middleware, _ := middleware.NewMiddleware(logService)
		aclController := controllers.NewACLController(levelDB, consulDB, logService)
		aclRoutes.POST("", middleware.ApiKeyAuthMiddleware(), aclController.Add)
		aclRoutes.PUT("", middleware.ApiKeyAuthMiddleware(), aclController.Check)
	}

}
