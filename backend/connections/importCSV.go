package db

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
	"time"

	_ "github.com/jinzhu/gorm/dialects/postgres"
	models "github.com/yogski/fullstack-go-vue/models"
)

// Import each file and insert to DB
// todo:
func ImportCustomerCompanies() {
	wd, err := os.Getwd()
	if err != nil {
		fmt.Println("Error getting current working directory:", err)
		return
	}

	// Construct the full path to the CSV file
	filePath := wd + "/connections/Test task - Postgres - customer_companies.csv"
	file, err := os.Open(filePath)
	if err != nil {
		fmt.Println("Error opening CSV file:", err)
		return
	}
	defer file.Close()

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		fmt.Println("Error reading CSV records:", err)
		return
	}

	for i, record := range records {
		// Skip the first row (headers)
		if i == 0 {
			continue
		}

		// Assuming the CSV structure is consistent, adjust index based on your CSV columns
		inputCompanyID, _ := strconv.Atoi(record[0])
		inputCompanyName := record[1]

		company := models.CustomerCompany{
			CompanyId:   inputCompanyID,
			CompanyName: inputCompanyName,
		}

		// Insert into the database
		if err := DB.Create(&company).Error; err != nil {
			fmt.Println("Error inserting record into 'customer_companies' table:", err)
			return
		}
	}

	fmt.Println("CSV data seeded into the 'customer_companies' table")
}

func ImportCustomer() {
	wd, err := os.Getwd()
	if err != nil {
		fmt.Println("Error getting current working directory:", err)
		return
	}

	// Construct the full path to the CSV file
	filePath := wd + "/connections/Test task - Postgres - customers.csv"
	file, err := os.Open(filePath)
	if err != nil {
		fmt.Println("Error opening CSV file:", err)
		return
	}
	defer file.Close()

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		fmt.Println("Error reading CSV records:", err)
		return
	}

	for i, record := range records {
		// Skip the first row (headers)
		if i == 0 {
			continue
		}

		// Assuming the CSV structure is consistent, adjust index based on your CSV columns
		inputUserId := record[0]
		inputLogin := record[1]
		inputPassword := record[2]
		inputName := record[3]
		inputCompanyId, _ := strconv.Atoi(record[4])
		inputCreditCards := record[5]
		// user_id,login,password,name,company_id,credit_cards
		customer := models.Customer{
			UserId:      inputUserId,
			Login:       inputLogin,
			Password:    inputPassword,
			Name:        inputName,
			CompanyId:   &inputCompanyId,
			CreditCards: &inputCreditCards,
		}

		// Insert into the database
		if err := DB.Create(&customer).Error; err != nil {
			fmt.Println("Error inserting record into 'customers' table:", err)
			return
		}
	}

	fmt.Println("CSV data seeded into 'customers' table")
}

func ImportOrder() {
	wd, err := os.Getwd()
	if err != nil {
		fmt.Println("Error getting current working directory:", err)
		return
	}

	// Construct the full path to the CSV file
	filePath := wd + "/connections/Test task - Postgres - orders.csv"
	file, err := os.Open(filePath)
	if err != nil {
		fmt.Println("Error opening CSV file:", err)
		return
	}
	defer file.Close()

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		fmt.Println("Error reading CSV records:", err)
		return
	}

	for i, record := range records {
		// Skip the first row (headers)
		if i == 0 {
			continue
		}

		// Assuming the CSV structure is consistent, adjust index based on your CSV columns

		inputId, _ := strconv.Atoi(record[0])
		inputCreatedAt, _ := time.Parse("2006-01-02T15:04:05Z", record[1])
		inputOrderName := record[2]
		inputCustomerId := record[3]

		order := models.Order{
			Id:         inputId,
			CreatedAt:  inputCreatedAt,
			OrderName:  inputOrderName,
			CustomerId: inputCustomerId,
		}

		// Insert into the database
		if err := DB.Create(&order).Error; err != nil {
			fmt.Println("Error inserting record into 'orders' table:", err)
			return
		}
	}

	fmt.Println("CSV data seeded into 'orders' table")
}

func ImportOrderItem() {
	wd, err := os.Getwd()
	if err != nil {
		fmt.Println("Error getting current working directory:", err)
		return
	}

	// Construct the full path to the CSV file
	filePath := wd + "/connections/Test task - Postgres - order_items.csv"
	file, err := os.Open(filePath)
	if err != nil {
		fmt.Println("Error opening CSV file:", err)
		return
	}
	defer file.Close()

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		fmt.Println("Error reading CSV records:", err)
		return
	}

	for i, record := range records {
		// Skip the first row (headers)
		if i == 0 {
			continue
		}

		// Assuming the CSV structure is consistent, adjust index based on your CSV columns
		inputId, _ := strconv.Atoi(record[0])
		inputOrderId, _ := strconv.Atoi(record[1])
		inputPricePerunit, _ := strconv.ParseFloat(record[2], 64)
		inputQuantity, _ := strconv.Atoi(record[3])
		inputProduct := record[4]

		orderItem := models.OrderItem{
			Id:           inputId,
			OrderId:      inputOrderId,
			PricePerUnit: inputPricePerunit,
			Quantity:     inputQuantity,
			Product:      inputProduct,
		}

		// Insert into the database
		if err := DB.Create(&orderItem).Error; err != nil {
			fmt.Println("Error inserting record into 'order_items' table:", err)
			return
		}
	}

	fmt.Println("CSV data seeded into 'order_items' table")
}

func ImportDelivery() {
	wd, err := os.Getwd()
	if err != nil {
		fmt.Println("Error getting current working directory:", err)
		return
	}

	// Construct the full path to the CSV file
	filePath := wd + "/connections/Test task - Postgres - deliveries.csv"
	file, err := os.Open(filePath)
	if err != nil {
		fmt.Println("Error opening CSV file:", err)
		return
	}
	defer file.Close()

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		fmt.Println("Error reading CSV records:", err)
		return
	}

	for i, record := range records {
		// Skip the first row (headers)
		if i == 0 {
			continue
		}

		// Assuming the CSV structure is consistent, adjust index based on your CSV columns
		inputId, _ := strconv.Atoi(record[0])
		inputOrderItemId, _ := strconv.Atoi(record[1])
		inputQuantity, _ := strconv.Atoi(record[2])

		customer := models.Delivery{
			Id:                inputId,
			OrderItemId:       inputOrderItemId,
			DeliveredQuantity: inputQuantity,
		}

		// Insert into the database
		if err := DB.Create(&customer).Error; err != nil {
			fmt.Println("Error inserting record into 'customers' table:", err)
			return
		}
	}

	fmt.Println("CSV data seeded into 'customers' table")
}
