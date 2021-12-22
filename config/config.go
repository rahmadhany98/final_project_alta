package config

import (
	"expense-manager/model"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

//initdb
func InitDB() {
	dsn := "root:Rahmadh4ny98@tcp(localhost:3306)/expense?charset=utf8mb4&parseTime=True&loc=Local"
	var err error
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err.Error())
	}
	initMigrate()
}

func initMigrate() {
	DB.AutoMigrate(&model.User{}, &model.Transaction{})
}
