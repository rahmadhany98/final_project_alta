package route

import (
	"expense-manager/constant"
	"expense-manager/controller"

	"github.com/labstack/echo"
	echoMiddleware "github.com/labstack/echo/middleware"
)

func New() *echo.Echo {
	e := echo.New()

	e.POST("/login", controller.LoginUserController)
	jwt := e.Group("/")
	jwt.Use(echoMiddleware.JWT([]byte(constant.SECRET_JWT)))
	jwt.GET("", controller.HelloController)
	return e
}
