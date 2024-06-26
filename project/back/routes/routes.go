package routes

import (
	"back/controllers"
	"back/middleware"
	"back/services"
	"database/sql"
	"log"
	"os"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine, db *sql.DB) {
	logService, err := services.NewLogService(os.Getenv("LOGS_FILE"))
	if err != nil {
		log.Fatalf("Failed to initialize LogService: %v", err)
	}

	userRoutes := r.Group("/api/users")
	{
		authController := controllers.NewUserController(db, logService)
		middleware := middleware.NewMiddleware(db, logService)
		userRoutes.POST("/login", authController.Login)
		userRoutes.POST("/logout", middleware.RequireAuth, authController.Logout)
	}

	aclRoutes := r.Group("/api/acl")
	{
		aclController := controllers.NewACLController(logService)
		middleware := middleware.NewMiddleware(db, logService)
		aclRoutes.POST("", middleware.RequireAuth, aclController.Add)
		aclRoutes.PUT("", middleware.RequireAuth, aclController.Check)
	}

	nameSpaceRoutes := r.Group("/api/ns/")
	{
		middleware := middleware.NewMiddleware(db, logService)
		nsController := controllers.NewNSController(logService)
		nameSpaceRoutes.GET("all", middleware.RequireAuth, nsController.Get)
		nameSpaceRoutes.GET(":key", middleware.RequireAuth, nsController.GetByNamespace)
		nameSpaceRoutes.POST("", middleware.RequireAuth, nsController.AddNamespace)
		nameSpaceRoutes.DELETE(":key", middleware.RequireAuth, nsController.Delete)
	}

	dataRoutes := r.Group("/api/data/")
	{
		dataController := controllers.NewDataController(logService)
		middleware := middleware.NewMiddleware(db, logService)
		dataRoutes.GET("all", middleware.RequireAuth, dataController.GetAll)
		dataRoutes.GET(":key", middleware.RequireAuth, dataController.GetByKey)
		dataRoutes.POST("", middleware.RequireAuth, dataController.Add)
		dataRoutes.DELETE(":key", middleware.RequireAuth, dataController.Delete)
	}

	fileRoutes := r.Group("/api/files")
	{
		fileController := controllers.NewFileController(db, logService)
		middleware := middleware.NewMiddleware(db, logService)
		fileRoutes.POST("/create", middleware.RequireAuth, fileController.Create)
		fileRoutes.PUT("/modify", middleware.RequireAuth, fileController.Modify)
		fileRoutes.POST("/share", middleware.RequireAuth, fileController.ShareAccess)
		fileRoutes.GET("/user", middleware.RequireAuth, fileController.GetUserFiles)
		fileRoutes.GET("/shared", middleware.RequireAuth, fileController.GetSharedFiles)
	}
}
