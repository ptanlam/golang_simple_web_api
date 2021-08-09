package services

import (
	"github.com/ptanlam/golang_simple_web_api/app/models"
	"github.com/ptanlam/golang_simple_web_api/app/repositories"
)

type CustomerService interface {
	GetAllCustomers() ([]models.Customer, error)
}

type DefaultCustomerService struct {
	repo repositories.ICustomerRepository
}

func (s DefaultCustomerService) GetAllCustomers() ([]models.Customer, error) {
	return s.repo.FindAll()
}

func NewCustomerService(repository repositories.ICustomerRepository) DefaultCustomerService {
	return DefaultCustomerService{repo: repository}
}
