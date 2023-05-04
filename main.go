package main

import (
	"fmt"

	"github.com/BEERBAL-HASH/Rule-Engine/config"

	"github.com/gofiber/fiber/v2"
)

func main() {
	//welcome message
	fmt.Println("Welcome to Rule Engine!!!!")

	//creating a new instance of fiber
	app := fiber.New()

	//connect the database
	config.ConnectDB()

	//server Listening on PORT 3000
	app.Listen(":3000")
}
