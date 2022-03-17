package bigcommerce

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"
)

// include_fields
var productFields = []string{"name", "sku", "custom_url", "is_visible", "price"}

// include (subresources, like variants images custom_fields bulk_pricing_rules primary_image modifiers options videos)
var productInclude []string

// Product is a BigCommerce product object
type Product struct {
	ID                      int64         `json:"id,omitempty"`
	Name                    string        `json:"name,omitempty"`
	Type                    string        `json:"type,omitempty"`
	Sku                     string        `json:"sku,omitempty"`
	Description             string        `json:"description,omitempty"`
	Weight                  float64       `json:"weight,omitempty"`
	Width                   int           `json:"width,omitempty"`
	Depth                   int           `json:"depth,omitempty"`
	Height                  int           `json:"height,omitempty"`
	Price                   float64       `json:"price,omitempty"`
	CostPrice               float64       `json:"cost_price,omitempty"`
	RetailPrice             float64       `json:"retail_price,omitempty"`
	SalePrice               float64       `json:"sale_price,omitempty"`
	MapPrice                float64       `json:"map_price,omitempty"`
	TaxClassID              int64         `json:"tax_class_id,omitempty"`
	ProductTaxCode          string        `json:"product_tax_code,omitempty"`
	CalculatedPrice         float64       `json:"calculated_price,omitempty"`
	Categories              []interface{} `json:"categories,omitempty"`
	BrandID                 int64         `json:"brand_id,omitempty"`
	OptionSetID             interface{}   `json:"option_set_id,omitempty"`
	OptionSetDisplay        string        `json:"option_set_display,omitempty"`
	InventoryLevel          int           `json:"inventory_level,omitempty"`
	InventoryWarningLevel   int           `json:"inventory_warning_level,omitempty"`
	InventoryTracking       string        `json:"inventory_tracking,omitempty"`
	ReviewsRatingSum        int           `json:"reviews_rating_sum,omitempty"`
	ReviewsCount            int           `json:"reviews_count,omitempty"`
	TotalSold               int           `json:"total_sold,omitempty"`
	FixedCostShippingPrice  float64       `json:"fixed_cost_shipping_price,omitempty"`
	IsFreeShipping          bool          `json:"is_free_shipping,omitempty"`
	IsVisible               bool          `json:"is_visible,omitempty"`
	IsFeatured              bool          `json:"is_featured,omitempty"`
	RelatedProducts         []int         `json:"related_products,omitempty"`
	Warranty                string        `json:"warranty,omitempty"`
	BinPickingNumber        string        `json:"bin_picking_number,omitempty"`
	LayoutFile              string        `json:"layout_file,omitempty"`
	Upc                     string        `json:"upc,omitempty"`
	Mpn                     string        `json:"mpn,omitempty"`
	Gtin                    string        `json:"gtin,omitempty"`
	SearchKeywords          string        `json:"search_keywords,omitempty"`
	Availability            string        `json:"availability,omitempty"`
	AvailabilityDescription string        `json:"availability_description,omitempty"`
	GiftWrappingOptionsType string        `json:"gift_wrapping_options_type,omitempty"`
	GiftWrappingOptionsList []interface{} `json:"gift_wrapping_options_list,omitempty"`
	SortOrder               int           `json:"sort_order,omitempty"`
	Condition               string        `json:"condition,omitempty"`
	IsConditionShown        bool          `json:"is_condition_shown,omitempty"`
	OrderQuantityMinimum    int           `json:"order_quantity_minimum,omitempty"`
	OrderQuantityMaximum    int           `json:"order_quantity_maximum,omitempty"`
	PageTitle               string        `json:"page_title,omitempty"`
	MetaKeywords            []interface{} `json:"meta_keywords,omitempty"`
	MetaDescription         string        `json:"meta_description,omitempty"`
	DateCreated             time.Time     `json:"date_created,omitempty"`
	DateModified            time.Time     `json:"date_modified,omitempty"`
	ViewCount               int           `json:"view_count,omitempty"`
	PreorderReleaseDate     interface{}   `json:"preorder_release_date,omitempty"`
	PreorderMessage         string        `json:"preorder_message,omitempty"`
	IsPreorderOnly          bool          `json:"is_preorder_only,omitempty"`
	IsPriceHidden           bool          `json:"is_price_hidden,omitempty"`
	PriceHiddenLabel        string        `json:"price_hidden_label,omitempty"`
	CustomURL               struct {
		URL          string `json:"url,omitempty"`
		IsCustomized bool   `json:"is_customized,omitempty"`
	} `json:"custom_url,omitempty"`
	BaseVariantID               int64  `json:"base_variant_id,omitempty"`
	OpenGraphType               string `json:"open_graph_type,omitempty"`
	OpenGraphTitle              string `json:"open_graph_title,omitempty"`
	OpenGraphDescription        string `json:"open_graph_description,omitempty"`
	OpenGraphUseMetaDescription bool   `json:"open_graph_use_meta_description,omitempty"`
	OpenGraphUseProductName     bool   `json:"open_graph_use_product_name,omitempty"`
	OpenGraphUseImage           bool   `json:"open_graph_use_image,omitempty"`
	Variants                    []struct {
		ID                        int64         `json:"id,omitempty"`
		ProductID                 int64         `json:"product_id,omitempty"`
		Sku                       string        `json:"sku,omitempty"`
		SkuID                     interface{}   `json:"sku_id,omitempty"`
		Price                     float64       `json:"price,omitempty"`
		CalculatedPrice           float64       `json:"calculated_price,omitempty"`
		SalePrice                 float64       `json:"sale_price,omitempty"`
		RetailPrice               float64       `json:"retail_price,omitempty"`
		MapPrice                  float64       `json:"map_price,omitempty"`
		Weight                    float64       `json:"weight,omitempty"`
		Width                     int           `json:"width,omitempty"`
		Height                    int           `json:"height,omitempty"`
		Depth                     int           `json:"depth,omitempty"`
		IsFreeShipping            bool          `json:"is_free_shipping,omitempty"`
		FixedCostShippingPrice    float64       `json:"fixed_cost_shipping_price,omitempty"`
		CalculatedWeight          float64       `json:"calculated_weight,omitempty"`
		PurchasingDisabled        bool          `json:"purchasing_disabled,omitempty"`
		PurchasingDisabledMessage string        `json:"purchasing_disabled_message,omitempty"`
		ImageURL                  string        `json:"image_url,omitempty"`
		CostPrice                 float64       `json:"cost_price,omitempty"`
		Upc                       string        `json:"upc,omitempty"`
		Mpn                       string        `json:"mpn,omitempty"`
		Gtin                      string        `json:"gtin,omitempty"`
		InventoryLevel            int           `json:"inventory_level,omitempty"`
		InventoryWarningLevel     int           `json:"inventory_warning_level,omitempty"`
		BinPickingNumber          string        `json:"bin_picking_number,omitempty"`
		OptionValues              []interface{} `json:"option_values,omitempty"`
	} `json:"variants,omitempty"`
	Images       []interface{} `json:"images,omitempty"`
	PrimaryImage interface{}   `json:"primary_image,omitempty"`
	Videos       []interface{} `json:"videos,omitempty"`
	CustomFields []struct {
		ID    int64  `json:"id,omitempty"`
		Name  string `json:"name,omitempty"`
		Value string `json:"value,omitempty"`
	} `json:"custom_fields,omitempty"`
	BulkPricingRules []interface{} `json:"bulk_pricing_rules,omitempty"`
	Options          []interface{} `json:"options,omitempty"`
	Modifiers        []interface{} `json:"modifiers,omitempty"`
}

// Metafield is a struct representing a BigCommerce product metafield
type Metafield struct {
	ID            int64     `json:"id,omitempty"`
	Key           string    `json:"key,omitempty"`
	Value         string    `json:"value,omitempty"`
	ResourceID    int64     `json:"resource_id,omitempty"`
	ResourceType  string    `json:"resource_type,omitempty"`
	Description   string    `json:"description,omitempty"`
	DateCreated   time.Time `json:"date_created,omitempty"`
	DateModified  time.Time `json:"date_modified,omitempty"`
	Namespace     string    `json:"namespace,omitempty"`
	PermissionSet string    `json:"permission_set,omitempty"`
}

// SetProductFields sets include_fields parameter for GetProducts, empty list will get all fields
func (bc *Client) SetProductFields(fields []string) {
	productFields = fields
}

// SetProductFields sets include_fields parameter for GetProducts, empty list will get all fields
func (bc *Client) SetProductInclude(subresources []string) {
	productInclude = subresources
}

// GetAllProducts gets all products from BigCommerce
func (bc *Client) GetAllProducts() ([]Product, error) {
	fields := productFields
	include := productInclude
	ps := []Product{}
	var psp []Product
	page := 1
	more := true
	var err error
	retries := 0
	for more {
		psp, more, err = bc.GetProducts(fields, include, page)
		log.Printf("page %d entries %d", page, len(psp))
		if err != nil {
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
	return ps, err
}

// GetProducts gets a page of products from BigCommerce
// fields is a list of fields to include in the response
// include is a list of subresources to include in the response
// page: the page number to download
func (bc *Client) GetProducts(fields, include []string, page int) ([]Product, bool, error) {
	fpart := ""
	if len(fields) != 0 {
		fpart = "&include_fields=" + strings.Join(fields, ",")
	}
	if len(include) != 0 {
		fpart = "&include=" + strings.Join(include, ",")
	}
	url := "/v3/catalog/products?page=" + strconv.Itoa(page) + fpart
	log.Printf("GET %s", url)

	req := bc.getAPIRequest(http.MethodGet, url, nil)
	res, err := bc.HTTPClient.Do(req)
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
		Status int       `json:"status"`
		Title  string    `json:"title"`
		Data   []Product `json:"data"`
		Meta   struct {
			Pagination Pagination `json:"pagination"`
		} `json:"meta"`
	}
	err = json.Unmarshal(body, &pp)
	if err != nil {
		return nil, false, err
	}
	log.Printf("%d products (%+v)", len(pp.Data), pp.Meta.Pagination)

	if pp.Status != 0 {
		return nil, false, errors.New(pp.Title)
	}
	return pp.Data, pp.Meta.Pagination.CurrentPage < pp.Meta.Pagination.TotalPages, nil
}

// GetProductByID gets a product from BigCommerce by ID
// productID: BigCommerce product ID to get
func (bc *Client) GetProductByID(productID int64) (*Product, error) {
	url := "/v3/catalog/products/" + strconv.FormatInt(productID, 10)
	req := bc.getAPIRequest(http.MethodGet, url, nil)
	res, err := bc.HTTPClient.Do(req)
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()
	body, err := processBody(res)
	if err != nil {
		return nil, err
	}

	var product Product
	err = json.Unmarshal(body, &product)
	if err != nil {
		return nil, err
	}
	return &product, nil
}

// GetProductMetafields gets metafields values for a product
// productID: BigCommerce product ID to get metafields for
func (bc *Client) GetProductMetafields(productID int64) (map[string]Metafield, error) {
	url := "/v3/catalog/products/" + strconv.FormatInt(productID, 10) + "/metafields"
	req := bc.getAPIRequest(http.MethodGet, url, nil)
	res, err := bc.HTTPClient.Do(req)
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()
	body, err := processBody(res)
	if err != nil {
		return nil, err
	}

	var metafieldsResponse struct {
		Metafields []Metafield `json:"data,omitempty"`
	}
	err = json.Unmarshal(body, &metafieldsResponse)
	if err != nil {
		return nil, err
	}
	ret := map[string]Metafield{}
	for _, mf := range metafieldsResponse.Metafields {
		ret[mf.Key] = mf
	}
	return ret, nil
}
