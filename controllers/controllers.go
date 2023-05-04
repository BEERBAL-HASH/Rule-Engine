package controllers

import (
	"github.com/BEERBAL-HASH/Rule-Engine/config"
	"github.com/BEERBAL-HASH/Rule-Engine/models"
	"github.com/gofiber/fiber/v2"
)

// fetching predefined triggers
func GetTriggers(c *fiber.Ctx) error {
	data := models.Triggers{
		Trigger1: "Age is less than 60 Yrs",
		Trigger2: "Age is above than 60 Yrs",
	}
	return c.Status(200).JSON(data)
}

// fetching predefined Actions
func GetActions(c *fiber.Ctx) error {
	data := models.Actions{
		Action1: "Copy one value to Another",
		Action2: "update the Salary after given increment",
		Action3: "Plot a graph for salary vs increment",
	}
	return c.Status(200).JSON(data)
}

// creating new user and storing its attributes in the database
func NewUser(c *fiber.Ctx) error {
	//gorm Database instance
	db := config.DBconn
	user := new(models.Employee)

	// Store the body in the user variable and return error if encountered
	err := c.BodyParser(user)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Review your input", "error": err.Error()})
	}

	// Create the user and return error if encountered
	err = db.Create(&user).Error

	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Could not create user", "error": err.Error()})
	}

	// Return the created user
	return c.Status(201).JSON(fiber.Map{"status": "success", "message": "Created a new user", "data": user})

}
