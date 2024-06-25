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

	aclRoutes := r.Group("/api/data")
	{
		aclController := controllers.NewACLController()
		middleware := middleware.NewMiddleware(db)
		aclRoutes.POST("", middleware.RequireAuth, aclController.Add)
		aclRoutes.PUT("", middleware.RequireAuth, aclController.Check)
	}
}
