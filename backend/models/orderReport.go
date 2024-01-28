package models

import "time"

type OrderReport struct {
	OrderName         string    `json:"order_name"`
	CompanyName       string    `json:"company_name"`
	CustomerName      string    `json:"customer_name"`
	OrderDate         time.Time `json:"order_date"`
	TotalAmount       float64   `json:"total_amount"`
	DeliveredQuantity int       `json:"delivered_quantity"`
}
