package bigcommerce

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

type WebhookPayload struct {
	Scope     string                 `json:"scope"`
	StoreID   string                 `json:"store_id"`
	Data      map[string]interface{} `json:"data"`
	Hash      string                 `json:"hash"`
	CreatedAt string                 `json:"created_at"`
	Producer  string                 `json:"producer"`
}

// GetWebhookPayload returns a WebhookPayload object and the raw payload from the BigCommerce API
// Arguments: r - the http.Request object
// Returns:
// *WebhookPayload - the WebhookPayload object
// []byte - the raw payload from the BigCommerce API
// error - the error, if any
func GetWebhookPayload(r *http.Request) (*WebhookPayload, []byte, error) {
	bytes, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return nil, nil, err
	}
	r.Body.Close()
	var payload WebhookPayload
	err = json.Unmarshal(bytes, &payload)
	if err != nil {
		return nil, bytes, err
	}
	return &payload, bytes, nil
}
