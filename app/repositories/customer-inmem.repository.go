package repositories

import "github.com/ptanlam/golang_simple_web_api/app/models"

type CustomerRepositoryStub struct {
	customers []models.Customer
}

func (s CustomerRepositoryStub) FindAll() ([]models.Customer, error) {
	return s.customers, nil
}

func NewCustomerRepositoryStub() CustomerRepositoryStub {
	customers := []models.Customer{
		{
			Id:          "1234",
			Name:        "Lam Phan",
			City:        "HCM",
			ZipCode:     "70000",
			DateOfBirth: "20/03/2000",
			Status:      "1",
		},
		{
			Id:          "1235",
			Name:        "Hien Tran",
			City:        "VT",
			ZipCode:     "78000",
			DateOfBirth: "11/02/1999",
			Status:      "1",
		},
	}

	return CustomerRepositoryStub{customers: customers}
}
