package bigcommerce

import (
	"bytes"
	"encoding/json"
	"net/http"
)

// Customer is a struct for the BigCommerce Customer API
type Customer struct {
	ID               int64             `json:"id"`
	Company          string            `json:"company"`
	Firstname        string            `json:"first_name"`
	Lastname         string            `json:"last_name"`
	Email            string            `json:"email"`
	Phone            string            `json:"phone"`
	FormFields       interface{}       `json:"form_fields"`
	DateCreated      string            `json:"date_created"`
	DateModified     string            `json:"date_modified"`
	StoreCredit      string            `json:"store_credit"`
	RegistrationIP   string            `json:"registration_ip_address"`
	CustomerGroup    int64             `json:"customer_group_id"`
	Notes            string            `json:"notes"`
	TaxExempt        string            `json:"tax_exempt_category"`
	ResetPassword    bool              `json:"reset_pass_on_login"`
	AcceptsMarketing bool              `json:"accepts_marketing"`
	Addresses        map[string]string `json:"addresses"`
}

func (bc *BigCommerce) ValidateCredentials(email, password string) (int64, error) {
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
