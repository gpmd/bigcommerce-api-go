package bigcommerce

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"
)

type Order struct {
	ID                                      int64        `json:"id"`
	CustomerID                              int64        `json:"customer_id"`
	DateCreated                             string       `json:"date_created"`
	DateModified                            string       `json:"date_modified"`
	DateShipped                             string       `json:"date_shipped"`
	StatusID                                int64        `json:"status_id"`
	Status                                  string       `json:"status"`
	SubtotalExTax                           string       `json:"subtotal_ex_tax"`
	SubtotalIncTax                          string       `json:"subtotal_inc_tax"`
	SubtotalTax                             string       `json:"subtotal_tax"`
	BaseShippingCost                        string       `json:"base_shipping_cost"`
	ShippingCostExTax                       string       `json:"shipping_cost_ex_tax"`
	ShippingCostIncTax                      string       `json:"shipping_cost_inc_tax"`
	ShippingCostTax                         string       `json:"shipping_cost_tax"`
	ShippingCostTaxClassID                  int64        `json:"shipping_cost_tax_class_id"`
	BaseHandlingCost                        string       `json:"base_handling_cost"`
	HandlingCostExTax                       string       `json:"handling_cost_ex_tax"`
	HandlingCostIncTax                      string       `json:"handling_cost_inc_tax"`
	HandlingCostTax                         string       `json:"handling_cost_tax"`
	HandlingCostTaxClassID                  int64        `json:"handling_cost_tax_class_id"`
	BaseWrappingCost                        string       `json:"base_wrapping_cost"`
	WrappingCostExTax                       string       `json:"wrapping_cost_ex_tax"`
	WrappingCostIncTax                      string       `json:"wrapping_cost_inc_tax"`
	WrappingCostTax                         string       `json:"wrapping_cost_tax"`
	WrappingCostTaxClassID                  int64        `json:"wrapping_cost_tax_class_id"`
	TotalExTax                              string       `json:"total_ex_tax"`
	TotalIncTax                             string       `json:"total_inc_tax"`
	TotalTax                                string       `json:"total_tax"`
	ItemsTotal                              int          `json:"items_total"`
	ItemsShipped                            int          `json:"items_shipped"`
	PaymentMethod                           string       `json:"payment_method"`
	PaymentProviderID                       string       `json:"payment_provider_id"`
	PaymentStatus                           string       `json:"payment_status"`
	RefundedAmount                          string       `json:"refunded_amount"`
	OrderIsDigital                          bool         `json:"order_is_digital"`
	StoreCreditAmount                       string       `json:"store_credit_amount"`
	GiftCertificateAmount                   string       `json:"gift_certificate_amount"`
	IPAddress                               string       `json:"ip_address"`
	IPAddressV6                             string       `json:"ip_address_v6"`
	GeoipCountry                            string       `json:"geoip_country"`
	GeoipCountryIso2                        string       `json:"geoip_country_iso2"`
	CurrencyID                              int64        `json:"currency_id"`
	CurrencyCode                            string       `json:"currency_code"`
	CurrencyExchangeRate                    string       `json:"currency_exchange_rate"`
	DefaultCurrencyID                       int64        `json:"default_currency_id"`
	DefaultCurrencyCode                     string       `json:"default_currency_code"`
	StaffNotes                              string       `json:"staff_notes"`
	CustomerMessage                         string       `json:"customer_message"`
	DiscountAmount                          string       `json:"discount_amount"`
	CouponDiscount                          string       `json:"coupon_discount"`
	ShippingAddressCount                    int          `json:"shipping_address_count"`
	IsDeleted                               bool         `json:"is_deleted"`
	EbayOrderID                             string       `json:"ebay_order_id"`
	CartID                                  string       `json:"cart_id"`
	BillingAddress                          OrderAddress `json:"billing_address"`
	IsEmailOptIn                            bool         `json:"is_email_opt_in"`
	CreditCardType                          interface{}  `json:"credit_card_type"`
	OrderSource                             string       `json:"order_source"`
	ChannelID                               int64        `json:"channel_id"`
	ExternalSource                          string       `json:"external_source"`
	Products                                interface{}  `json:"products"`
	ShippingAddresses                       interface{}  `json:"shipping_addresses"`
	Coupons                                 interface{}  `json:"coupons"`
	ExternalID                              interface{}  `json:"external_id"`
	ExternalMerchantID                      interface{}  `json:"external_merchant_id"`
	TaxProviderID                           string       `json:"tax_provider_id"`
	StoreDefaultCurrencyCode                string       `json:"store_default_currency_code"`
	StoreDefaultToTransactionalExchangeRate string       `json:"store_default_to_transactional_exchange_rate"`
	CustomStatus                            string       `json:"custom_status"`
	CustomerLocale                          string       `json:"customer_locale"`
}

type OrderAddress struct {
	FirstName   string        `json:"first_name"`
	LastName    string        `json:"last_name"`
	Company     string        `json:"company"`
	Street1     string        `json:"street_1"`
	Street2     string        `json:"street_2"`
	City        string        `json:"city"`
	State       string        `json:"state"`
	Zip         string        `json:"zip"`
	Country     string        `json:"country"`
	CountryIso2 string        `json:"country_iso2"`
	Phone       string        `json:"phone"`
	Email       string        `json:"email"`
	FormFields  []interface{} `json:"form_fields"`
}

type OrderProduct struct {
	ID                   int64             `json:"id"`
	OrderID              int64             `json:"order_id"`
	ProductID            int64             `json:"product_id"`
	OrderAddressID       int64             `json:"order_address_id"`
	Name                 string            `json:"name"`
	NameCustomer         string            `json:"name_customer"`
	NameMerchant         string            `json:"name_merchant"`
	Sku                  string            `json:"sku"`
	Upc                  string            `json:"upc"`
	Type                 string            `json:"type"`
	BasePrice            string            `json:"base_price"`
	PriceExTax           string            `json:"price_ex_tax"`
	PriceIncTax          string            `json:"price_inc_tax"`
	PriceTax             string            `json:"price_tax"`
	BaseTotal            string            `json:"base_total"`
	TotalExTax           string            `json:"total_ex_tax"`
	TotalIncTax          string            `json:"total_inc_tax"`
	TotalTax             string            `json:"total_tax"`
	Weight               string            `json:"weight"`
	Quantity             int               `json:"quantity"`
	BaseCostPrice        string            `json:"base_cost_price"`
	CostPriceIncTax      string            `json:"cost_price_inc_tax"`
	CostPriceExTax       string            `json:"cost_price_ex_tax"`
	CostPriceTax         string            `json:"cost_price_tax"`
	IsRefunded           bool              `json:"is_refunded"`
	QuantityRefunded     int               `json:"quantity_refunded"`
	RefundAmount         string            `json:"refund_amount"`
	ReturnID             int64             `json:"return_id"`
	WrappingName         string            `json:"wrapping_name"`
	BaseWrappingCost     string            `json:"base_wrapping_cost"`
	WrappingCostExTax    string            `json:"wrapping_cost_ex_tax"`
	WrappingCostIncTax   string            `json:"wrapping_cost_inc_tax"`
	WrappingCostTax      string            `json:"wrapping_cost_tax"`
	WrappingMessage      string            `json:"wrapping_message"`
	QuantityShipped      int               `json:"quantity_shipped"`
	FixedShippingCost    string            `json:"fixed_shipping_cost"`
	EbayItemID           string            `json:"ebay_item_id"`
	EbayTransactionID    string            `json:"ebay_transaction_id"`
	OptionSetID          int64             `json:"option_set_id"`
	ParentOrderProductID interface{}       `json:"parent_order_product_id"`
	IsBundledProduct     bool              `json:"is_bundled_product"`
	BinPickingNumber     string            `json:"bin_picking_number"`
	ExternalID           interface{}       `json:"external_id"`
	FulfillmentSource    string            `json:"fulfillment_source"`
	AppliedDiscounts     []ProductDiscount `json:"applied_discounts"`
	ProductOptions       []ProductOption   `json:"product_options"`
	ConfigurableFields   []interface{}     `json:"configurable_fields"`
	EventName            interface{}       `json:"event_name"`
	EventDate            interface{}       `json:"event_date"`
}

type ProductDiscount struct {
	ID     string      `json:"id"`
	Amount string      `json:"amount"`
	Name   string      `json:"name"`
	Code   interface{} `json:"code"`
	Target string      `json:"target"`
}

type ProductOption struct {
	ID                   int64  `json:"id"`
	OptionID             int64  `json:"option_id"`
	OrderProductID       int64  `json:"order_product_id"`
	ProductOptionID      int64  `json:"product_option_id"`
	DisplayName          string `json:"display_name"`
	DisplayNameCustomer  string `json:"display_name_customer"`
	DisplayNameMerchant  string `json:"display_name_merchant"`
	DisplayValue         string `json:"display_value"`
	DisplayValueCustomer string `json:"display_value_customer"`
	DisplayValueMerchant string `json:"display_value_merchant"`
	Value                string `json:"value"`
	Type                 string `json:"type"`
	Name                 string `json:"name"`
	DisplayStyle         string `json:"display_style"`
}

type OrderShippingAddress struct {
	ID                     int64         `json:"id"`
	OrderID                int64         `json:"order_id"`
	FirstName              string        `json:"first_name"`
	LastName               string        `json:"last_name"`
	Company                string        `json:"company"`
	Street1                string        `json:"street_1"`
	Street2                string        `json:"street_2"`
	City                   string        `json:"city"`
	Zip                    string        `json:"zip"`
	Country                string        `json:"country"`
	CountryIso2            string        `json:"country_iso2"`
	State                  string        `json:"state"`
	Email                  string        `json:"email"`
	Phone                  string        `json:"phone"`
	ItemsTotal             int           `json:"items_total"`
	ItemsShipped           int           `json:"items_shipped"`
	ShippingMethod         string        `json:"shipping_method"`
	BaseCost               string        `json:"base_cost"`
	CostExTax              string        `json:"cost_ex_tax"`
	CostIncTax             string        `json:"cost_inc_tax"`
	CostTax                string        `json:"cost_tax"`
	CostTaxClassID         int64         `json:"cost_tax_class_id"`
	BaseHandlingCost       string        `json:"base_handling_cost"`
	HandlingCostExTax      string        `json:"handling_cost_ex_tax"`
	HandlingCostIncTax     string        `json:"handling_cost_inc_tax"`
	HandlingCostTax        string        `json:"handling_cost_tax"`
	HandlingCostTaxClassID int64         `json:"handling_cost_tax_class_id"`
	ShippingZoneID         int64         `json:"shipping_zone_id"`
	ShippingZoneName       string        `json:"shipping_zone_name"`
	ShippingQuotes         interface{}   `json:"shipping_quotes"`
	FormFields             []interface{} `json:"form_fields"`
}

type OrderCoupon struct {
	ID       int64  `json:"id"`
	CouponID int64  `json:"coupon_id"`
	OrderID  int64  `json:"order_id"`
	Code     string `json:"code"`
	Amount   int    `json:"amount"`
	Type     int    `json:"type"`
	Discount int    `json:"discount"`
}

// GetOrders returns all orders using filters
// filters: request query parameters for BigCommerce orders endpoint, for example {"customer_id": "41"}
func (bc *Client) GetOrders(filters map[string]string) ([]Order, error) {
	params := []string{}
	for k, v := range filters {
		params = append(params, fmt.Sprintf("%s=%s", k, v))
	}
	url := "/v2/orders?" + strings.Join(params, "&")

	req := bc.getAPIRequest(http.MethodGet, url, nil)
	res, err := bc.HTTPClient.Do(req)
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()
	body, err := processBody(res)
	if err != nil {
		if res.StatusCode == http.StatusNoContent {
			return []Order{}, nil
		}
		return nil, err
	}

	var orders []Order
	err = json.Unmarshal(body, &orders)
	if err != nil {
		return nil, err
	}
	return orders, nil
}

// GetOrder returns a given order
// filters: request query parameters for BigCommerce orders endpoint, for example {"customer_id": "41"}
func (bc *Client) GetOrder(orderID int64) (*Order, error) {
	url := "/v2/orders/" + strconv.FormatInt(orderID, 10)

	req := bc.getAPIRequest(http.MethodGet, url, nil)
	res, err := bc.HTTPClient.Do(req)
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()
	body, err := processBody(res)
	if err != nil {
		return nil, err
	}

	var order Order
	err = json.Unmarshal(body, &order)
	if err != nil {
		return nil, err
	}
	products, err := bc.GetOrderProducts(orderID)
	if err != nil {
		return &order, nil // well, we got the order, but we can't get the products
	}
	order.Products = products // this is why we used interface{} for products instead of OrderResource
	addresses, err := bc.GetOrderShippingAddresses(orderID)
	if err != nil {
		return &order, nil // well, we got the order, but we can't get the addresses
	}
	order.ShippingAddresses = addresses
	coupons, err := bc.GetOrderCoupons(orderID)
	if err != nil {
		return &order, nil // well, we got the order, but we can't get the coupons
	}
	order.Coupons = coupons
	return &order, nil
}

// GetOrderProducts returns all products for a given order
func (bc *Client) GetOrderProducts(orderID int64) ([]OrderProduct, error) {
	url := "/v2/orders/" + strconv.FormatInt(orderID, 10) + "/products"

	req := bc.getAPIRequest(http.MethodGet, url, nil)
	res, err := bc.HTTPClient.Do(req)
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()
	body, err := processBody(res)
	if err != nil {
		return nil, err
	}

	var products []OrderProduct
	err = json.Unmarshal(body, &products)
	if err != nil {
		return nil, err
	}
	return products, nil
}

// GetOrderShippingAddresses returns all shipping addresses for a given order
func (bc *Client) GetOrderShippingAddresses(orderID int64) ([]OrderShippingAddress, error) {
	url := "/v2/orders/" + strconv.FormatInt(orderID, 10) + "/shipping_addresses"

	req := bc.getAPIRequest(http.MethodGet, url, nil)
	res, err := bc.HTTPClient.Do(req)
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()
	body, err := processBody(res)
	if err != nil {
		return nil, err
	}

	var addresses []OrderShippingAddress
	err = json.Unmarshal(body, &addresses)
	if err != nil {
		return nil, err
	}
	for i := range addresses {
		addresses[i].ShippingQuotes = nil // we don't need this
	}
	return addresses, nil
}

// GetOrderCoupons returns all coupons for a given order
func (bc *Client) GetOrderCoupons(orderID int64) ([]OrderCoupon, error) {
	url := "/v2/orders/" + strconv.FormatInt(orderID, 10) + "/coupons"

	req := bc.getAPIRequest(http.MethodGet, url, nil)
	res, err := bc.HTTPClient.Do(req)
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()
	body, err := processBody(res)
	if err != nil {
		return nil, err
	}

	var coupons []OrderCoupon
	err = json.Unmarshal(body, &coupons)
	if err != nil {
		return nil, err
	}
	return coupons, nil
}
