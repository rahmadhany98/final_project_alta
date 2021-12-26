package controller

import (
	"net/http"

	"expense-manager/config"
	"expense-manager/middleware"
	"expense-manager/model"

	"github.com/labstack/echo"
)

func LoginUserController(e echo.Context) error {
	user := model.User{}
	e.Bind(&user)

	err := config.DB.Where("email = ? AND password = ?", user.Email, user.Password).First(&user).Error
	if err != nil {
		return e.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "login failed",
			"error":   err.Error(),
		})
	}

	token, err := middleware.CreateToken(user.ID, user.Name)
	if err != nil {
		return e.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "login failed",
			"error":   err.Error(),
		})
	}

	return e.JSON(http.StatusOK, map[string]interface{}{
		"message": "success",
		"token":   token,
	})
}

func CreateUserController(e echo.Context) error {

	user := model.User{}
	e.Bind(&user)

	err := config.DB.Save(&user).Error
	if err != nil {
		return e.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": err.Error(),
		})
	}

	return e.JSON(http.StatusOK, map[string]interface{}{
		"message": "success",
		"user":    user,
	})
}
