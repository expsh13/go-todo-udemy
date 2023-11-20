package main

import (
	"fmt"
	"sqlite/app/controllers"
	"sqlite/app/models"
)

func main() {
	fmt.Println(models.Db)

	controllers.StartMainServer()
}