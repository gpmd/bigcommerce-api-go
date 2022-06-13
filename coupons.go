package bigcommerce

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

type Coupon struct {
	ID          int64  `json:"id"`
	Name        string `json:"name"`
	Type        string `json:"type"`
	Amount      string `json:"amount"`
	MinPurchase string `json:"min_purchase"`
	Expires     string `json:"expires"`
	Enabled     bool   `json:"enabled"`
	Code        string `json:"code"`
	AppliesTo   struct {
		Entity string  `json:"entity"`
		Ids    []int64 `json:"ids"`
	} `json:"applies_to"`
	NumUses            int           `json:"num_uses"`
	MaxUses            int           `json:"max_uses"`
	MaxUsesPerCustomer int           `json:"max_uses_per_customer"`
	RestrictedTo       []interface{} `json:"restricted_to"`
	ShippingMethods    struct {
	} `json:"shipping_methods"`
	DateCreated string `json:"date_created"`
}

func (bc *Client) CreateCoupon(coupon Coupon) (*Coupon, error) {
	var body []byte
	body, _ = json.Marshal(coupon)
	req := bc.getAPIRequest(http.MethodPost, "/v3/coupons", bytes.NewReader(body))
	res, err := bc.HTTPClient.Do(req)
	if err != nil {
		return nil, err
	}
	b, err := processBody(res)
	if err != nil {
		return nil, err
	}
	var couponResponse struct {
		Data Coupon `json:"data"`
		Meta struct {
		} `json:"meta"`
	}
	err = json.Unmarshal(b, &couponResponse)
	if err != nil {
		return nil, err
	}
	return &couponResponse.Data, nil
}

func (bc *Client) GetCoupon(couponID int64) (*Coupon, error) {
	req := bc.getAPIRequest(http.MethodGet, "/v3/coupons/"+strconv.FormatInt(couponID, 10), nil)
	res, err := bc.HTTPClient.Do(req)
	if err != nil {
		return nil, err
	}
	b, err := processBody(res)
	if err != nil {
		return nil, err
	}
	var couponResponse struct {
		Data Coupon `json:"data"`
		Meta struct {
		} `json:"meta"`
	}
	err = json.Unmarshal(b, &couponResponse)
	if err != nil {
		return nil, err
	}
	return &couponResponse.Data, nil
}

func (bc *Client) UpdateCoupon(couponID int64, coupon Coupon) (*Coupon, error) {
	var body []byte
	body, _ = json.Marshal(coupon)
	req := bc.getAPIRequest(http.MethodPut, "/v3/coupons/"+strconv.FormatInt(couponID, 10), bytes.NewReader(body))
	res, err := bc.HTTPClient.Do(req)
	if err != nil {
		return nil, err
	}
	b, err := processBody(res)
	if err != nil {
		return nil, err
	}
	var couponResponse struct {
		Data Coupon `json:"data"`
		Meta struct {
		} `json:"meta"`
	}
	err = json.Unmarshal(b, &couponResponse)
	if err != nil {
		return nil, err
	}
	return &couponResponse.Data, nil
}

func (bc *Client) DeleteCoupon(couponID int64) error {
	req := bc.getAPIRequest(http.MethodDelete, "/v3/coupons/"+strconv.FormatInt(couponID, 10), nil)
	res, err := bc.HTTPClient.Do(req)
	if err != nil {
		return err
	}
	b, err := processBody(res)
	if err != nil {
		return err
	}
	var couponResponse struct {
		Data Coupon `json:"data"`
		Meta struct {
		} `json:"meta"`
	}
	err = json.Unmarshal(b, &couponResponse)
	if err != nil {
		return err
	}
	return nil
}

func (bc *Client) GetAllCoupons(args map[string]string) ([]Coupon, error) {
	cs := []Coupon{}
	var csp []Coupon
	page := 1
	more := true
	var err error
	var retries int
	for more {
		csp, more, err = bc.GetCoupons(args, page)
		if err != nil {
			retries++
			if retries > bc.MaxRetries {
				return cs, fmt.Errorf("max retries reached")
			}
			break
		}
		cs = append(cs, csp...)
		page++
	}
	return cs, err
}

func (bc *Client) GetCoupons(args map[string]string, page int) ([]Coupon, bool, error) {
	fpart := ""
	for k, v := range args {
		fpart += "&" + k + "=" + v
	}
	url := "/v3/coupons?page=" + strconv.Itoa(page) + fpart
	req := bc.getAPIRequest(http.MethodGet, url, nil)
	res, err := bc.HTTPClient.Do(req)
	if err != nil {
		return nil, false, err
	}
	b, err := processBody(res)
	if err != nil {
		return nil, false, err
	}
	var couponResponse struct {
		Data []Coupon `json:"data"`
		Meta struct {
			Pagination Pagination `json:"pagination"`
		} `json:"meta"`
	}
	err = json.Unmarshal(b, &couponResponse)
	if err != nil {
		return nil, false, err
	}
	return couponResponse.Data, couponResponse.Meta.Pagination.CurrentPage < couponResponse.Meta.Pagination.TotalPages, nil
}
