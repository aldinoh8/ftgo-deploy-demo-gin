package main

import (
	"aldinoh8/example-gin/config"
	"aldinoh8/example-gin/controller"
	"aldinoh8/example-gin/repository"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	_ "github.com/joho/godotenv/autoload"
)

func main() {
	db := config.InitDb(os.Getenv("DB"))

	mainRepository := repository.New(db)
	mainController := controller.New(mainRepository)

	app := gin.New()

	todos := app.Group("/todos")
	todos.POST("", mainController.Create)
	todos.GET("", mainController.FindAll)

	if err := app.Run(":8080"); err != nil {
		log.Fatal(err.Error())
	}
}
