package controller

import (
	"backend-marketplace-openidea/models/payload"
	"backend-marketplace-openidea/service"
	"net/http"

	"github.com/labstack/echo/v4"
)

func RegisterUserController(c echo.Context) error {
	payloadUser := payload.CreateUserRequest{}
	c.Bind(&payloadUser)

	if err := c.Validate(&payloadUser); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "error payload create user",
			"error":   err.Error(),
		})
	}

	response, err := service.RegisterUser(&payloadUser)
	if err != nil {
		return c.JSON(http.StatusConflict, map[string]interface{}{
			"messages": "error create user",
			"error":    err.Error(),
		})
	}

	return c.JSON(http.StatusCreated, payload.Response{
		Message: "User registered successfully",
		Data:    response,
	})
}

func LoginUserController(c echo.Context) error {
	payloadUser := payload.LoginUserRequest{}

	c.Bind(&payloadUser)

	if err := c.Validate(&payloadUser); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "error payload login user",
			"error":   err.Error(),
		})
	}

	response, err := service.LoginUser(&payloadUser)
	if err != nil {
		if err.Error() == "Email Not Registered" {
			return c.JSON(http.StatusNotFound, err.Error())
		}
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	return c.JSON(200, payload.Response{
		Message: "User logged successfully",
		Data:    response,
	})
}
