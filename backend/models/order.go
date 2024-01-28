package models

import (
	"time"
)

// Product represents the main_products table
type Order struct {
	Id         int `gorm:"primary_key"`
	CreatedAt  time.Time
	OrderName  string
	CustomerId string   `gorm:"not null"`
	Customer   Customer `gorm:"foreignKey:customer_id"`
}
