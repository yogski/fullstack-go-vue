package models

type Customer struct {
	UserId          string `gorm:"primary_key"`
	Login           string `gorm:"not null"`
	Password        string `gorm:"not null"`
	Name            string `gorm:"not null"`
	CompanyId       *int
	CreditCards     *string
	Orders          []Order         `gorm:"foreignKey:customer_id"`
	CustomerCompany CustomerCompany `gorm:"foreignKey:company_id;references:company_id"`
}
