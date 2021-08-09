package repositories

import (
	"context"
	"database/sql"
	"fmt"
	"log"

	"github.com/ptanlam/golang_simple_web_api/app/models"
)

type CustomerRepository struct {
	db *sql.DB
}

func (cr *CustomerRepository) FindAll() ([]models.Customer, error) {
	ctx := context.Background()
	err := cr.db.PingContext(ctx)
	if err != nil {
		log.Fatal(err.Error())
	}
	fmt.Printf("Connected!\n")

	sqlCmd := "SELECT * FROM Customers"

	rows, err := cr.db.Query(sqlCmd)
	if err != nil {
		log.Fatal("Error while querying database: ", err.Error())
	}

	var customers []models.Customer
	for rows.Next() {
		var c models.Customer
		err := rows.Scan(&c.Id, &c.Name, &c.City, &c.DateOfBirth, &c.Status, &c.ZipCode)
		if err != nil {
			log.Fatal("Error while scanning customers: ", err.Error())
		}
		customers = append(customers, c)
	}
	return customers, nil
}

func NewCustomerRepository() CustomerRepository {
	connString := fmt.Sprintf("server=%s;user id=%s;password=%s;port=%d;database=%s;",
		"localhost", "sa", "Admin@123", 1433, "Golang_Test")

	db, err := sql.Open("sqlserver", connString)
	if err != nil {
		log.Fatal("Error creating connection pool: ", err.Error())
	}

	return CustomerRepository{db: db}
}
