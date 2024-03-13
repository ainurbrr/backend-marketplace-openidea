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

	ver := e.Group("/v1")

	user := ver.Group("/user")
	user.POST("/register", controller.RegisterUserController)
	user.POST("/login", controller.LoginUserController)
}
