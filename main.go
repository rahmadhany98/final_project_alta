package main

import (
	"expense-manager/config"
	"expense-manager/route"
)

//controller

func main() {

	config.InitDB()

	e := route.New()
	e.Start(":8080")
}
