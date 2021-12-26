package route

import (
	"expense-manager/constant"
	"expense-manager/controller"
	"expense-manager/middleware"
	"os"

	"github.com/labstack/echo"
	echoMiddleware "github.com/labstack/echo/middleware"
	"github.com/sirupsen/logrus"
)

func New() *echo.Echo {
	e := echo.New()
	logrus.SetOutput(os.Stdout)
	e.Use(middleware.LogMiddleware)
	e.POST("/login", controller.LoginUserController)
	e.POST("users", controller.CreateUserController)
	jwt := e.Group("/")
	jwt.Use(echoMiddleware.JWT([]byte(constant.SECRET_JWT)))
	jwt.POST("users", controller.CreateUserController)
	jwt.GET("", controller.HelloController)
	jwt.POST("", controller.AddNewController)
	jwt.GET("transaction/:id", controller.GetTransactionController)
	jwt.PUT("transaction/:id", controller.UpdateTransactionController)
	jwt.DELETE("transaction/:id", controller.DeleteTransactionController)
	jwt.GET("expense", controller.GetExpenseController)
	jwt.GET("income", controller.GetIncomeController)
	jwt.GET("report", controller.GetReportController)
	return e
}
