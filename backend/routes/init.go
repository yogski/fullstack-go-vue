package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

func InitRoutes(router *gin.Engine, dbInstance *gorm.DB) {
	db = dbInstance

	router.GET("/orders", getOrders)

	// Add more routes
}
