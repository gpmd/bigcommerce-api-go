package bigcommerce

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"strings"
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

type SaveAccountPayload struct {
	ID                int64     `json:"id"`
	Company           string    `json:"company"`
	FirstName         string    `json:"first_name"`
	LastName          string    `json:"last_name"`
	Email             string    `json:"email"`
	Phone             string    `json:"phone"`
	Notes             string    `json:"notes"`
	TaxExemptCategory string    `json:"tax_exempt_category"`
	CustomerGroupID   int64     `json:"customer_group_id"`
	Addresses         []Address `json:"addresses"`
	Authentication    struct {
		ForcePasswordReset bool   `json:"force_password_reset"`
		NewPassword        string `json:"new_password"`
	} `json:"authentication"`
	AcceptsProductReviewAbandonedCartEmails bool `json:"accepts_product_review_abandoned_cart_emails"`
	StoreCreditAmounts                      []struct {
		Amount float64 `json:"amount"`
	} `json:"store_credit_amounts"`
	OriginChannelID int   `json:"origin_channel_id"`
	ChannelIDs      []int `json:"channel_ids"`
	FormFields      []struct {
		Name  string `json:"name"`
		Value string `json:"value"`
	} `json:"form_fields"`
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
	ChannelIDs                              []int          `json:"channel_ids"`
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
	if payload.OriginChannelID == 0 {
		payload.OriginChannelID = bc.ChannelID
	}
	if payload.ChannelIDs == nil {
		payload.ChannelIDs = []int{bc.ChannelID}
	}
	var b []byte
	b, _ = json.Marshal([]CreateAccountPayload{*payload})
	req := bc.getAPIRequest(http.MethodPost, "/v3/customers", bytes.NewBuffer(b))
	res, err := bc.HTTPClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	body, err := processBody(res)
	if err != nil {
		if res.StatusCode == http.StatusUnprocessableEntity {
			var errResp ErrorResult
			err = json.Unmarshal(body, &errResp)
			if err != nil {
				log.Printf("Error: %s\nResult: %s", err, string(body))
				return nil, err
			}
			if len(errResp.Errors) > 0 {
				errors := []string{}
				for _, e := range errResp.Errors {
					errors = append(errors, e)
				}
				return nil, fmt.Errorf("%s", strings.Join(errors, ", "))
			}
			return nil, errors.New("unknown error")
		}
		log.Printf("Error: %s\nResult: %s", err, string(body))
		return nil, err
	}
	var ret struct {
		Customers []Customer `json:"data"`
	}
	err = json.Unmarshal(body, &ret)
	if err != nil {
		return nil, err
	}
	return &ret.Customers[0], nil
}

// CreateAccount creates a new customer account in BigCommerce and returns the customer or error
func (bc *Client) SaveAccount(payload *SaveAccountPayload) (*Customer, error) {
	if payload.OriginChannelID == 0 {
		payload.OriginChannelID = bc.ChannelID
	}
	if payload.ChannelIDs == nil {
		payload.ChannelIDs = []int{bc.ChannelID}
	}
	var b []byte
	b, _ = json.Marshal([]SaveAccountPayload{*payload})
	req := bc.getAPIRequest(http.MethodPut, "/v3/customers", bytes.NewBuffer(b))
	res, err := bc.HTTPClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	body, err := processBody(res)
	if err != nil {
		if res.StatusCode == http.StatusUnprocessableEntity {
			var errResp ErrorResult
			err = json.Unmarshal(body, &errResp)
			if err != nil {
				log.Printf("Error: %s\nResult: %s", err, string(body))
				return nil, err
			}
			if len(errResp.Errors) > 0 {
				errors := []string{}
				for _, e := range errResp.Errors {
					errors = append(errors, e)
				}
				return nil, fmt.Errorf("%s", strings.Join(errors, ", "))
			}
			return nil, errors.New("unknown error")
		}
		log.Printf("Error: %s\nResult: %s", err, string(body))
		return nil, err
	}
	var ret struct {
		Customers []Customer `json:"data"`
	}
	err = json.Unmarshal(body, &ret)
	if err != nil {
		return nil, err
	}
	return &ret.Customers[0], nil
}

// CustomerSetFormFields sets the form fields for a customer
func (bc *Client) CustomerSetFormFields(customerID int64, formFields []FormField) error {
	if customerID == 0 {
		return errors.New("customerID cannot be 0")
	}
	for i := range formFields {
		formFields[i].CustomerID = customerID
	}
	var b []byte
	b, _ = json.Marshal(formFields)
	log.Printf("Fields: %s", string(b))
	req := bc.getAPIRequest(http.MethodPut, "/v3/customers/form-field-values", bytes.NewBuffer(b))
	res, err := bc.HTTPClient.Do(req)
	if err != nil {
		return err
	}
	defer res.Body.Close()
	body, err := processBody(res)
	if err != nil {
		if res.StatusCode == http.StatusUnprocessableEntity {
			var errResp ErrorResult
			err = json.Unmarshal(body, &errResp)
			if err != nil {
				log.Printf("Error: %s\nResult: %s", err, string(body))
				return err
			}
			if len(errResp.Errors) > 0 {
				errors := []string{}
				for _, e := range errResp.Errors {
					errors = append(errors, e)
				}
				return fmt.Errorf("%s", strings.Join(errors, ", "))
			}
			return errors.New("unknown error")
		}
		return err
	}
	return nil
}

func (bc *Client) CustomerGetFormFields(customerID int64) ([]FormField, error) {
	req := bc.getAPIRequest(http.MethodGet, fmt.Sprintf("/v3/customers/form-field-values?customer_id=%d", customerID), nil)
	res, err := bc.HTTPClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	body, err := processBody(res)
	if err != nil {
		return nil, err
	}
	var ret struct {
		Data []FormField `json:"data"`
	}
	err = json.Unmarshal(body, &ret)
	if err != nil {
		return nil, err
	}
	log.Printf("Form fields: %s", string(body))
	return ret.Data, nil
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
	var ret struct {
		Data []Customer `json:"data"`
	}
	err = json.Unmarshal(body, &ret)
	if err != nil {
		return nil, err
	}
	if len(ret.Data) == 0 {
		return nil, ErrNotFound
	}
	return &ret.Data[0], nil
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
	var ret struct {
		Data []Customer `json:"data"`
	}
	err = json.Unmarshal(body, &ret)
	if err != nil {
		return nil, err
	}
	if len(ret.Data) == 0 {
		return nil, ErrNotFound
	}
	return &ret.Data[0], nil // return the first customer
}
