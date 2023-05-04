package config

import (
	"fmt"
	"log"
	"os"

	"github.com/BEERBAL-HASH/Rule-Engine/models"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func loadEnv() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file", err.Error())
	}
}

func ConnectDB() {
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s sslmode=%s", PostgresHost, PostgresUser, PostgresPass, PostgresName, PostgresSSLMode,
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("Failed to connect to the database server! \n", err.Error())
		panic(err)
	}
	fmt.Println("Database connected successfully")
	db.AutoMigrate(&models.Employee{})
	DBconn = db
}

var (
	DBconn *gorm.DB
)
var PostgresHost string
var PostgresPort string
var PostgresUser string
var PostgresPass string
var PostgresName string
var PostgresSSLMode string

func init() {
	loadEnv() //load .env file

	PostgresHost = os.Getenv("POSTGRES_HOST")
	PostgresPort = os.Getenv("POSTGRES_PORT")
	PostgresUser = os.Getenv("POSTGRES_USER")
	PostgresName = os.Getenv("POSTGRES_NAME")
	PostgresPass = os.Getenv("POSTGRES_PASS")
	PostgresSSLMode = os.Getenv("POSTGRES_SSLMODE")

}
