# Bigcommerce API Client for Go

## Usage

```go
package main

import (
    "github.com/gpmd/bigcommerce-api-go"
    "log"
)

func main() {
    // for app development, you need to provide arguments
    // for CLI and other web apps you can use empty strings
    client := bigcommerce.NewClient("** my store's hash like '123abcdefg' **", "**my X-Auth-Token generated in BigCommerce admin**")
    products, err := client.GetAllProducts()
    if err != nil {
        log.Fatalf("Error while getting products: %v", err)
    }
    for _, product := range products {
        log.Println(product.Name)
    }
}
```

## Errors

```go
var ErrNoContent = errors.New("no content 204 from BigCommerce API")
var ErrNoMainThumbnail = errors.New("no main thumbnail")
var ErrNotFound = errors.New("404 not found")
```

## Types

#### type App

```go
type App struct {
	Hostname        string
	AppClientID     string
	AppClientSecret string
	HTTPClient      *http.Client
	MaxRetries      int
	ChannelID       int
}
```

BigCommerce is the BigCommerce API client object for BigCommerce Apps holds no
client specific information

#### func  NewApp

```go
func NewApp(hostname, appClientID, appClientSecret string) *App
```
New returns a new BigCommerce API object with the given hostname, client ID, and
client secret The client ID and secret are the App's client ID and secret from
the BigCommerce My Apps dashboard The hostname is the domain name of the app
from the same page (e.g. app.exampledomain.com)

#### func (*App) CheckSignature

```go
func (bc *App) CheckSignature(signedPayload string) ([]byte, error)
```
CheckSignature checks the signature of the request whith SHA256 HMAC

#### func (*App) GetAuthContext

```go
func (bc *App) GetAuthContext(requestURLQuery url.Values) (*AuthContext, error)
```
GetAuthContext returns an AuthContext object from the BigCommerce API Call it
with r.URL.Query() - will return BigCommerce Auth Context or error

#### func (*App) GetClientRequest

```go
func (bc *App) GetClientRequest(requestURLQuery url.Values) (*ClientRequest, error)
```
GetClientRequest returns a ClientRequest object from the BigCommerce API Call it
with r.URL.Query() - will return BigCommerce Client Request or error

#### type AuthContext

```go
type AuthContext struct {
	AccessToken string `json:"access_token"` // used later as X-Auth-Token header
	Scope       string `json:"scope"`
	User        BCUser `json:"user"`
	Context     string `json:"context"`
	URL         string `json:"url"`
	Error       string `json:"error"`
}
```

AuthContext is a BigCommerce auth context object

#### type AuthContexter

```go
type AuthContexter interface {
	GetAuthContext(clientID, clientSecret string, q url.Values) (*AuthContext, error)
}
```

AuthContexter interface for GetAuthContext

#### type AuthTokenRequest

```go
type AuthTokenRequest struct {
	ClientID     string `json:"client_id"`
	ClientSecret string `json:"client_secret"`
	Code         string `json:"code"`
	Scope        string `json:"scope"`
	GrantType    string `json:"grant_type"`
	RedirectURI  string `json:"redirect_uri"`
	Context      string `json:"context"`
}
```

AuthTokenRequest is sent to BigCommerce to get AuthContext

#### type BCUser

```go
type BCUser struct {
	ID       int64  `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
}
```

BCUser is a BigCommerce shorthand object type that's in many other responses

#### type Brand

```go
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
```

Brand is BigCommerce brand object

#### type Cart

```go
type Cart struct {
	ID         string `json:"id"`
	CustomerID int64  `json:"customer_id,omitempty"`
	ChannelID  int64  `json:"channel_id,omitempty"`
	Email      string `json:"email,omitempty"`
	Currency   struct {
		Code string `json:"code,omitempty"`
	} `json:"currency,omitempty"`
	TaxIncluded    bool       `json:"tax_included,omitempty"`
	BaseAmount     float64    `json:"base_amount,omitempty"`
	DiscountAmount float64    `json:"discount_amount,omitempty"`
	CartAmount     float64    `json:"cart_amount,omitempty"`
	Discounts      []Discount `json:"discounts,omitempty"`
	Coupons        []Coupon   `json:"coupons,omitempty"`
	LineItems      struct {
		PhysicalItems    []LineItem `json:"physical_items,omitempty"`
		DigitalItems     []LineItem `json:"digital_items,omitempty"`
		GiftCertificates []LineItem `json:"gift_certificates,omitempty"`
		CustomItems      []LineItem `json:"custom_items,omitempty"`
	} `json:"line_items"`
	CreatedTime time.Time `json:"created_time,omitempty"`
	UpdatedTime time.Time `json:"updated_time,omitempty"`
	Locale      string    `json:"locale,omitempty"`
}
```

Cart is a BigCommerce cart object

#### type Category

```go
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
```

Category is a BC category object

#### type Channel

```go
type Channel struct {
	IconURL          string    `json:"icon_url"`
	IsListableFromUI bool      `json:"is_listable_from_ui"`
	IsVisible        bool      `json:"is_visible"`
	DateCreated      time.Time `json:"date_created"`
	ExternalID       string    `json:"external_id"`
	Type             string    `json:"type"`
	Platform         string    `json:"platform"`
	IsEnabled        bool      `json:"is_enabled"`
	DateModified     time.Time `json:"date_modified"`
	Name             string    `json:"name"`
	ID               int       `json:"id"`
	Status           string    `json:"status"`
}
```


#### type Client

```go
type Client struct {
	StoreHash  string `json:"store-hash"`
	XAuthToken string `json:"x-auth-token"`
	MaxRetries int
	HTTPClient *http.Client
	ChannelID  int
}
```


#### func  NewClient

```go
func NewClient(storeHash, xAuthToken string) *Client
```

#### func (*Client) CartAddItems

```go
func (bc *Client) CartAddItems(cartID string, items []LineItem) (*Cart, error)
```
CartAddItem adds line items to a cart

#### func (*Client) CartDeleteItem

```go
func (bc *Client) CartDeleteItem(cartID string, item LineItem) (*Cart, error)
```
DeleteItem deletes a line item from a cart, returns the updated cart Arguments:

    cartID: the cart ID
    item: the line item, must have an existing line item ID

#### func (*Client) CartEditItem

```go
func (bc *Client) CartEditItem(cartID string, item LineItem) (*Cart, error)
```
EditItem edits a line item in a cart, returns the updated cart Arguments:

    cartID: the cart ID
    item: the line item to edit. Must have an ID, quantity, and product ID

#### func (*Client) CartUpdateCustomerID

```go
func (bc *Client) CartUpdateCustomerID(cartID, customerID string) (*Cart, error)
```
CartUpdateCustomerID updates the customer ID for a cart Arguments: cartID: the
BigCommerce cart ID customerID: the new BigCommerce customer ID

#### func (*Client) CreateCart

```go
func (bc *Client) CreateCart(items []LineItem) (*Cart, error)
```
CreateCart creates a new cart in BigCommerce and returns it

#### func (*Client) GetAllBrands

```go
func (bc *Client) GetAllBrands() ([]Brand, error)
```
GetAllBrands returns all brands, handling pagination

#### func (*Client) GetAllCategories

```go
func (bc *Client) GetAllCategories() ([]Category, error)
```
GetAllCategories returns a list of categories, handling pagination

#### func (*Client) GetAllChannels

```go
func (bc *Client) GetAllChannels() ([]Channel, error)
```

#### func (*Client) GetAllPosts

```go
func (bc *Client) GetAllPosts(context, xAuthToken string) ([]Post, error)
```
GetAllPosts downloads all posts from BigCommerce, handling pagination context:
the BigCommerce context (e.g. stores/23412341234) where 23412341234 is the store
hash xAuthToken: the BigCommerce Store's X-Auth-Token coming from store
credentials (see AuthContext)

#### func (*Client) GetAllProducts

```go
func (bc *Client) GetAllProducts() ([]Product, error)
```
GetAllProducts gets all products from BigCommerce

#### func (*Client) GetBrands

```go
func (bc *Client) GetBrands(page int) ([]Brand, bool, error)
```
GetBrands returns all brands, handling pagination context: the BigCommerce
context (e.g. stores/23412341234) where 23412341234 is the store hash
xAuthToken: the BigCommerce Store's X-Auth-Token coming from store credentials
(see AuthContext) page: the page number to download

#### func (*Client) GetCart

```go
func (bc *Client) GetCart(cartID string) (*Cart, error)
```
GetCart gets a cart by ID from BigCommerce and returns it

#### func (*Client) GetCategories

```go
func (bc *Client) GetCategories(page int) ([]Category, bool, error)
```
GetCategories returns a list of categories, handling pagination page: the page
number to download

#### func (*Client) GetChannels

```go
func (bc *Client) GetChannels(page int) ([]Channel, bool, error)
```

#### func (*Client) GetMainThumbnailURL

```go
func (bc *Client) GetMainThumbnailURL(productID int64) (string, error)
```
GetMainThumbnailURL returns the main thumbnail URL for a product this is due to
the fact that the Product API does not return the main thumbnail URL

#### func (*Client) GetPosts

```go
func (bc *Client) GetPosts(page int) ([]Post, bool, error)
```
GetPosts downloads all posts from BigCommerce, handling pagination context: the
BigCommerce context (e.g. stores/23412341234) where 23412341234 is the store
hash xAuthToken: the BigCommerce Store's X-Auth-Token coming from store
credentials (see AuthContext) page: the page number to download

#### func (*Client) GetProductByID

```go
func (bc *Client) GetProductByID(productID int64) (*Product, error)
```
GetProductByID gets a product from BigCommerce by ID productID: BigCommerce
product ID to get

#### func (*Client) GetProducts

```go
func (bc *Client) GetProducts(page int) ([]Product, bool, error)
```
GetProducts gets a page of products from BigCommerce page: the page number to
download

#### func (*Client) GetStoreInfo

```go
func (bc *Client) GetStoreInfo() (StoreInfo, error)
```
GetStoreInfo returns the store info for the current store page: the page number
to download

#### func (*Client) SetProductFields

```go
func (bc *Client) SetProductFields(fields []string)
```
SetProductFields sets include_fields parameter for GetProducts, empty list will
get all fields

#### func (*Client) SetProductInclude

```go
func (bc *Client) SetProductInclude(subresources []string)
```
SetProductFields sets include_fields parameter for GetProducts, empty list will
get all fields

#### func (*Client) ValidateCredentials

```go
func (bc *Client) ValidateCredentials(email, password string) (int64, error)
```

#### type ClientRequest

```go
type ClientRequest struct {
	User      UserPart `json:"user"`
	Owner     UserPart `json:"owner"`
	Context   string   `json:"context"`
	StoreHash string   `json:"store_hash"`
}
```

ClientRequest is a BigCommerce client request object that comes with most App
callbacks in the GET request signed_payload parameter

#### type Coupon

```go
type Coupon struct {
	Code             string  `json:"code"`
	ID               string  `json:"id"`
	CouponType       string  `json:"coupon_type"`
	DiscountedAmount float64 `json:"discounted_amount"`
}
```


#### type Customer

```go
type Customer struct {
	ID               int64             `json:"id"`
	Company          string            `json:"company"`
	Firstname        string            `json:"first_name"`
	Lastname         string            `json:"last_name"`
	Email            string            `json:"email"`
	Phone            string            `json:"phone"`
	FormFields       interface{}       `json:"form_fields"`
	DateCreated      string            `json:"date_created"`
	DateModified     string            `json:"date_modified"`
	StoreCredit      string            `json:"store_credit"`
	RegistrationIP   string            `json:"registration_ip_address"`
	CustomerGroup    int64             `json:"customer_group_id"`
	Notes            string            `json:"notes"`
	TaxExempt        string            `json:"tax_exempt_category"`
	ResetPassword    bool              `json:"reset_pass_on_login"`
	AcceptsMarketing bool              `json:"accepts_marketing"`
	Addresses        map[string]string `json:"addresses"`
}
```

Customer is a struct for the BigCommerce Customer API

#### type Discount

```go
type Discount struct {
	ID               string  `json:"id"`
	DiscountedAmount float64 `json:"discounted_amount"`
}
```


#### type Image

```go
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
```

Image is entry for BC product images

#### type InventoryEntry

```go
type InventoryEntry struct {
	ProductID int64   `json:"product_id"`
	Method    string  `json:"method"`
	Value     float64 `json:"value"`
	VariantID int64   `json:"variant_id"`
}
```


#### type LineItem

```go
type LineItem struct {
	ID                string     `json:"id,omitempty"`
	ParentID          int64      `json:"parent_id,omitempty"`
	VariantID         int64      `json:"variant_id,omitempty"`
	ProductID         int64      `json:"product_id,omitempty"`
	Sku               string     `json:"sku,omitempty"`
	Name              string     `json:"name,omitempty"`
	URL               string     `json:"url,omitempty"`
	Quantity          float64    `json:"quantity,omitempty"`
	Taxable           bool       `json:"taxable,omitempty"`
	ImageURL          string     `json:"image_url,omitempty"`
	Discounts         []Discount `json:"discounts,omitempty"`
	Coupons           []Coupon   `json:"coupons,omitempty"`
	DiscountAmount    float64    `json:"discount_amount,omitempty"`
	CouponAmount      float64    `json:"coupon_amount,omitempty"`
	ListPrice         float64    `json:"list_price,omitempty"`
	SalePrice         float64    `json:"sale_price,omitempty"`
	ExtendedListPrice float64    `json:"extended_list_price,omitempty"`
	ExtendedSalePrice float64    `json:"extended_sale_price,omitempty"`
	IsRequireShipping bool       `json:"is_require_shipping,omitempty"`
	IsMutable         bool       `json:"is_mutable,omitempty"`
}
```

LineItem is a BigCommerce line item object for cart

#### type LoadContext

```go
type LoadContext struct {
	User      BCUser  `json:"user"`
	Owner     BCUser  `json:"owner"`
	Context   string  `json:"context"`
	StoreHash string  `json:"store_hash"`
	Timestamp float64 `json:"timestamp"`
	URL       string  `json:"url"`
}
```

LoadContext is a BigCommerce load context object

#### type Post

```go
type Post struct {
	ID                   int64       `json:"id"`
	Title                string      `json:"title"`
	URL                  string      `json:"url"`
	PreviewURL           string      `json:"preview_url"`
	Body                 string      `json:"body"`
	Tags                 []string    `json:"tags"`
	Summary              string      `json:"summary"`
	IsPublished          bool        `json:"is_published"`
	PublishedDate        interface{} `json:"publisheddate"`
	PublishedDateISO8601 string      `json:"publisheddate_iso8601"`
	MetaDescription      string      `json:"meta_description"`
	MetaKeywords         string      `json:"meta_keywords"`
	Author               string      `json:"author"`
	ThumbnailPath        string      `json:"thumbnailpath"`
}
```

Post is a BC blog post

#### type Product

```go
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
```

Product is a BigCommerce product object

#### type StoreInfo

```go
type StoreInfo struct {
	ID          string `json:"id"`
	Domain      string `json:"domain"`
	SecureURL   string `json:"secure_url"`
	Status      string `json:"status"`
	Name        string `json:"name"`
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	Address     string `json:"address"`
	Country     string `json:"country"`
	CountryCode string `json:"country_code"`
	Phone       string `json:"phone"`
	AdminEmail  string `json:"admin_email"`
	OrderEmail  string `json:"order_email"`
	FaviconURL  string `json:"favicon_url"`
	Timezone    struct {
		Name          string `json:"name"`
		RawOffset     int    `json:"raw_offset"`
		DstOffset     int    `json:"dst_offset"`
		DstCorrection bool   `json:"dst_correction"`
		DateFormat    struct {
			Display         string `json:"display"`
			Export          string `json:"export"`
			ExtendedDisplay string `json:"extended_display"`
		} `json:"date_format"`
	} `json:"timezone"`
	Language                string        `json:"language"`
	Currency                string        `json:"currency"`
	CurrencySymbol          string        `json:"currency_symbol"`
	DecimalSeparator        string        `json:"decimal_separator"`
	ThousandsSeparator      string        `json:"thousands_separator"`
	DecimalPlaces           int           `json:"decimal_places"`
	CurrencySymbolLocation  string        `json:"currency_symbol_location"`
	WeightUnits             string        `json:"weight_units"`
	DimensionUnits          string        `json:"dimension_units"`
	DimensionDecimalPlaces  int           `json:"dimension_decimal_places"`
	DimensionDecimalToken   string        `json:"dimension_decimal_token"`
	DimensionThousandsToken string        `json:"dimension_thousands_token"`
	PlanName                string        `json:"plan_name"`
	PlanLevel               string        `json:"plan_level"`
	PlanIsTrial             bool          `json:"plan_is_trial"`
	Industry                string        `json:"industry"`
	Logo                    interface{}   `json:"logo"`
	IsPriceEnteredWithTax   bool          `json:"is_price_entered_with_tax"`
	ActiveComparisonModules []interface{} `json:"active_comparison_modules"`
	Features                struct {
		StencilEnabled       bool   `json:"stencil_enabled"`
		SitewidehttpsEnabled bool   `json:"sitewidehttps_enabled"`
		FacebookCatalogID    string `json:"facebook_catalog_id"`
		CheckoutType         string `json:"checkout_type"`
		WishlistsEnabled     bool   `json:"wishlists_enabled"`
	} `json:"features"`
}
```

StoreInfo is a BigCommerce store info object

#### type UserPart

```go
type UserPart struct {
	ID    int64  `json:"id"`
	Email string `json:"email"`
}
```

UserPart is a BigCommerce user shorthand object type that's in many other
responses

#### type WebhookPayload

```go
type WebhookPayload struct {
	Scope   string `json:"scope"`
	StoreID string `json:"store_id"`
	Data    struct {
		Type     string `json:"type"`
		ID       int64  `json:"id"`
		CouponID string `json:"couponId"`
		CartID   string `json:"cartId"`
		OrderID  int64  `json:"orderId"`
		Address  struct {
			CustomerID int64 `json:"customer_id"`
		} `json:"address"`
		Inventory InventoryEntry `json:"inventory"`
		Message   struct {
			OrderMessageID int64 `json:"order_message_id"`
		} `json:"message"`
		Sku struct {
			ProductID int64 `json:"product_id"`
			VariantID int64 `json:"variant_id"`
		} `json:"sku"`
		Status struct {
			PreviousStatusID int64 `json:"previous_status_id"`
			NewStatusID      int64 `json:"new_status_id"`
		} `json:"status"`
	} `json:"data"`
	Hash      string `json:"hash"`
	CreatedAt int64  `json:"created_at"`
	Producer  string `json:"producer"`
}
```


#### func  GetWebhookPayload

```go
func GetWebhookPayload(r *http.Request) (*WebhookPayload, []byte, error)
```
GetWebhookPayload returns a WebhookPayload object and the raw payload from the
BigCommerce API Arguments: r - the http.Request object Returns: *WebhookPayload
- the WebhookPayload object []byte - the raw payload from the BigCommerce API
error - the error, if any
