package mocks

import "github.com/gpmd/bigcommerce-api-go"

type CartClient struct {
}

func (cm *CartClient) CreateCart(items []bigcommerce.LineItem) (*bigcommerce.Cart, error) {
	return nil, nil
}

func (cm *CartClient) GetCart(cartID string) (*bigcommerce.Cart, error) {
	return nil, nil
}

func (cm *CartClient) CartAddItems(cartID string, items []bigcommerce.LineItem) (*bigcommerce.Cart, error) {
	return nil, nil
}

func (cm *CartClient) CartEditItem(cartID string, item bigcommerce.LineItem) (*bigcommerce.Cart, error) {
	return nil, nil
}

func (cm *CartClient) CartDeleteItem(cartID string, item bigcommerce.LineItem) (*bigcommerce.Cart, error) {
	return nil, nil
}

func (cm *CartClient) CartUpdateCustomerID(cartID, customerID string) (*bigcommerce.Cart, error) {
	return nil, nil
}

func (cm *CartClient) ValidateCredentials(email, password string) (int64, error) {
	return 0, nil
}
