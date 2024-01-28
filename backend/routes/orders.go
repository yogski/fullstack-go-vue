package routes

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	models "github.com/yogski/fullstack-go-vue/models"
)

var db *gorm.DB

// Failed attempt of using proper model. insufficient time to resolve everything.
func getOrdersOld(c *gin.Context) {
	var orders []models.Order
	if err := db.Debug().Preload("Customer").Preload("Customer.CustomerCompany").Find(&orders).Error; err != nil {
		c.JSON(500, gin.H{"error": "Failed to fetch orders"})
		return
	}

	c.JSON(200, orders)
}

func getOrders(c *gin.Context) {
	var orders []models.OrderReport
	rawBuildQuery := `
	SELECT
		t1.order_name,
		t3.company_name,
		t2.name as "customer_name",
		t1.created_at as "order_date",
		SUM(t4.price_per_unit * t4.quantity) as "total_amount",
		SUM(t5.delivered_quantity) as "delivered_quantity"
	FROM orders t1
	INNER JOIN customers t2 ON t1.customer_id = t2.user_id
	INNER JOIN customer_companies t3 ON t3.company_id = t2.company_id
	INNER JOIN order_items t4 ON t4.order_id = t1.id
	LEFT JOIN deliveries t5 ON t5.order_item_id = t4.id
	`
	// this is an ugly and dangerous raw builder query
	// because I haven't work with Gorm for so long
	rawBuildQuery += " WHERE "

	if startDate := c.Query("start_date"); startDate != "" {
		rawBuildQuery += fmt.Sprintf("t1.created_at >= '%s' AND ", startDate)
	} else {
		rawBuildQuery += "TRUE AND "
	}

	if endDate := c.Query("end_date"); endDate != "" {
		rawBuildQuery += fmt.Sprintf(" t1.created_at <= '%s' AND ", endDate)
	} else {
		rawBuildQuery += "TRUE AND "
	}

	if textQuery := c.Query("q"); textQuery != "" {
		rawBuildQuery += fmt.Sprintf(" t4.product LIKE '%%%s%%' OR t1.order_name LIKE '%%%s%%' AND ", textQuery, textQuery)
	} else {
		rawBuildQuery += "TRUE AND "
	}

	rawBuildQuery = rawBuildQuery[:len(rawBuildQuery)-4] // This removes AND with space
	rawBuildQuery += " GROUP BY t1.id, t1.order_name, t3.company_name, t2.name, t1.created_at "

	if limit := c.Query("limit"); limit != "" {
		rawBuildQuery += fmt.Sprintf(" LIMIT %s ", limit)
	}

	if offset := c.Query("offset"); offset != "" {
		rawBuildQuery += fmt.Sprintf(" OFFSET %s ", offset)
	}

	if err := db.Raw(rawBuildQuery).Scan(&orders).Error; err != nil {
		c.JSON(500, gin.H{"error": "Failed to fetch orders"})
		return
	}

	c.JSON(200, orders)
}

// Define other order-related routes and functions as needed
