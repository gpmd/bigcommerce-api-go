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
