package services

import (
	"back/repositories"
	"database/sql"
	"errors"
	"os"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
)

type UserService struct {
	repo repositories.UserRepository
}

func NewUserService(db *sql.DB) UserService {
	return UserService{repo: repositories.NewUserRepository(db)}
}

// request body
type LoginInput struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func (us UserService) Login(input LoginInput) (string, error) {
	// looked up requested user by email
	user, err := us.repo.GetUserByEmail(input.Email)
	if err != nil {
		return "", errors.New("invalid username or password")
	}

	// compare sent password with saved user hash password
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(input.Password))
	if err != nil {
		return "", errors.New("invalid username or password")
	}

	// generate a jwt token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub":  user.Id,
		"role": strconv.Itoa(1),
		"exp":  time.Now().Add(time.Hour * 24).Unix(), // 1 day
	})

	// Sign and get the complete encoded token as a string using the secret
	tokenString, err := token.SignedString([]byte(os.Getenv("API_SECRET")))
	if err != nil {
		return "", errors.New("failed to create token")
	}

	return tokenString, nil
}
