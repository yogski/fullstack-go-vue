// main.go

package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	db "github.com/yogski/fullstack-go-vue/connections"
	models "github.com/yogski/fullstack-go-vue/models"
	routes "github.com/yogski/fullstack-go-vue/routes"
)

var migrateFlag bool

func init() {
	// Define the --migration flag
	flag.BoolVar(&migrateFlag, "migration", false, "Run migrations")
	flag.Parse()
}

func main() {
	db.Init()

	// handling the --migration flag. Only for first setup. Will exit without running server.
	if migrateFlag {

		// Run the migrations
		db.Migrate(
			&models.CustomerCompany{},
			&models.Customer{},
			&models.Order{},
			&models.OrderItem{},
			&models.Delivery{},
		)

		// Seed data from CSV
		db.ImportCSV()

		os.Exit(0)
	}

	if err := godotenv.Load(); err != nil {
		fmt.Println("Error loading .env file")
		panic(err)
	}
	appPort := os.Getenv("APP_PORT")

	router := gin.Default()
	routes.InitRoutes(router, db.DB)

	router.Run(":" + appPort)
}
