package database

import (
	"backend-marketplace-openidea/config"
	"backend-marketplace-openidea/models"
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

func IsUsernameAvailable(username string) bool {
	var count int64
	if err := config.DB.QueryRow(`SELECT COUNT(*) FROM users WHERE username = $1`, username).Scan(&count); err != nil {
		echo.NewHTTPError(http.StatusNotFound, err)
		return false
	}

	return count == 0
}

func CreateUser(user *models.User) (error, string) {
	var id string
	if err := config.DB.QueryRow(`INSERT INTO users (username, name, password) VALUES ($1, $2, $3) RETURNING id`, user.Username, user.Name, user.Password).Scan(&id); err != nil {
		return err, id
	}
	fmt.Printf("Inserted a single record %v", id)
	return nil, id
}

func GetUserByUsername(username string) (*models.User, error) {
	var user models.User
	if err := config.DB.QueryRow(`SELECT id, username, name, password FROM users WHERE username = $1`, username).Scan(&user.ID, &user.Username, &user.Name, &user.Password); err != nil {
		return &user, err
	}

	return &user, nil
}
