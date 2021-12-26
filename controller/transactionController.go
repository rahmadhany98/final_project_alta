package controller

import (
	"expense-manager/config"
	"expense-manager/model"
	"net/http"
	"strconv"
	"time"

	"github.com/labstack/echo"
)

func HelloController(e echo.Context) error {
	id := ReturnID(e)
	date := time.Now().Format("2006-01-02")
	income := "income"
	expense := "expense"
	var incomeToday, expenseToday int
	config.DB.Debug().Table("transactions").Select("sum(amount)").Where("category = ? AND date = ? AND user_id = ?", income, date, id).Row().Scan(&incomeToday)
	config.DB.Debug().Table("transactions").Select("sum(amount)").Where("category = ? AND date = ? AND user_id = ?", expense, date, id).Row().Scan(&expenseToday)
	return e.JSON(http.StatusOK, map[string]interface{}{
		"massage": "success",
		"income":  incomeToday,
		"expense": expenseToday,
	})
}

func AddNewController(e echo.Context) error {
	id := ReturnID(e)
	transactions := model.Transaction{}
	e.Bind(&transactions)
	transactions.UserID = id
	err := config.DB.Debug().Save(&transactions).Error
	if err != nil {
		return e.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": err.Error(),
		})
	}

	return e.JSON(http.StatusOK, map[string]interface{}{
		"message": "success",
		"data":    transactions,
	})
}

func GetTransactionController(e echo.Context) error {
	iduser := ReturnID(e)
	id, _ := strconv.Atoi(e.Param("id"))
	var transactions []model.Transaction
	err := config.DB.Debug().Where("id = ? AND user_id = ?", id, iduser).Find(&transactions).Error
	if err != nil {
		return e.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": err.Error(),
		})
	}
	if len(transactions) == 0 {
		return e.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "Data Tidak Ditemukan atau Anda TIdak Mempunyai Akses ke Data Tersebut!",
		})
	}

	return e.JSON(http.StatusOK, map[string]interface{}{
		"message": "success",
		"data":    transactions,
	})
}

func UpdateTransactionController(e echo.Context) error {
	iduser := ReturnID(e)
	id, _ := strconv.Atoi(e.Param("id"))
	transactions := model.Transaction{}
	err := config.DB.Debug().Where("id = ? AND user_id = ?", id, iduser).First(&transactions).Error
	if err != nil {
		return e.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": err.Error(),
		})
	}
	e.Bind(&transactions)

	err1 := config.DB.Debug().Save(&transactions).Error
	if err1 != nil {
		return e.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": err1.Error(),
		})
	}

	return e.JSON(http.StatusOK, map[string]interface{}{
		"message": "success",
		"data":    transactions,
	})
}

func DeleteTransactionController(e echo.Context) error {
	iduser := ReturnID(e)
	id, _ := strconv.Atoi(e.Param("id"))
	var transactions []model.Transaction
	err := config.DB.Debug().Model(&transactions).Where("id = ? AND user_id = ?", id, iduser).Delete(&transactions)
	if err.Error != nil {
		return e.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": err.Error,
		})
	}

	return e.JSON(http.StatusOK, map[string]interface{}{
		"message":      "deleted",
		"affectedrows": err.RowsAffected,
	})
}
