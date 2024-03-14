package routes

import (
	"backend-marketplace-openidea/controller"
	"backend-marketplace-openidea/middleware"
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

	product := ver.Group("/product", middleware.IsLoggedIn)
	product.POST("", controller.CreateProductController)


	bankA := ver.Group("/bank/account", middleware.IsLoggedIn)
	bankA.POST("", controller.CreateBankAccountController)
	bankA.GET("", controller.GetBankAccountByUserIdController)
	bankA.PATCH("/:bankAccountId", controller.UpdateBankAccountController)
	bankA.DELETE("/:bankAccountId", controller.DeleteBankAccountByUserIdController)

}
