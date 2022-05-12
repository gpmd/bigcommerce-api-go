package mocks

import (
	"strconv"

	"github.com/gpmd/bigcommerce-api-go"
)

type CartClient struct {
	carts       map[string]*bigcommerce.Cart
	CartContent map[string]bigcommerce.LineItem
	CartID      string
	CustomerID  int64
}

func (cm *CartClient) CreateCart(items []bigcommerce.LineItem) (*bigcommerce.Cart, error) {
	if cm.carts == nil {
		cm.carts = map[string]*bigcommerce.Cart{}
	}
	cart := bigcommerce.Cart{
		ID:         cm.CartID,
		CustomerID: cm.CustomerID,
	}
	cart.LineItems.PhysicalItems = append(cart.LineItems.PhysicalItems, items...)
	cm.carts[cart.ID] = &cart
	return &cart, nil
}

func (cm *CartClient) GetCart(cartID string) (*bigcommerce.Cart, error) {
	if cm.carts == nil {
		cm.carts = map[string]*bigcommerce.Cart{}
		return nil, bigcommerce.ErrNotFound
	}
	if cart, ok := cm.carts[cartID]; ok {
		return cart, nil
	}
	return nil, bigcommerce.ErrNotFound
}

func (cm *CartClient) CartAddItems(cartID string, items []bigcommerce.LineItem) (*bigcommerce.Cart, error) {
	if cm.carts == nil {
		cm.carts = map[string]*bigcommerce.Cart{}
		return nil, bigcommerce.ErrNotFound
	}
	if cart, ok := cm.carts[cartID]; ok {
		cart.LineItems.PhysicalItems = append(cart.LineItems.PhysicalItems, items...)
		return cart, nil
	}
	return nil, bigcommerce.ErrNotFound
}

func (cm *CartClient) CartEditItem(cartID string, item bigcommerce.LineItem) (*bigcommerce.Cart, error) {
	if cm.carts == nil {
		cm.carts = map[string]*bigcommerce.Cart{}
		return nil, bigcommerce.ErrNotFound
	}
	if cart, ok := cm.carts[cartID]; ok {
		for i, lineItem := range cart.LineItems.PhysicalItems {
			if lineItem.ID == item.ID {
				cart.LineItems.PhysicalItems[i] = item
				return cart, nil
			}
		}
	}
	return nil, bigcommerce.ErrNotFound
}

func (cm *CartClient) CartDeleteItem(cartID string, item bigcommerce.LineItem) (*bigcommerce.Cart, error) {
	if cm.carts == nil {
		cm.carts = map[string]*bigcommerce.Cart{}
		return nil, bigcommerce.ErrNotFound
	}
	if cart, ok := cm.carts[cartID]; ok {
		for i, lineItem := range cart.LineItems.PhysicalItems {
			if lineItem.ID == item.ID {
				cart.LineItems.PhysicalItems = append(cart.LineItems.PhysicalItems[:i], cart.LineItems.PhysicalItems[i+1:]...)
				return cart, nil
			}
		}
	}
	return nil, bigcommerce.ErrNotFound
}

func (cm *CartClient) CartUpdateCustomerID(cartID, customerID string) (*bigcommerce.Cart, error) {
	if cm.carts == nil {
		cm.carts = map[string]*bigcommerce.Cart{}
		return nil, bigcommerce.ErrNotFound
	}
	if cart, ok := cm.carts[cartID]; ok {
		cid, err := strconv.ParseInt(customerID, 10, 64)
		if err != nil {
			return nil, err
		}
		cart.CustomerID = cid
		return cart, nil
	}
	return nil, bigcommerce.ErrNotFound
}

func (cm *CartClient) DeleteCart(cartID string) error {
	if cm.carts == nil {
		cm.carts = map[string]*bigcommerce.Cart{}
		return bigcommerce.ErrNotFound
	}
	delete(cm.carts, cartID)
	return nil
}
