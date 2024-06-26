package middleware

import (
	"back/repositories"
	"back/services"
	"back/utils"
	"database/sql"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

type Middleware struct {
	repo       repositories.UserRepository
	logService *services.LogService
}

func NewMiddleware(db *sql.DB, logService *services.LogService) Middleware {
	mw := Middleware{
		repo:       repositories.NewUserRepository(db),
		logService: logService,
	}
	mw.logService.Info("Middleware initialized successfully")
	return mw
}

func (mw Middleware) RequireAuth(c *gin.Context) {
	tokenString, err := c.Cookie("Authorization")
	if err != nil {
		mw.logService.Error("Failed to retrieve Authorization cookie: " + err.Error())
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(os.Getenv("API_SECRET")), nil
	})

	if err != nil || token == nil {
		mw.logService.Error("Token validation failed: " + err.Error())
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		expFloat, ok := claims["exp"].(float64)
		if !ok {
			mw.logService.Error("Failed to parse expiration time from token claims")
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}
		exp := time.Unix(int64(expFloat), 0)
		if time.Now().After(exp) {
			mw.logService.Info("Token expired")
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		subFloat, ok := claims["sub"].(float64)
		if !ok {
			mw.logService.Error("Failed to parse subject ID from token claims")
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}
		sub := int(subFloat)
		user, err := mw.repo.GetUserById(sub)
		if err != nil {
			mw.logService.Error("Failed to retrieve user from repository: " + err.Error())
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		mw.logService.Info("User authenticated successfully")
		c.Set("user", user)
		c.Next()
	} else {
		mw.logService.Error("Token claims validation failed")
		c.AbortWithStatus(http.StatusUnauthorized)
	}
}

func UserMiddleware(c *gin.Context) {
	cookie, err := c.Cookie("Authorization")
	if err != nil {
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	claims, err := utils.ParseToken(cookie)
	if err != nil {
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	role, ok := claims["role"].(string)
	if !ok {
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	if role != "1" {
		c.AbortWithStatus(http.StatusForbidden)
		return
	}

	c.Next()
}
