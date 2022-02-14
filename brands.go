package bigcommerce

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"time"
)

// Brand is BigCommerce brand object
type Brand struct {
	ExtID           int64         `json:"id" db:"ext_id"`
	Name            string        `json:"name"`
	PageTitle       string        `json:"page_title"`
	MetaKeywords    []string      `json:"meta_keywords"`
	MetaDescription string        `json:"meta_description"`
	ID              int           `json:"-" db:"id"`
	ImageURL        string        `json:"image_url"`
	ClientID        sql.NullInt64 `json:"-" db:"client_id"`
	SearchKeywords  string        `json:"search_keywords"`
	CustomURL       struct {
		URL          string `json:"url"`
		IsCustomized bool   `json:"is_customized"`
	} `json:"custom_url"`
	URL string `json:"-" db:"url"`
}

func (bc *BigCommerce) GetAllBrands(context, client, token string) ([]Brand, error) {
	cs := []Brand{}
	var csp []Brand
	page := 1
	more := true
	extidmap := map[int64]int{}
	var err error
	var retries int
	for more {
		csp, more, err = bc.GetBrands(context, client, token, page)
		if err != nil {
			log.Println(err)
			retries++
			if retries > bc.MaxRetries {
				return cs, fmt.Errorf("max retries reached")
			}
			break
		}
		log.Println("More brands:", more, " count:", len(csp))
		cs = append(cs, csp...)
		page++
	}
	for i, c := range cs {
		extidmap[c.ExtID] = i
	}
	for i := range cs {
		cs[i].URL = cs[i].CustomURL.URL
	}
	return cs, err
}

func (bc *BigCommerce) GetBrands(context, client, token string, page int) ([]Brand, bool, error) {
	url := context + "/v3/catalog/brands?page=" + strconv.Itoa(page)

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
