package bigcommerce

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
)

// Brand is BigCommerce brand object
type Brand struct {
	ID              int64    `json:"id"`
	Name            string   `json:"name"`
	PageTitle       string   `json:"page_title"`
	MetaKeywords    []string `json:"meta_keywords"`
	MetaDescription string   `json:"meta_description"`
	ImageURL        string   `json:"image_url"`
	SearchKeywords  string   `json:"search_keywords"`
	CustomURL       struct {
		URL          string `json:"url"`
		IsCustomized bool   `json:"is_customized"`
	} `json:"custom_url"`
	URL string `json:"-"`
}

// GetAllBrands returns all brands, handling pagination
// context: the BigCommerce context (e.g. stores/23412341234) where 23412341234 is the store hash
// xAuthToken: the BigCommerce Store's X-Auth-Token coming from store credentials (see AuthContext)
func (bc *BigCommerce) GetAllBrands(context, xAuthToken string) ([]Brand, error) {
	cs := []Brand{}
	var csp []Brand
	page := 1
	more := true
	extidmap := map[int64]int{}
	var err error
	var retries int
	for more {
		csp, more, err = bc.GetBrands(context, xAuthToken, page)
		if err != nil {
			log.Println(err)
			retries++
			if retries > bc.MaxRetries {
				return cs, fmt.Errorf("max retries reached")
			}
			break
		}
		cs = append(cs, csp...)
		page++
	}
	for i, c := range cs {
		extidmap[c.ID] = i
	}
	for i := range cs {
		cs[i].URL = cs[i].CustomURL.URL
	}
	return cs, err
}

// GetBrands returns all brands, handling pagination
// context: the BigCommerce context (e.g. stores/23412341234) where 23412341234 is the store hash
// xAuthToken: the BigCommerce Store's X-Auth-Token coming from store credentials (see AuthContext)
// page: the page number to download
func (bc *BigCommerce) GetBrands(context, xAuthToken string, page int) ([]Brand, bool, error) {
	url := context + "/v3/catalog/brands?page=" + strconv.Itoa(page)

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
		Data []Brand `json:"data"`
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
