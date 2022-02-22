package bigcommerce

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"strings"
)

var productFields = []string{"name", "sku", "custom_url", "is_visible", "price"}

// Product is a BigCommerce product object
type Product struct {
	ID        int64   `json:"id"`
	Name      string  `json:"name"`
	SKU       string  `json:"sku"`
	Visible   bool    `json:"is_visible"`
	Thumbnail string  `json:"thumbnail"`
	URL       string  `json:"-"`
	Price     float64 `json:"price"`
	CustomURL struct {
		URL          string `json:"url"`
		IsCustomised bool   `json:"is_customized"`
	} `json:"custom_url"`
}

// SetProductFields sets include_fields parameter for GetProducts, empty list will get all fields
func (bc *BigCommerce) SetProductFields(f []string) {
	productFields = f
}

// GetAllProducts gets all products from BigCommerce
// context: the BigCommerce context (e.g. stores/23412341234) where 23412341234 is the store hash
// xAuthToken: the BigCommerce Store's X-Auth-Token coming from store credentials (see AuthContext)
func (bc *BigCommerce) GetAllProducts(context, xAuthToken string) ([]Product, error) {
	ps := []Product{}
	var psp []Product
	page := 1
	more := true
	var err error
	retries := 0
	for more {
		psp, more, err = bc.GetProducts(context, xAuthToken, page)
		if err != nil {
			log.Println(err)
			retries++
			if retries > bc.MaxRetries {
				log.Println("Max retries reached")
				return ps, err
			}
			break
		}
		ps = append(ps, psp...)
		page++
	}
	return ps, nil
}

// GetProducts gets a page of products from BigCommerce
// context: the BigCommerce context (e.g. stores/23412341234) where 23412341234 is the store hash
// xAuthToken: the BigCommerce Store's X-Auth-Token coming from store credentials (see AuthContext)
// page: the page number to download
func (bc *BigCommerce) GetProducts(context, xAuthToken string, page int) ([]Product, bool, error) {
	fpart := ""
	if len(productFields) != 0 {
		fpart = "&include_fields=" + strings.Join(productFields, ",")
	}
	url := context + "/v3/catalog/products?page=" + strconv.Itoa(page) + fpart

	req := bc.getAPIRequest(http.MethodGet, url, xAuthToken, nil)
	res, err := bc.DefaultClient.Do(req)
	if err != nil {
		return nil, false, err
	}
	defer res.Body.Close()
	if res.StatusCode == http.StatusNoContent {
		return nil, false, ErrNoContent
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, false, err
	}
	var pp struct {
		Data []Product `json:"data"`
		Meta struct {
			Pagination struct {
				Total       int64       `json:"total"`
				Count       int64       `json:"count"`
				PerPage     int64       `json:"per_page"`
				CurrentPage int64       `json:"current_page"`
				TotalPages  int64       `json:"total_pages"`
				Links       interface{} `json:"links"`
				TooMany     bool        `json:"too_many"`
			} `json:"pagination"`
		} `json:"meta"`
	}
	err = json.Unmarshal(body, &pp)
	if err != nil {
		return nil, false, err
	}
	return pp.Data, pp.Meta.Pagination.CurrentPage < pp.Meta.Pagination.TotalPages, nil
}
