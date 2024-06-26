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

	ver.POST("/image", controller.UploadImage)

	user := ver.Group("/user")
	user.POST("/register", controller.RegisterUserController)
	user.POST("/login", controller.LoginUserController)

	product := ver.Group("/product", middleware.IsLoggedIn)
	product.POST("", controller.CreateProductController)
	product.PATCH("/:productId", controller.UpdateProductController)
	product.GET("/:productId", controller.GetProductById)
	product.GET("", controller.GetAllProductsController)
	product.DELETE("/:productId", controller.DeleteProductController)
	product.POST("/:productId/stock", controller.UpdateProductStockController)
	product.POST("/:productId/buy", controller.CreatePaymentController)

	bankA := ver.Group("/bank/account", middleware.IsLoggedIn)
	bankA.POST("", controller.CreateBankAccountController)
	bankA.GET("", controller.GetBankAccountByUserIdController)
	bankA.PATCH("/:bankAccountId", controller.UpdateBankAccountController)
	bankA.DELETE("/:bankAccountId", controller.DeleteBankAccountByUserIdController)

}
