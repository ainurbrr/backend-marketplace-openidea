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
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"messages": "error create user",
			"error":    err.Error(),
		})
	}

	return c.JSON(http.StatusOK, payload.Response{
		Message: "success register user",
		Data:    response,
	})
}
