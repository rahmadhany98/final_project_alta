package controller

import (
	"expense-manager/config"
	"expense-manager/model"
	"net/http"

	"github.com/labstack/echo"
)

func GetExpenseController(e echo.Context) error {
	iduser := ReturnID(e)
	category := "expense"
	var transactions []model.Transaction
	startdate := e.QueryParam("start_date")
	enddate := e.QueryParam("end_date")
	var expense int
	if startdate != "" && enddate != "" {
		config.DB.Debug().Table("transactions").Select("sum(amount)").Where("category = ? AND user_id = ? AND date between ? AND ?", category, iduser, startdate, enddate).Row().Scan(&expense)
		err := config.DB.Debug().Where("category = ? AND user_id = ? AND date between ? AND ?", category, iduser, startdate, enddate).Find(&transactions).Error
		if err != nil {
			return e.JSON(http.StatusInternalServerError, map[string]interface{}{
				"message": err.Error(),
			})
		}
		if len(transactions) == 0 {
			return e.JSON(http.StatusInternalServerError, map[string]interface{}{
				"message": "Data Tidak Ditemukan!",
			})
		}
		return e.JSON(http.StatusOK, map[string]interface{}{
			"message": "success",
			"data":    transactions,
			"total":   expense,
		})
	} else {
		config.DB.Debug().Table("transactions").Select("sum(amount)").Where("category = ? AND user_id = ? ", category, iduser).Row().Scan(&expense)
		err := config.DB.Debug().Where("category = ? AND user_id = ?", category, iduser).Find(&transactions).Error
		if err != nil {
			return e.JSON(http.StatusInternalServerError, map[string]interface{}{
				"message": err.Error(),
			})
		}
		if len(transactions) == 0 {
			return e.JSON(http.StatusInternalServerError, map[string]interface{}{
				"message": "Data Tidak Ditemukan!",
			})
		}

		return e.JSON(http.StatusOK, map[string]interface{}{
			"message": "success",
			"data":    transactions,
			"total":   expense,
		})
	}
}

func GetIncomeController(e echo.Context) error {
	iduser := ReturnID(e)
	category := "income"
	var transactions []model.Transaction
	startdate := e.QueryParam("start_date")
	enddate := e.QueryParam("end_date")
	var income int
	if startdate != "" && enddate != "" {
		config.DB.Debug().Table("transactions").Select("sum(amount)").Where("category = ? AND user_id = ? AND date between ? AND ?", category, iduser, startdate, enddate).Row().Scan(&income)
		err := config.DB.Debug().Where("category = ? AND user_id = ? AND date between ? AND ?", category, iduser, startdate, enddate).Find(&transactions).Error
		if err != nil {
			return e.JSON(http.StatusInternalServerError, map[string]interface{}{
				"message": err.Error(),
			})
		}
		if len(transactions) == 0 {
			return e.JSON(http.StatusInternalServerError, map[string]interface{}{
				"message": "Data Tidak Ditemukan!",
			})
		}

		return e.JSON(http.StatusOK, map[string]interface{}{
			"message": "success",
			"data":    transactions,
			"total":   income,
		})
	} else {
		config.DB.Debug().Table("transactions").Select("sum(amount)").Where("category = ? AND user_id = ?", category, iduser).Row().Scan(&income)
		err := config.DB.Debug().Where("category = ? AND user_id = ?", category, iduser).Find(&transactions).Error
		if err != nil {
			return e.JSON(http.StatusInternalServerError, map[string]interface{}{
				"message": err.Error(),
			})
		}

	}
	if len(transactions) == 0 {
		return e.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "Data Tidak Ditemukan!",
		})
	}
	return e.JSON(http.StatusOK, map[string]interface{}{
		"message": "success",
		"data":    transactions,
		"total":   income,
	})
}

func GetReportController(e echo.Context) error {
	iduser := ReturnID(e)
	var transactions []model.Transaction
	startdate := e.QueryParam("start_date")
	enddate := e.QueryParam("end_date")
	category1 := "income"
	category2 := "expense"
	var expense, income int
	if startdate != "" && enddate != "" {
		config.DB.Debug().Table("transactions").Select("sum(amount)").Where("category = ? AND user_id = ? AND date between ? AND ?", category1, iduser, startdate, enddate).Row().Scan(&income)
		config.DB.Debug().Table("transactions").Select("sum(amount)").Where("category = ? AND user_id = ? AND date between ? AND ?", category2, iduser, startdate, enddate).Row().Scan(&expense)
		profit := income - expense
		err := config.DB.Debug().Where("user_id = ? AND date between ? AND ?", iduser, startdate, enddate).Find(&transactions).Error
		if err != nil {
			return e.JSON(http.StatusInternalServerError, map[string]interface{}{
				"message": err.Error(),
			})
		}
		if len(transactions) == 0 {
			return e.JSON(http.StatusInternalServerError, map[string]interface{}{
				"message": "Data Tidak Ditemukan!",
			})
		}

		return e.JSON(http.StatusOK, map[string]interface{}{
			"message":       "success",
			"data":          transactions,
			"total-income":  income,
			"total-expense": expense,
			"profit":        profit,
		})
	} else {
		config.DB.Debug().Table("transactions").Select("sum(amount)").Where("category = ? AND user_id = ? ", category1, iduser).Row().Scan(&income)
		config.DB.Debug().Table("transactions").Select("sum(amount)").Where("category = ? AND user_id = ?", category2, iduser).Row().Scan(&expense)
		profit := income - expense
		err := config.DB.Debug().Where("user_id = ?", iduser).Find(&transactions).Error
		if err != nil {
			return e.JSON(http.StatusInternalServerError, map[string]interface{}{
				"message": err.Error(),
			})
		}
		if len(transactions) == 0 {
			return e.JSON(http.StatusInternalServerError, map[string]interface{}{
				"message": "Data Tidak Ditemukan!",
			})
		}
		return e.JSON(http.StatusOK, map[string]interface{}{
			"message":       "success",
			"data":          transactions,
			"total-income":  income,
			"total-expense": expense,
			"profit":        profit,
		})
	}
}
