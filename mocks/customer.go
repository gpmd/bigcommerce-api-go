package mocks

import (
	"github.com/gpmd/bigcommerce-api-go"
)

type CustomerClient struct {
	CustomerID int64
	Email      string
	Password   string
	Customer   *bigcommerce.Customer
}

func (cm *CustomerClient) ValidateCredentials(email, password string) (int64, error) {
	if email == cm.Email && password == cm.Password {
		return 1, nil
	}
	return 0, bigcommerce.ErrNotFound
}

func (cm *CustomerClient) CreateAccount(customer *bigcommerce.CreateAccountPayload) (*bigcommerce.Customer, error) {
	return cm.Customer, nil
}

func (cm *CustomerClient) CustomerSetFormFields(customerID int64, formFields []bigcommerce.FormField) error {
	return nil
}
