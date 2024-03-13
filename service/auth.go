package service

import (
	"backend-marketplace-openidea/middleware"
	"backend-marketplace-openidea/models"
	"backend-marketplace-openidea/models/payload"
	"backend-marketplace-openidea/repository/database"
	"errors"

	"golang.org/x/crypto/bcrypt"
)

func RegisterUser(req *payload.CreateUserRequest) (resp payload.CreateOrLoginUserResponse, err error) {
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

	resp = payload.CreateOrLoginUserResponse{
		Username:   newUser.Username,
		Name:       newUser.Name,
		AccessToken: token,
	}

	return
}

func LoginUser(req *payload.LoginUserRequest) (res payload.CreateOrLoginUserResponse, err error) {

	user, err := database.GetUserByUsername(req.Username)
	if err != nil {
		return res, errors.New("Email Not Registered")
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password))
	if err != nil && err == bcrypt.ErrMismatchedHashAndPassword {
		return res, errors.New("Wrong Password")
	}

	token, err := middleware.CreateToken(user.ID)
	if err != nil {
		return res, errors.New("Failed To Create Token")
	}

	res = payload.CreateOrLoginUserResponse{
		Username: user.Username,
		Name: user.Name,
		AccessToken: token,
	}

	return
}
