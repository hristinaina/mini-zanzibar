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
}
