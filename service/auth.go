package service

import (
	"backend-marketplace-openidea/middleware"
	"backend-marketplace-openidea/models"
	"backend-marketplace-openidea/models/payload"
	"backend-marketplace-openidea/repository/database"
	"errors"

	"golang.org/x/crypto/bcrypt"
)

func RegisterUser(req *payload.CreateUserRequest) (resp payload.CreateUserResponse, err error) {
	passwordHash, err := bcrypt.GenerateFromPassword([]byte(req.Password), 8)
	if err != nil {
		return
	}

	if !database.IsUsernameAvailable(req.Username) {
		return resp, errors.New("Username is already registered")
	}

	newUser := &models.User{
		Username: req.Username,
		Name:     req.Name,
		Password: string(passwordHash),
	}

	err, id := database.CreateUser(newUser)
	if err != nil {
		return
	}

	token, err := middleware.CreateToken(id)
	if err != nil {
		return resp, errors.New("Failed To Create Token")
	}

	resp = payload.CreateUserResponse{
		Username:   newUser.Username,
		Name:       newUser.Name,
		AccesToken: token,
	}

	return
}
