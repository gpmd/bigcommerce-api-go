package bigcommerce

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"time"
)

// Category is a BC category object
type Category struct {
	ExtID     int64       `json:"id" db:"ext_id"`
	Name      string      `json:"name"`
	ParentID  int64       `json:"parent_id"`
	Visible   bool        `json:"is_visible"`
	ID        interface{} `json:"-" db:"id"`
	ClientID  interface{} `json:"-" db:"client_id"`
	FullName  string      `json:"-" db:"full_name"`
	CustomURL struct {
		URL        string `json:"url" db:"url"`
		Customized bool   `json:"is_customized"`
	} `json:"custom_url"`
	URL string `json:"-" db:"url"`
}

func (bc *BigCommerce) GetAllCategories(context, client, token string) ([]Category, error) {
	cs := []Category{}
	var csp []Category
	page := 1
	more := true
	var err error
	retries := 0
	for more {
		csp, more, err = bc.GetCategories(context, client, token, page)
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
	extidmap := map[int64]int{}
	for i, c := range cs {
		extidmap[c.ExtID] = i
	}
	for i := range cs {
		cs[i].URL = cs[i].CustomURL.URL
		// get A > B > C fancy name
		cs[i].FullName = bc.getFullCategoryName(cs, i, extidmap)
	}
	return cs, err
}

func (bc *BigCommerce) GetCategories(context, client, token string, page int) ([]Category, bool, error) {
	url := context + "/v3/catalog/categories?include_fields=name,parent_id,is_visible,custom_url&page=" + strconv.Itoa(page)

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
		Data []Category `json:"data"`
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

func (bc *BigCommerce) getFullCategoryName(cs []Category, i int, extidmap map[int64]int) string {
	if cs[i].ParentID == 0 {
		return cs[i].Name
	}
	return bc.getFullCategoryName(cs, extidmap[cs[i].ParentID], extidmap) + " > " + cs[i].Name
}
