package bigcommerce

import (
	"encoding/json"
	"fmt"
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
func (bc *Client) GetAllBrands() ([]Brand, error) {
	cs := []Brand{}
	var csp []Brand
	page := 1
	more := true
	extidmap := map[int64]int{}
	var err error
	var retries int
	for more {
		csp, more, err = bc.GetBrands(page)
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
	for i, c := range cs {
		extidmap[c.ID] = i
	}
	for i := range cs {
		cs[i].URL = cs[i].CustomURL.URL
	}
	return cs, err
}

// GetBrands returns all brands, handling pagination
// page: the page number to download
func (bc *Client) GetBrands(page int) ([]Brand, bool, error) {
	url := "/v3/catalog/brands?page=" + strconv.Itoa(page)

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
		Data []Brand `json:"data"`
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
