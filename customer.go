package bigcommerce

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

// Customer is a struct for the BigCommerce Customer API
type Customer struct {
	ID               int64       `json:"id"`
	Company          string      `json:"company"`
	Firstname        string      `json:"first_name"`
	Lastname         string      `json:"last_name"`
	Email            string      `json:"email"`
	Phone            string      `json:"phone"`
	FormFields       interface{} `json:"form_fields"`
	DateCreated      string      `json:"date_created"`
	DateModified     string      `json:"date_modified"`
	StoreCredit      string      `json:"store_credit"`
	RegistrationIP   string      `json:"registration_ip_address"`
	CustomerGroup    int64       `json:"customer_group_id"`
	Notes            string      `json:"notes"`
	TaxExempt        string      `json:"tax_exempt_category"`
	ResetPassword    bool        `json:"reset_pass_on_login"`
	AcceptsMarketing bool        `json:"accepts_marketing"`
	Addresses        []Address   `json:"addresses"`
}

type CreateAccountPayload struct {
	Company                                 string         `json:"company"`
	FirstName                               string         `json:"first_name"`
	LastName                                string         `json:"last_name"`
	Email                                   string         `json:"email"`
	Phone                                   string         `json:"phone"`
	Notes                                   string         `json:"notes"`
	TaxExemptCategory                       string         `json:"tax_exempt_category"`
	CustomerGroupID                         int64          `json:"customer_group_id"`
	Addresses                               []Address      `json:"addresses"`
	Authentication                          Authentication `json:"authentication"`
	AcceptsProductReviewAbandonedCartEmails bool           `json:"accepts_product_review_abandoned_cart_emails"`
	StoreCreditAmounts                      []StoreCredit  `json:"store_credit_amounts"`
	OriginChannelID                         int            `json:"origin_channel_id"`
	ChannelIds                              []int          `json:"channel_ids"`
}

// StoreCredit is for CreateAccountPayload's store_credit_ammounts field
type StoreCredit struct {
	Amount float64 `json:"amount"`
}

// AccountAuthentication is for CreateAccountPayload's authentication field
type Authentication struct {
	ForcePasswordReset bool   `json:"force_password_reset"`
	Password           string `json:"new_password"`
}

// Address is for CreateAccountPayload's addresses field
type Address struct {
	Address1        string `json:"address1"`
	Address2        string `json:"address2"`
	AddressType     string `json:"address_type"`
	City            string `json:"city"`
	Company         string `json:"company"`
	CountryCode     string `json:"country_code"`
	FirstName       string `json:"first_name"`
	LastName        string `json:"last_name"`
	Phone           string `json:"phone"`
	PostalCode      string `json:"postal_code"`
	StateOrProvince string `json:"state_or_province"`
}

// FormField is a struct for the BigCommerce Customer API Form Fiel values
type FormField struct {
	CustomerID int64  `json:"customer_id"`
	Name       string `json:"name"`
	Value      string `json:"value"`
}

// ValidateCredentials returns customer ID or error (i.e. ErrNotfound) if the provided credentials are valid in BigCommerce
func (bc *Client) ValidateCredentials(email, password string) (int64, error) {
	var credReq struct {
		Email     string `json:"email"`
		Password  string `json:"password"`
		ChannelID int    `json:"channel_id"`
	}
	credReq.Email = email
	credReq.Password = password
	credReq.ChannelID = bc.ChannelID
	var b []byte
	b, _ = json.Marshal(credReq)
	req := bc.getAPIRequest(http.MethodPost, "/v3/customers/validate-credentials", bytes.NewBuffer(b))
	res, err := bc.HTTPClient.Do(req)
	if err != nil {
		return 0, err
	}
	defer res.Body.Close()
	body, err := processBody(res)
	if err != nil {
		return 0, err
	}
	var credResptype struct {
		IsValid    bool  `json:"is_valid"`
		CustomerID int64 `json:"customer_id"`
	}
	err = json.Unmarshal(body, &credResptype)
	if err != nil {
		return 0, err
	}
	if !credResptype.IsValid {
		return 0, ErrNotFound
	}
	return credResptype.CustomerID, nil
}

// CreateAccount creates a new customer account in BigCommerce and returns the customer or error
func (bc *Client) CreateAccount(payload *CreateAccountPayload) (*Customer, error) {
	var b []byte
	b, _ = json.Marshal(payload)
	req := bc.getAPIRequest(http.MethodPost, "/v3/customers", bytes.NewBuffer(b))
	res, err := bc.HTTPClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	body, err := processBody(res)
	if err != nil {
		return nil, err
	}
	var customer Customer
	err = json.Unmarshal(body, &customer)
	if err != nil {
		return nil, err
	}
	return &customer, nil
}

// CustomerSetFormFields sets the form fields for a customer
func (bc *Client) CustomerSetFormFields(customerID int64, formFields []FormField) error {
	for _, formField := range formFields {
		formField.CustomerID = customerID
	}
	var b []byte
	b, _ = json.Marshal(formFields)
	req := bc.getAPIRequest(http.MethodPut, "/v3/customers/form-field-values", bytes.NewBuffer(b))
	res, err := bc.HTTPClient.Do(req)
	if err != nil {
		return err
	}
	defer res.Body.Close()
	_, err = processBody(res)
	if err != nil {
		return err
	}
	return nil
}

func (bc *Client) CustomerGetFormFields(customerID int64) ([]FormField, error) {
	req := bc.getAPIRequest(http.MethodGet, fmt.Sprintf("/v3/customers/form-fields?customer_id=%d", customerID), nil)
	res, err := bc.HTTPClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	body, err := processBody(res)
	if err != nil {
		return nil, err
	}
	var formFields []FormField
	err = json.Unmarshal(body, &formFields)
	if err != nil {
		return nil, err
	}
	return formFields, nil
}

func (bc *Client) GetCustomerByID(customerID int64) (*Customer, error) {
	req := bc.getAPIRequest(http.MethodGet, fmt.Sprintf("/v3/customers?id:in=%d", customerID), nil)
	res, err := bc.HTTPClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	body, err := processBody(res)
	if err != nil {
		return nil, err
	}
	var customers []Customer
	err = json.Unmarshal(body, &customers)
	if err != nil {
		return nil, err
	}
	if len(customers) == 0 {
		return nil, ErrNotFound
	}
	return &customers[0], nil
}

func (bc *Client) GetCustomerByEmail(email string) (*Customer, error) {
	req := bc.getAPIRequest(http.MethodGet, fmt.Sprintf("/v3/customers?email:in=%s", email), nil)
	res, err := bc.HTTPClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	body, err := processBody(res)
	if err != nil {
		return nil, err
	}
	var customers []Customer
	err = json.Unmarshal(body, &customers)
	if err != nil {
		return nil, err
	}
	if len(customers) == 0 {
		return nil, ErrNotFound
	}
	return &customers[0], nil // BigCommerce can have multiple customers with same email, we are returning the first one
}
