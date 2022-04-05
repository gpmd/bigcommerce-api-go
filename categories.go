package bigcommerce

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

// Category is a BC category object
type Category struct {
	ID        int64  `json:"id"`
	Name      string `json:"name"`
	ParentID  int64  `json:"parent_id"`
	Visible   bool   `json:"is_visible"`
	FullName  string `json:"-"`
	CustomURL struct {
		URL        string `json:"url"`
		Customized bool   `json:"is_customized"`
	} `json:"custom_url"`
	URL string `json:"-"`
}

// GetAllCategories returns a list of categories, handling pagination
// args is a map of arguments to pass to the API
func (bc *Client) GetAllCategories(args map[string]string) ([]Category, error) {
	cs := []Category{}
	var csp []Category
	page := 1
	more := true
	var err error
	retries := 0
	for more {
		csp, more, err = bc.GetCategories(args, page)
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
	extidmap := map[int64]int{}
	for i, c := range cs {
		extidmap[c.ID] = i
	}
	for i := range cs {
		cs[i].URL = cs[i].CustomURL.URL
		// get A > B > C fancy name
		cs[i].FullName = bc.getFullCategoryName(cs, i, extidmap)
	}
	return cs, err
}

// GetCategories returns a list of categories, handling pagination
// args is a map of arguments to pass to the API
// page: the page number to download
func (bc *Client) GetCategories(args map[string]string, page int) ([]Category, bool, error) {
	fpart := ""
	for k, v := range args {
		fpart += "&" + k + "=" + v
	}
	url := "/v3/catalog/categories?page=" + strconv.Itoa(page) + fpart

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
		Data []Category `json:"data"`
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

func (bc *Client) getFullCategoryName(cs []Category, i int, extidmap map[int64]int) string {
	if cs[i].ParentID == 0 {
		return cs[i].Name
	}
	return bc.getFullCategoryName(cs, extidmap[cs[i].ParentID], extidmap) + " > " + cs[i].Name
}
