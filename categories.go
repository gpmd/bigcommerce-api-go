package bigcommerce

import (
	"encoding/json"
	"fmt"
	"net/http"
	"sort"
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
	cats := map[int64]Category{}
	ids := []int64{}
	for _, c := range cs {
		cats[c.ID] = c
		//		log.Printf("%d: %s", c.ID, c.Name)
		ids = append(ids, c.ID)
	}
	// sort ids
	sort.Slice(ids, func(i, j int) bool {
		return ids[i] < ids[j]
	})
	cs = []Category{}
	for _, i := range ids {
		c := cats[i]
		c.URL = c.CustomURL.URL
		// get A > B > C fancy name
		c.FullName = bc.getFullCategoryName(cats, i)
		cs = append(cs, c)
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

func (bc *Client) getFullCategoryName(cats map[int64]Category, i int64) string {
	c := cats[i]
	if c.ParentID == 0 {
		return c.Name
	}
	if c.FullName != "" {
		return c.FullName
	}
	pn := bc.getFullCategoryName(cats, cats[i].ParentID)
	c.FullName = pn + " > " + c.Name
	cats[i] = c
	return c.FullName
}
