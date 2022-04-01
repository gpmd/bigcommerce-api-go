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
	CheckoutURL string `json:"checkout_url,omitempty"`
	CustomerID  int64  `json:"customer_id,omitempty"`
	ChannelID   int64  `json:"channel_id,omitempty"`
	Email       string `json:"email,omitempty"`
	Currency    struct {
		Code string `json:"code,omitempty"`
	} `json:"currency,omitempty"`
	TaxIncluded    bool       `json:"tax_included,omitempty"`
	BaseAmount     float64    `json:"base_amount,omitempty"`
	DiscountAmount float64    `json:"discount_amount,omitempty"`
	CartAmount     float64    `json:"cart_amount,omitempty"`
	Discounts      []Discount `json:"discounts,omitempty"`
	Coupons        []Coupon   `json:"coupons,omitempty"`
	LineItems      struct {
		PhysicalItems    []LineItem `json:"physical_items,omitempty"`
		DigitalItems     []LineItem `json:"digital_items,omitempty"`
		GiftCertificates []LineItem `json:"gift_certificates,omitempty"`
		CustomItems      []LineItem `json:"custom_items,omitempty"`
	} `json:"line_items"`
	CreatedTime  time.Time `json:"created_time,omitempty"`
	UpdatedTime  time.Time `json:"updated_time,omitempty"`
	RedirectUrls struct {
		CartURL             string `json:"cart_url"`
		CheckoutURL         string `json:"checkout_url"`
		EmbeddedCheckoutURL string `json:"embedded_checkout_url"`
	} `json:"redirect_urls,omitempty"`
	Locale string `json:"locale,omitempty"`
}

// LineItem is a BigCommerce line item object for cart
type LineItem struct {
	ID                string     `json:"id,omitempty"`
	ParentID          int64      `json:"parent_id,omitempty"`
	VariantID         int64      `json:"variant_id,omitempty"`
	ProductID         int64      `json:"product_id,omitempty"`
	Sku               string     `json:"sku,omitempty"`
	Name              string     `json:"name,omitempty"`
	URL               string     `json:"url,omitempty"`
	Quantity          float64    `json:"quantity,omitempty"`
	Taxable           bool       `json:"taxable,omitempty"`
	ImageURL          string     `json:"image_url,omitempty"`
	Discounts         []Discount `json:"discounts,omitempty"`
	Coupons           []Coupon   `json:"coupons,omitempty"`
	DiscountAmount    float64    `json:"discount_amount,omitempty"`
	CouponAmount      float64    `json:"coupon_amount,omitempty"`
	ListPrice         float64    `json:"list_price,omitempty"`
	SalePrice         float64    `json:"sale_price,omitempty"`
	ExtendedListPrice float64    `json:"extended_list_price,omitempty"`
	ExtendedSalePrice float64    `json:"extended_sale_price,omitempty"`
	IsRequireShipping bool       `json:"is_require_shipping,omitempty"`
	IsMutable         bool       `json:"is_mutable,omitempty"`
}

type CartURLs struct {
	CartURL             string `json:"cart_url,omitempty"`
	CheckoutURL         string `json:"checkout_url,omitempty"`
	EmbeddedCheckoutURL string `json:"embedded_checkout_url,omitempty"`
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
		Data Cart `json:"data,omitempty"`
		Meta struct {
		} `json:"meta,omitempty"`
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
		Data Cart `json:"data,omitempty"`
		Meta struct {
		} `json:"meta,omitempty"`
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
		Data Cart `json:"data,omitempty"`
		Meta struct {
		} `json:"meta,omitempty"`
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
func (bc *Client) CartDeleteItem(cartID string, item LineItem) (*Cart, error) {
	req := bc.getAPIRequest(http.MethodDelete, "/v3/carts/"+cartID+"/items/"+item.ID+"?include=redirect_urls", nil)
	res, err := bc.HTTPClient.Do(req)
	if err != nil {
		return nil, err
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
		Data Cart `json:"data,omitempty"`
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
