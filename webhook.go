package bigcommerce

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

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
