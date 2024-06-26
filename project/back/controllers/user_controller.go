package controllers

import (
	"back/models"
	"back/services"
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	service    services.UserService
	logService *services.LogService
}

func NewUserController(db *sql.DB, logService *services.LogService) UserController {
	return UserController{service: services.NewUserService(db), logService: logService}
}

func (uc UserController) Login(c *gin.Context) {
	var input services.LoginInput
	if err := c.ShouldBindJSON(&input); err != nil {
		uc.logService.Error("Failed to read body in Login")
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to read body"})
		return
	}

	uc.logService.Info("Processing Login request for user: " + input.Email)
	token, err := uc.service.Login(input)
	if err != nil {
		uc.logService.Error("Failed to login user: " + input.Email)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.SetSameSite(http.SameSiteLaxMode)
	c.SetCookie("Authorization", token, 3600*24*30, "", "", true, true)

	uc.logService.Info("User logged in successfully: " + input.Email)
	c.JSON(http.StatusOK, gin.H{"token": token})
}

func (uc UserController) Logout(c *gin.Context) {
	currentUserFromCookie, _ := c.Get("user")
	currentUser := currentUserFromCookie.(*models.User)

	uc.logService.Info("Processing Logout request for user: " + currentUser.Email)
	c.SetCookie("Authorization", "", -1, "", "", false, true)
	uc.logService.Info("User logged out successfully: " + currentUser.Email)
	c.JSON(http.StatusOK, gin.H{"message": "Successful logout!"})
}
