package bigcommerce

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"time"
)

// Product is a BigCommerce product object
type Product struct {
	ID        int64   `json:"-" db:"id"`
	ClientID  int64   `json:"-" db:"client_id"`
	ExtID     int64   `json:"id" db:"ext_id"`
	Name      string  `json:"name" db:"name"`
	SKU       string  `json:"sku" db:"sku"`
	Visible   bool    `json:"is_visible" db:"visible"`
	Thumbnail string  `json:"thumbnail" db:"thumbnail_url"`
	URL       string  `json:"-" db:"url"`
	Price     float64 `json:"price" db:"price"`
	CustomURL struct {
		URL          string `json:"url" db:"-"`
		IsCustomised bool   `json:"is_customized" db:"-"`
	} `json:"custom_url" db:"-"`
}

// GetAllProducts gets all products from BigCommerce
func (bc *BigCommerce) GetAllProducts(context, client, token string) []Product {
	ps := []Product{}
	var psp []Product
	page := 1
	more := true
	var err error
	retries := 0
	for more {
		psp, more, err = bc.GetProducts(context, client, token, page)
		if err != nil {
			log.Println(err)
			retries++
			if retries > bc.MaxRetries {
				log.Println("Max retries reached")
				return ps
			}
			break
		}
		log.Println("More prods:", more, " count:", len(psp))
		ps = append(ps, psp...)
		page++
	}
	return ps
}

// GetProducts gets a page of products from BigCommerce
func (bc *BigCommerce) GetProducts(context, client, token string, page int) ([]Product, bool, error) {
	url := context + "/v3/catalog/products?include_fields=name,sku,custom_url,is_visible,price&page=" + strconv.Itoa(page)

	req := bc.getAPIRequest(http.MethodGet, url, client, token)
	var c = &http.Client{
		Timeout: time.Second * 10,
	}
	res, err := c.Do(req)
	if err != nil {
		return nil, false, err
	}
	defer res.Body.Close()
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
