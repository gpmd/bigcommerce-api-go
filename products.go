package bigcommerce

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"time"
)

// Product is a BigCommerce product object
type Product struct {
	ID                      int64         `json:"id"`
	Name                    string        `json:"name"`
	Type                    string        `json:"type"`
	Sku                     string        `json:"sku"`
	Description             string        `json:"description"`
	Weight                  float64       `json:"weight"`
	Width                   int           `json:"width"`
	Depth                   int           `json:"depth"`
	Height                  int           `json:"height"`
	Price                   float64       `json:"price"`
	CostPrice               float64       `json:"cost_price"`
	RetailPrice             float64       `json:"retail_price"`
	SalePrice               float64       `json:"sale_price"`
	MapPrice                float64       `json:"map_price"`
	TaxClassID              int64         `json:"tax_class_id"`
	ProductTaxCode          string        `json:"product_tax_code"`
	CalculatedPrice         float64       `json:"calculated_price"`
	Categories              []interface{} `json:"categories"`
	BrandID                 int64         `json:"brand_id"`
	OptionSetID             interface{}   `json:"option_set_id"`
	OptionSetDisplay        string        `json:"option_set_display"`
	InventoryLevel          int           `json:"inventory_level"`
	InventoryWarningLevel   int           `json:"inventory_warning_level"`
	InventoryTracking       string        `json:"inventory_tracking"`
	ReviewsRatingSum        int           `json:"reviews_rating_sum"`
	ReviewsCount            int           `json:"reviews_count"`
	TotalSold               int           `json:"total_sold"`
	FixedCostShippingPrice  float64       `json:"fixed_cost_shipping_price"`
	IsFreeShipping          bool          `json:"is_free_shipping"`
	IsVisible               bool          `json:"is_visible"`
	IsFeatured              bool          `json:"is_featured"`
	RelatedProducts         []int         `json:"related_products"`
	Warranty                string        `json:"warranty"`
	BinPickingNumber        string        `json:"bin_picking_number"`
	LayoutFile              string        `json:"layout_file"`
	Upc                     string        `json:"upc"`
	Mpn                     string        `json:"mpn"`
	Gtin                    string        `json:"gtin"`
	SearchKeywords          string        `json:"search_keywords"`
	Availability            string        `json:"availability"`
	AvailabilityDescription string        `json:"availability_description"`
	GiftWrappingOptionsType string        `json:"gift_wrapping_options_type"`
	GiftWrappingOptionsList []interface{} `json:"gift_wrapping_options_list"`
	SortOrder               int           `json:"sort_order"`
	Condition               string        `json:"condition"`
	IsConditionShown        bool          `json:"is_condition_shown"`
	OrderQuantityMinimum    int           `json:"order_quantity_minimum"`
	OrderQuantityMaximum    int           `json:"order_quantity_maximum"`
	PageTitle               string        `json:"page_title"`
	MetaKeywords            []interface{} `json:"meta_keywords"`
	MetaDescription         string        `json:"meta_description"`
	DateCreated             time.Time     `json:"date_created"`
	DateModified            time.Time     `json:"date_modified"`
	ViewCount               int           `json:"view_count"`
	PreorderReleaseDate     interface{}   `json:"preorder_release_date"`
	PreorderMessage         string        `json:"preorder_message"`
	IsPreorderOnly          bool          `json:"is_preorder_only"`
	IsPriceHidden           bool          `json:"is_price_hidden"`
	PriceHiddenLabel        string        `json:"price_hidden_label"`
	CustomURL               struct {
		URL          string `json:"url"`
		IsCustomized bool   `json:"is_customized"`
	} `json:"custom_url"`
	BaseVariantID               int64  `json:"base_variant_id"`
	OpenGraphType               string `json:"open_graph_type"`
	OpenGraphTitle              string `json:"open_graph_title"`
	OpenGraphDescription        string `json:"open_graph_description"`
	OpenGraphUseMetaDescription bool   `json:"open_graph_use_meta_description"`
	OpenGraphUseProductName     bool   `json:"open_graph_use_product_name"`
	OpenGraphUseImage           bool   `json:"open_graph_use_image"`
	Variants                    []struct {
		ID                        int64         `json:"id"`
		ProductID                 int64         `json:"product_id"`
		Sku                       string        `json:"sku"`
		SkuID                     interface{}   `json:"sku_id"`
		Price                     float64       `json:"price"`
		CalculatedPrice           float64       `json:"calculated_price"`
		SalePrice                 float64       `json:"sale_price"`
		RetailPrice               float64       `json:"retail_price"`
		MapPrice                  float64       `json:"map_price"`
		Weight                    float64       `json:"weight"`
		Width                     int           `json:"width"`
		Height                    int           `json:"height"`
		Depth                     int           `json:"depth"`
		IsFreeShipping            bool          `json:"is_free_shipping"`
		FixedCostShippingPrice    float64       `json:"fixed_cost_shipping_price"`
		CalculatedWeight          float64       `json:"calculated_weight"`
		PurchasingDisabled        bool          `json:"purchasing_disabled"`
		PurchasingDisabledMessage string        `json:"purchasing_disabled_message"`
		ImageURL                  string        `json:"image_url"`
		CostPrice                 float64       `json:"cost_price"`
		Upc                       string        `json:"upc"`
		Mpn                       string        `json:"mpn"`
		Gtin                      string        `json:"gtin"`
		InventoryLevel            int           `json:"inventory_level"`
		InventoryWarningLevel     int           `json:"inventory_warning_level"`
		BinPickingNumber          string        `json:"bin_picking_number"`
		OptionValues              []interface{} `json:"option_values"`
	} `json:"variants"`
	Images       []interface{} `json:"images"`
	PrimaryImage interface{}   `json:"primary_image"`
	Videos       []interface{} `json:"videos"`
	CustomFields []struct {
		ID    int64  `json:"id"`
		Name  string `json:"name"`
		Value string `json:"value"`
	} `json:"custom_fields"`
	BulkPricingRules []interface{} `json:"bulk_pricing_rules"`
	Options          []interface{} `json:"options"`
	Modifiers        []interface{} `json:"modifiers"`
}

// Metafield is a struct representing a BigCommerce product metafield
type Metafield struct {
	ID            int64     `json:"id"`
	Key           string    `json:"key"`
	Value         string    `json:"value"`
	ResourceID    int64     `json:"resource_id"`
	ResourceType  string    `json:"resource_type"`
	Description   string    `json:"description"`
	DateCreated   time.Time `json:"date_created"`
	DateModified  time.Time `json:"date_modified"`
	Namespace     string    `json:"namespace"`
	PermissionSet string    `json:"permission_set"`
}

// GetAllProducts gets all products from BigCommerce
// args is a key-value map of additional arguments to pass to the API
func (bc *Client) GetAllProducts(args map[string]string) ([]Product, error) {
	ps := []Product{}
	var psp []Product
	page := 1
	more := true
	var err error
	retries := 0
	for more {
		psp, more, err = bc.GetProducts(args, page)
		//		log.Printf("products page %d entries %d %v", page, len(psp), args)
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
// args is a key-value map of additional arguments to pass to the API
// page: the page number to download
func (bc *Client) GetProducts(args map[string]string, page int) ([]Product, bool, error) {
	fpart := ""
	for k, v := range args {
		fpart += "&" + k + "=" + v
	}
	url := "/v3/catalog/products?page=" + strconv.Itoa(page) + fpart
	// log.Printf("GET %s", url)

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
	//	log.Printf("%d products (%+v)", len(pp.Data), pp.Meta.Pagination)

	if pp.Status != 0 {
		return nil, false, errors.New(pp.Title)
	}
	return pp.Data, pp.Meta.Pagination.CurrentPage < pp.Meta.Pagination.TotalPages, nil
}

// GetProductByID gets a product from BigCommerce by ID
// productID: BigCommerce product ID to get
func (bc *Client) GetProductByID(productID int64) (*Product, error) {
	url := "/v3/catalog/products/" + strconv.FormatInt(productID, 10) + "?include=variants,images,custom_fields,bulk_pricing_rules,primary_image,modifiers,options,videos"
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

	var productResponse struct {
		Data Product `json:"data"`
	}
	err = json.Unmarshal(body, &productResponse)
	if err != nil {
		return nil, err
	}
	//	log.Printf("Product %d: %s", productID, string(body))
	return &productResponse.Data, nil
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
		Metafields []Metafield `json:"data"`
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
