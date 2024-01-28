package models

// Product represents the order_items table
type Delivery struct {
	Id                int `gorm:"primary_key"`
	OrderItemId       int
	DeliveredQuantity int
}
