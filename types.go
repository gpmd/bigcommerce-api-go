package bigcommerce

// WebhookPayload is a BigCommerce webhook payload object
type WebhookPayload struct {
	Scope     string                 `json:"scope"`
	StoreID   string                 `json:"store_id"`
	Data      map[string]interface{} `json:"data"`
	Hash      string                 `json:"hash"`
	CreatedAt string                 `json:"created_at"`
	Producer  string                 `json:"producer"`
}

// AuthTokenRequest is sent to BigCommerce to get AuthContext
type AuthTokenRequest struct {
	ClientID     string `json:"client_id"`
	ClientSecret string `json:"client_secret"`
	Code         string `json:"code"`
	Scope        string `json:"scope"`
	GrantType    string `json:"grant_type"`
	RedirectURI  string `json:"redirect_uri"`
	Context      string `json:"context"`
}

// BCUser is a BigCommerce shorthand object type that's in many other responses
type BCUser struct {
	ID       int64  `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
}

// LoadContext is a BigCommerce load context object
type LoadContext struct {
	User      BCUser  `json:"user"`
	Owner     BCUser  `json:"owner"`
	Context   string  `json:"context"`
	StoreHash string  `json:"store_hash"`
	Timestamp float64 `json:"timestamp"`
	URL       string  `json:"url"`
}

// StoreInfo is a BigCommerce store info object
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

// AuthContext is a BigCommerce auth context object
type AuthContext struct {
	AccessToken string `json:"access_token"` // used later as X-Auth-Token header
	Scope       string `json:"scope"`
	User        BCUser `json:"user"`
	Context     string `json:"context"`
	URL         string `json:"url"`
	Error       string `json:"error"`
}

// UserPart is a BigCommerce user shorthand object type that's in many other responses
type UserPart struct {
	ID    int64  `json:"id"`
	Email string `json:"email"`
}

// ClientRequest is a BigCommerce client request object that comes with most App callbacks
// in the GET request signed_payload parameter
type ClientRequest struct {
	User      UserPart `json:"user"`
	Owner     UserPart `json:"owner"`
	Context   string   `json:"context"`
	StoreHash string   `json:"store_hash"`
}
