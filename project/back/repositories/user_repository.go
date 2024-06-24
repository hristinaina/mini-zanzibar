package repositories

import (
	"back/models"
	"database/sql"
	"errors"
	"fmt"
)

type UserRepository interface {
	GetAll() []models.User
	GetUserByEmail(email string) (*models.User, error)
	GetUserById(id int) (*models.User, error)
}

type UserRepositoryImpl struct {
	db *sql.DB
}

var ErrUserNotFound = errors.New("user not found")

func NewUserRepository(db *sql.DB) UserRepository {
	return &UserRepositoryImpl{db: db}
}

func (res *UserRepositoryImpl) GetAll() []models.User {
	query := "SELECT * FROM users"
	rows, err := res.db.Query(query)
	if CheckIfError(err) {
		return nil
	}
	defer rows.Close()

	var users []models.User
	for rows.Next() {
		var (
			user models.User
		)

		if err := rows.Scan(&user.Id,
			&user.Name, &user.Surname, &user.Email, &user.Password); err != nil {
			fmt.Println("Error: ", err.Error())
			return []models.User{}
		}
		users = append(users, user)
	}

	return users
}

func (res *UserRepositoryImpl) GetUserByEmail(email string) (*models.User, error) {
	var user models.User

	query := "SELECT * FROM users WHERE email = $1"
	row := res.db.QueryRow(query, email)

	err := row.Scan(&user.Id, &user.Name, &user.Surname, &user.Email, &user.Password)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, ErrUserNotFound
		}
		return nil, err
	}
	return &user, nil
}

func (res *UserRepositoryImpl) GetUserById(id int) (*models.User, error) {
	var user models.User

	query := "SELECT * FROM users WHERE id = $1"
	row := res.db.QueryRow(query, id)

	err := row.Scan(&user.Id, &user.Name, &user.Surname, &user.Email, &user.Password)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, ErrUserNotFound
		}
		return nil, err
	}

	return &user, nil
}

func CheckIfError(err error) bool {
	if err != nil {
		fmt.Println("Error: ", err.Error())
		return true
	}
	return false
}
