package bigcommerce

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

// Cart is a BigCommerce cart object
type Cart struct {
	ID          string `json:"id"`
	CheckoutURL string `json:"checkout_url"`
	CustomerID  int64  `json:"customer_id"`
	ChannelID   int64  `json:"channel_id"`
	Email       string `json:"email"`
	Currency    struct {
		Code string `json:"code"`
	} `json:"currency"`
	TaxIncluded    bool         `json:"tax_included"`
	BaseAmount     float64      `json:"base_amount"`
	DiscountAmount float64      `json:"discount_amount"`
	CartAmount     float64      `json:"cart_amount"`
	Discounts      []Discount   `json:"discounts"`
	Coupons        []CartCoupon `json:"coupons"`
	LineItems      struct {
		PhysicalItems    []LineItem `json:"physical_items"`
		DigitalItems     []LineItem `json:"digital_items"`
		GiftCertificates []LineItem `json:"gift_certificates"`
		CustomItems      []LineItem `json:"custom_items"`
	} `json:"line_items"`
	CreatedTime  time.Time `json:"created_time"`
	UpdatedTime  time.Time `json:"updated_time"`
	RedirectUrls struct {
		CartURL             string `json:"cart_url"`
		CheckoutURL         string `json:"checkout_url"`
		EmbeddedCheckoutURL string `json:"embedded_checkout_url"`
	} `json:"redirect_urls"`
	Locale string `json:"locale"`
}

// LineItem is a BigCommerce line item object for cart
type LineItem struct {
	ID                string     `json:"id"`
	ParentID          int64      `json:"parent_id"`
	VariantID         int64      `json:"variant_id"`
	ProductID         int64      `json:"product_id"`
	Sku               string     `json:"sku"`
	Name              string     `json:"name"`
	URL               string     `json:"url"`
	Quantity          float64    `json:"quantity"`
	Taxable           bool       `json:"taxable"`
	ImageURL          string     `json:"image_url"`
	Discounts         []Discount `json:"discounts"`
	Coupons           []Coupon   `json:"coupons"`
	DiscountAmount    float64    `json:"discount_amount"`
	CouponAmount      float64    `json:"coupon_amount"`
	ListPrice         float64    `json:"list_price"`
	SalePrice         float64    `json:"sale_price"`
	ExtendedListPrice float64    `json:"extended_list_price"`
	ExtendedSalePrice float64    `json:"extended_sale_price"`
	IsRequireShipping bool       `json:"is_require_shipping"`
	IsMutable         bool       `json:"is_mutable"`
}

type CartURLs struct {
	CartURL             string `json:"cart_url"`
	CheckoutURL         string `json:"checkout_url"`
	EmbeddedCheckoutURL string `json:"embedded_checkout_url"`
}

// CreateCart creates a new cart in BigCommerce and returns it
func (bc *Client) CreateCart(items []LineItem) (*Cart, error) {
	var body []byte
	body, _ = json.Marshal(map[string]interface{}{
		"channel_id": bc.ChannelID,
		"line_items": items,
	})
	req := bc.getAPIRequest(http.MethodPost, "/v3/carts?include=redirect_urls", bytes.NewReader(body))
	res, err := bc.HTTPClient.Do(req)
	if err != nil {
		return nil, err
	}
	b, err := processBody(res)
	if err != nil {
		return nil, err
	}
	var cartResponse struct {
		Data Cart `json:"data"`
		Meta struct {
		} `json:"meta"`
	}
	err = json.Unmarshal(b, &cartResponse)
	if err != nil {
		return nil, err
	}
	return &cartResponse.Data, nil
}

// GetCart gets a cart by ID from BigCommerce and returns it
func (bc *Client) GetCart(cartID string) (*Cart, error) {
	req := bc.getAPIRequest(http.MethodGet, "/v3/carts/"+cartID+"?include=redirect_urls", nil)
	res, err := bc.HTTPClient.Do(req)
	if err != nil {
		return nil, err
	}
	b, err := processBody(res)
	if err != nil {
		return nil, err
	}
	var cartResponse struct {
		Data Cart `json:"data"`
		Meta struct {
		} `json:"meta"`
	}
	err = json.Unmarshal(b, &cartResponse)
	if err != nil {
		return nil, err
	}
	return &cartResponse.Data, nil
}

// CartAddItem adds line items to a cart
func (bc *Client) CartAddItems(cartID string, items []LineItem) (*Cart, error) {
	var body []byte
	body, _ = json.Marshal(map[string]interface{}{
		"line_items": items,
	})
	req := bc.getAPIRequest(http.MethodPost, "/v3/carts/"+cartID+"/items?include=redirect_urls", bytes.NewReader(body))
	res, err := bc.HTTPClient.Do(req)
	if err != nil {
		return nil, err
	}
	b, err := processBody(res)
	if err != nil {
		return nil, err
	}
	var cartResponse struct {
		Data Cart `json:"data"`
		Meta struct {
		} `json:"meta"`
	}
	err = json.Unmarshal(b, &cartResponse)
	if err != nil {
		return nil, err
	}
	return &cartResponse.Data, nil
}

// EditItem edits a line item in a cart, returns the updated cart
// Arguments:
// 		cartID: the cart ID
// 		item: the line item to edit. Must have an ID, quantity, and product ID
func (bc *Client) CartEditItem(cartID string, item LineItem) (*Cart, error) {
	var body []byte
	body, _ = json.Marshal(map[string]interface{}{
		"line_item": item,
	})
	req := bc.getAPIRequest(http.MethodPut, "/v3/carts/"+cartID+"/items/"+item.ID+"?include=redirect_urls", bytes.NewReader(body))
	res, err := bc.HTTPClient.Do(req)
	if err != nil {
		return nil, err
	}
	b, err := processBody(res)
	if err != nil {
		return nil, fmt.Errorf("%s", string(b))
	}
	return bc.GetCart(cartID)
}

// DeleteItem deletes a line item from a cart, returns the updated cart
// Arguments:
// 		cartID: the cart ID
// 		item: the line item, must have an existing line item ID
// returns nil for empty cart
func (bc *Client) CartDeleteItem(cartID string, item LineItem) (*Cart, error) {
	req := bc.getAPIRequest(http.MethodDelete, "/v3/carts/"+cartID+"/items/"+item.ID+"?include=redirect_urls", nil)
	res, err := bc.HTTPClient.Do(req)
	if err != nil {
		return nil, err
	}
	if res.StatusCode == 204 {
		return nil, nil
	}
	_, err = processBody(res)
	if err != nil {
		return nil, err
	}
	return bc.GetCart(cartID)
}

// CartUpdateCustomerID updates the customer ID for a cart
// Arguments:
// cartID: the BigCommerce cart ID
// customerID: the new BigCommerce customer ID
func (bc *Client) CartUpdateCustomerID(cartID, customerID string) (*Cart, error) {
	req := bc.getAPIRequest(http.MethodPut, "/v3/carts/"+cartID+"?include=redirect_urls",
		bytes.NewReader([]byte(fmt.Sprintf(`{"customer_id": %s}`, customerID))))
	res, err := bc.HTTPClient.Do(req)
	if err != nil {
		return nil, err
	}
	b, err := processBody(res)
	if err != nil {
		return nil, err
	}
	var cartResponse struct {
		Data Cart `json:"data"`
	}
	err = json.Unmarshal(b, &cartResponse)
	if err != nil {
		return nil, err
	}
	return &cartResponse.Data, nil
}

// DeleteCart deletes a cart by ID from BigCommerce
func (bc *Client) DeleteCart(cartID string) error {
	req := bc.getAPIRequest(http.MethodDelete, "/v3/carts/"+cartID, nil)
	res, err := bc.HTTPClient.Do(req)
	if err != nil {
		return err
	}
	if res.StatusCode != http.StatusNoContent {
		return fmt.Errorf("unexpected status code: %d", res.StatusCode)
	}
	return nil
}
