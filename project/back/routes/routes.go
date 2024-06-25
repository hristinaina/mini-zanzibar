package routes

import (
	"back/controllers"
	"back/middleware"
	"database/sql"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine, db *sql.DB) {
	userRoutes := r.Group("/api/users")
	{
		authController := controllers.NewUserController(db)
		middleware := middleware.NewMiddleware(db)
		userRoutes.POST("/login", authController.Login)
		userRoutes.POST("/logout", middleware.RequireAuth, authController.Logout)
	}

	aclRoutes := r.Group("/api/acl")
	{
		aclController := controllers.NewACLController()
		middleware := middleware.NewMiddleware(db)
		aclRoutes.POST("", middleware.RequireAuth, aclController.Add)
		aclRoutes.PUT("", middleware.RequireAuth, aclController.Check)
	}

	nameSpaceRoutes := r.Group("/api/ns/")
	{
		middleware := middleware.NewMiddleware(db)
		nsController := controllers.NewNSController()
		nameSpaceRoutes.GET("all", middleware.RequireAuth, nsController.Get)
		nameSpaceRoutes.GET(":key", middleware.RequireAuth, nsController.GetByNamespace)
		nameSpaceRoutes.POST("", middleware.RequireAuth, nsController.AddNamespace)
		nameSpaceRoutes.DELETE(":key", middleware.RequireAuth, nsController.Delete)
	}

	dataRoutes := r.Group("/api/data/")
	{
		dataController := controllers.NewDataController()
		dataRoutes.GET("all", dataController.GetAll)
		dataRoutes.GET(":key", dataController.GetByKey)
		dataRoutes.POST("", dataController.Add)
		dataRoutes.DELETE(":key", dataController.Delete)
	}
}
