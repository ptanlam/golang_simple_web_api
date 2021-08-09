package repositories

import "github.com/ptanlam/golang_simple_web_api/app/models"

type ICustomerRepository interface {
	FindAll() ([]models.Customer, error)
}
