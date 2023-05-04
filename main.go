package main

import (
	"fmt"

	"github.com/BEERBAL-HASH/Rule-Engine/config"
	"github.com/BEERBAL-HASH/Rule-Engine/controllers"
	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	app.Get("/getTriggers", controllers.GetTriggers)
	app.Get("/getActions", controllers.GetActions)
	app.Post("/user/new", controllers.NewUser)

	//grouping for action handlers
	action := app.Group("/Action")
	action.Put("/arithmetic", controllers.ArithmeticAction)
	action.Put("/update", controllers.UpdateAction)
	action.Get("/plot", controllers.PlotAction)
}

func main() {
	//welcome message
	fmt.Println("Welcome to Rule Engine!!!!")

	//creating a new instance of fiber
	app := fiber.New()

	//connect the database
	config.ConnectDB()

	SetupRoutes(app)
	//server Listening on PORT 3000
	app.Listen(":3000")

}
