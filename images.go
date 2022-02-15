package bigcommerce

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"time"
)

// Image is entry for BC product images
type Image struct {
	ID           int64  `json:"id"`
	ProductID    int64  `json:"product_id"`
	IsThumbnail  bool   `json:"is_thumbnail"`
	SortOrder    int64  `json:"sort_order"`
	Description  string `json:"description"`
	ImageFile    string `json:"image_file"`
	URLZoom      string `json:"url_zoom"`
	URLStandard  string `json:"url_standard"`
	URLThumbnail string `json:"url_thumbnail"`
	URLTiny      string `json:"url_tiny"`
	DateModified string `json:"date_modified"`
}

// GetMainThumbnailURL returns the main thumbnail URL for a product
// this is due to the fact that the Product API does not return the main thumbnail URL
func (bc *BigCommerce) GetMainThumbnailURL(context, xAuthClient, xAuthToken string, bigcommerceProductID int64) string {
	url := context + "/v3/catalog/products/" + strconv.FormatInt(bigcommerceProductID, 10) + "/images"

	req := bc.getAPIRequest(http.MethodGet, url, xAuthClient, xAuthToken)
	var c = &http.Client{
		Timeout: time.Second * 10,
	}
	res, err := c.Do(req)
	if err != nil {
		return ""
	}

	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Println(err)
		return ""
	}

	var pp struct {
		Data []Image `json:"data"`
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
		log.Println(err)
		return ""
	}
	for _, p := range pp.Data {
		if p.IsThumbnail {
			return p.URLThumbnail
		}
	}
	return ""
}
