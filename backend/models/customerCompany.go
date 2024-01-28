package models

type CustomerCompany struct {
	CompanyId   int    `gorm:"primary_key"`
	CompanyName string `gorm:"not null"`
}
