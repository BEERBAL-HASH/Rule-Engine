package models

import "gorm.io/gorm"

type Employee struct {
	gorm.Model
	EmployeeName   string `json:"EmployeeName"`
	DepartmentName string `json:"DepartmentName"`
	Profile        string `json:"Profile"`
	Age            int    `json:"Age"`
	Salary         int    `json:"Salary"`
}
