package main

import (
	"task4/controller"
	"task4/dto"
)

func main() {
	// gorm
	dto.CreateTable()
	// gin
	controller.Router.Run()
}
