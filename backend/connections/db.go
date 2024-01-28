package db

import (
	"fmt"
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/joho/godotenv"
)

// DB represents the database connection
var DB *gorm.DB

// Init initializes the database connection
func Init() {
	// Load environment variables from .env file
	if err := godotenv.Load(); err != nil {
		fmt.Println("Error loading .env file")
		panic(err)
	}

	// Retrieve environment variables
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")
	dbSSLMode := os.Getenv("DB_SSL_MODE")

	// Create database connection string
	dbConnectionString := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=%s", dbUser, dbPassword, dbName, dbSSLMode)

	var err error
	DB, err = gorm.Open("postgres", dbConnectionString)
	if err != nil {
		fmt.Println("Failed to connect to the database")
		panic(err)
	}

	// Other database configurations, if needed

	fmt.Println("Connected to the database")
}

// Migrate creates tables in the database
func Migrate(modelsList ...interface{}) {
	// Automigrate all provided models
	DB.AutoMigrate(modelsList...)

	fmt.Println("Database migrated")
}

func ImportCSV() {
	ImportCustomerCompanies()
	ImportCustomer()
	ImportOrder()
	ImportOrderItem()
	ImportDelivery()
}
