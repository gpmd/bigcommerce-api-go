package bigcommerce

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
)

// Address is for Customer Address endpoint
type Address struct {
	ID              int64       `json:"id"`
	CustomerID      int64       `json:"customer_id"`
	Address1        string      `json:"address1"`
	Address2        string      `json:"address2"`
	AddressType     string      `json:"address_type"`
	City            string      `json:"city"`
	Company         string      `json:"company"`
	Country         string      `json:"country"`
	CountryCode     string      `json:"country_code"`
	FirstName       string      `json:"first_name"`
	LastName        string      `json:"last_name"`
	Phone           string      `json:"phone"`
	PostalCode      string      `json:"postal_code"`
	StateOrProvince string      `json:"state_or_province"`
	FormFields      []FormField `json:"form_fields"`
}

// GetAddresses returns all addresses for a curstomer, handling pagination
// customerID is bigcommerce customer id
func (bc *Client) GetAddresses(customerID int64) ([]Address, error) {
	cs := []Address{}
	var csp []Address
	page := 1
	more := true
	var err error
	var retries int
	for more {
		csp, more, err = bc.GetAddressPage(customerID, page)
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

// GetAddressPage returns all addresses for a curstomer, handling pagination
// customerID is bigcommerce customer id
// page: the page number to download
func (bc *Client) GetAddressPage(customerID int64, page int) ([]Address, bool, error) {
	url := "/v3/customers/addresses?customer_id:in=" + strconv.FormatInt(customerID, 10) + "&page=" + strconv.Itoa(page)

	req := bc.getAPIRequest(http.MethodGet, url, nil)
	res, err := bc.HTTPClient.Do(req)
	if err != nil {
		return nil, false, err
	}

	defer res.Body.Close()
	body, err := processBody(res)
	if err != nil {
		return nil, false, err
	}

	var pp struct {
		Data []Address `json:"data"`
		Meta struct {
			Pagination Pagination `json:"pagination"`
		} `json:"meta"`
	}
	err = json.Unmarshal(body, &pp)
	if err != nil {
		return nil, false, err
	}
	return pp.Data, pp.Meta.Pagination.CurrentPage < pp.Meta.Pagination.TotalPages, nil
}

// CreateAddress creates a new address for a customer from given data, ignoring ID (duplicating address)
func (bc *Client) CreateAddress(customerID int64, address *Address) (*Address, error) {
	url := "/v3/customers/addresses"
	// extra safety feature so we don't edit other customers' address
	address.CustomerID = customerID
	addressJSON, _ := json.Marshal([]Address{*address})
	//	log.Printf("addressJSON: %s", string(addressJSON))
	req := bc.getAPIRequest(http.MethodPost, url, bytes.NewReader(addressJSON))
	res, err := bc.HTTPClient.Do(req)
	if err != nil {
		return nil, err
	}
	var addr struct {
		Addresses []Address `json:"data"`
	}
	defer res.Body.Close()
	body, err := processBody(res)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(body, &addr)
	if err != nil {
		return nil, fmt.Errorf("error parsing body: %s %s", err, string(body))
	}
	if len(addr.Addresses) == 0 {
		return nil, fmt.Errorf("no address returned, %s", string(body))
	}
	return &addr.Addresses[0], nil
}

// UpdateAddress updates an existing address, address ID is required
func (bc *Client) UpdateAddress(customerID int64, address *Address) (*Address, error) {
	url := "/v3/customers/addresses"
	// extra safety feature so we don't edit other customers' address
	address.CustomerID = customerID
	if address.ID == 0 {
		return nil, fmt.Errorf("address ID is required")
	}
	addressJSON, _ := json.Marshal([]Address{*address})
	//	log.Printf("addressJSON: %s", string(addressJSON))
	req := bc.getAPIRequest(http.MethodPut, url, bytes.NewReader(addressJSON))
	res, err := bc.HTTPClient.Do(req)
	if err != nil {
		return nil, err
	}
	var addr struct {
		Addresses []Address `json:"data"`
	}
	defer res.Body.Close()
	body, err := processBody(res)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(body, &addr)
	if err != nil {
		return nil, fmt.Errorf("error parsing body: %s %s", err, string(body))
	}
	if len(addr.Addresses) == 0 {
		return nil, fmt.Errorf("no address returned, %s", string(body))
	}
	return &addr.Addresses[0], nil
}

// DeleteAddress deletes an existing address, address ID is required
func (bc *Client) DeleteAddress(customerID, addressID int64) error {
	url := "/v3/customers/addresses?id:in=" + strconv.FormatInt(addressID, 10)
	req := bc.getAPIRequest(http.MethodDelete, url, nil)
	res, err := bc.HTTPClient.Do(req)
	if err != nil {
		return err
	}
	defer res.Body.Close()
	if res.StatusCode == http.StatusNoContent {
		return nil
	}
	body, err := processBody(res)
	if err != nil {
		log.Printf("error processing body: %s", err)
		return err
	}
	var addr Address
	err = json.Unmarshal(body, &addr)
	if err != nil {
		log.Printf("error parsing body: %s %s", err, string(body))
		return err
	}
	return nil
}
