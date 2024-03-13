package routes

import (
	"backend-marketplace-openidea/controller"
	"backend-marketplace-openidea/util"
	"database/sql"

	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
)

func Routes(e *echo.Echo, db *sql.DB) {
	e.Validator = &util.CustomValidator{Validator: validator.New()}

	e.POST("/register", controller.RegisterUserController)
}
