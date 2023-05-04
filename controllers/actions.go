package controllers

import (
	"image/color"

	"strconv"

	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/vg"

	"github.com/BEERBAL-HASH/Rule-Engine/config"
	"github.com/BEERBAL-HASH/Rule-Engine/models"
	"github.com/gofiber/fiber/v2"
)

func ArithmeticAction(c *fiber.Ctx) error {
	//fetching employeeId using requestHeader
	employeeId, _ := strconv.Atoi(c.Get("employeeId"))

	//creating gorm DB instance
	db := config.DBconn

	//user variable of type Employee
	var user models.Employee

	//querying database to check whether given id exists or not
	err := db.Model(&user).Where("id=?", employeeId).First(&user).Error

	//error handling if given id doesn't exist
	if err != nil {
		return c.Status(404).JSON(fiber.Map{
			"status":  "success",
			"error":   err.Error(),
			"message": "no such user with corresponding id",
		})
	}

	//fetching salary for current user
	userSalary := user.Salary

	increment := 25

	// if(age>60){
	// 	increment=10
	// }

	//updating salary using below formula
	// updatedSalary = userSalary + (25% of increment after one financial year)
	UpdatedSalary := userSalary + (userSalary*increment)/100
	user.Salary = UpdatedSalary

	//saving updatedSalary in db
	db.Save(&user)

	//Response
	return c.JSON(fiber.Map{
		"status":        "success",
		"UpdatedSalary": UpdatedSalary,
	})
}

func UpdateAction(c *fiber.Ctx) error {

	//fetching requestHeaders
	employeeId, _ := strconv.Atoi(c.Get("employeeId"))
	param1 := c.Get("param1")
	param2 := c.Get("param2")

	db := config.DBconn

	var user models.Employee

	//copy one value with another
	// Assumption: user will try to change value with similar datatypes, in this case string will be preferable

	err := db.Model(&user).Where("id=?", employeeId).First(&user).Error

	if err != nil {
		return c.Status(404).JSON(fiber.Map{
			"status":  "failed",
			"error":   err.Error(),
			"message": "no such user with corresponding id",
		})
	}

	if param1 == "department_name" {

		if param2 == "employee_name" {
			user.DepartmentName = user.EmployeeName
		}
		if param2 == "profile" {
			user.DepartmentName = user.Profile

		}
	}

	if param1 == "employee_name" {

		if param2 == "deparment_name" {
			user.EmployeeName = user.DepartmentName
		}
		if param2 == "profile" {
			user.EmployeeName = user.Profile
		}
	}
	if param1 == "profile" {

		if param2 == "employee_name" {
			user.Profile = user.EmployeeName
		}
		if param2 == "department_name" {
			user.Profile = user.DepartmentName
		}
	}

	//error handling while updating user in the database
	err = db.Save(&user).Error
	if err != nil {

		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Could not update the user", "error": err.Error()})
	}

	//returning successful response
	return c.Status(200).JSON(fiber.Map{
		"status":  "success",
		"message": "data has been successfully updated",
	})
}

func PlotAction(c *fiber.Ctx) error {
	p := plot.New()

	p.Title.Text = "Salary Graph"
	p.X.Label.Text = "Current Salary"  // current Salary
	p.Y.Label.Text = "Expected Salary" //expected Salary

	// expected salary calculator function
	y := plotter.NewFunction(func(x float64) float64 { return 1.3*x + 100000 })
	y.Color = color.RGBA{B: 255, A: 255}

	// Add the function and legend entries to the plot
	p.Add(y)
	p.Legend.Add("1.3*x+100000", y)
	p.Legend.ThumbnailWidth = 0.5 * vg.Inch

	// Set the axis ranges.  Unlike other data sets,
	// functions don't set the axis ranges automatically
	// since functions don't necessarily have a
	// finite range of x and y values.
	p.X.Min = 1000000
	p.X.Max = 3000000
	p.Y.Min = 1000000
	p.Y.Max = 10000000

	// Save the plot to a PNG file.
	if err := p.Save(4*vg.Inch, 4*vg.Inch, "file.png"); err != nil {
		return c.Status(500).JSON(fiber.Map{
			"status":  "failed",
			"message": "OOOOOOOOOOOOOOPS..........could not plot the graph",
			"error":   err.Error(),
		})
	}

	return c.Status(200).JSON(fiber.Map{
		"status":  "success",
		"message": "File named file.png successfully created",
	})

}
