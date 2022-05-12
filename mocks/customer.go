package mocks

import (
	"github.com/gpmd/bigcommerce-api-go"
)

type CustomerClient struct {
	CustomerID int64
	Email      string
	Password   string
	Customer   *bigcommerce.Customer
	FormFields []bigcommerce.FormField
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

func (cm *CustomerClient) CustomerGetFormFields(customerID int64) ([]bigcommerce.FormField, error) {
	return cm.FormFields, nil
}

func (cm *CustomerClient) GetCustomerByEmail(email string) (*bigcommerce.Customer, error) {
	return cm.Customer, nil
}

func (cm *CustomerClient) GetCustomerByID(customerID int64) (*bigcommerce.Customer, error) {
	return cm.Customer, nil
}

func (cm *CustomerClient) SaveAccount(customer *bigcommerce.SaveAccountPayload) (*bigcommerce.Customer, error) {
	return cm.Customer, nil
}
