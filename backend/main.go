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

	gg := gin.Default()
	gg.Use(CORSMiddleware())
	routes.InitRoutes(gg, db.DB)

	gg.Run(":" + appPort)
}

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {

		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Credentials", "true")
		c.Header("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Header("Access-Control-Allow-Methods", "POST,HEAD,PATCH, OPTIONS, GET, PUT")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}
