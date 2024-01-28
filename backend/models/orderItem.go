package models

// Product represents the order_items table
type OrderItem struct {
	Id           int `gorm:"primary_key"`
	OrderId      int
	PricePerUnit float64
	Quantity     int
	Product      string `gorm:"not null"`
}
