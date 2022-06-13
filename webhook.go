package bigcommerce

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"reflect"
	"strconv"
	"strings"
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

type Webhook struct {
	ID          int64             `json:"id"`
	ClientID    string            `json:"client_id"`
	StoreHash   string            `json:"store_hash"`
	CreatedAt   int64             `json:"created_at"`
	UpdatedAt   int64             `json:"updated_at"`
	Scope       string            `json:"scope"`
	Destination string            `json:"destination"`
	IsActive    bool              `json:"is_active"`
	Headers     map[string]string `json:"headers"`
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

func (bc *Client) GetWebhooks() ([]Webhook, error) {
	url := "/v3/hooks?limit=250"

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

	var webhooksResponse struct {
		Data []Webhook `json:"data"`
		Meta struct {
			Pagination Pagination `json:"pagination"`
		} `json:"meta"`
	}
	err = json.Unmarshal(body, &webhooksResponse)
	if err != nil {
		return nil, err
	}
	return webhooksResponse.Data, nil
}

// CreateWebhook creates a new webhook or activates it if it already exists but inactive
func (bc *Client) CreateWebhook(scope, destination string, headers map[string]string) (int64, error) {
	url := "/v3/hooks"

	webhooks, err := bc.GetWebhooks()
	if err != nil {
		return 0, err
	}
	for _, webhook := range webhooks {
		if webhook.Scope == scope && webhook.Destination == destination && reflect.DeepEqual(headers, headers) {
			if webhook.IsActive {
				return webhook.ID, nil
			}
			req := bc.getAPIRequest(http.MethodPut, url+"/"+strconv.FormatInt(webhook.ID, 10), strings.NewReader(`{"is_active": true}`))
			res, err := bc.HTTPClient.Do(req)
			if err != nil {
				return 0, err
			}
			defer res.Body.Close()
			body, err := processBody(res)
			if err != nil {
				return 0, fmt.Errorf("error processing response body: %v %s", err, string(body))
			}
			return webhook.ID, nil
		}
	}

	payload := struct {
		Scope       string            `json:"scope"`
		Destination string            `json:"destination"`
		Headers     map[string]string `json:"headers"`
	}{
		Scope:       scope,
		Destination: destination,
	}
	if headers != nil {
		payload.Headers = headers
	}
	reqJSON, _ := json.Marshal(payload)

	req := bc.getAPIRequest(http.MethodPost, url, bytes.NewReader(reqJSON))
	res, err := bc.HTTPClient.Do(req)
	if err != nil {
		return 0, err
	}

	defer res.Body.Close()
	body, err := processBody(res)
	if err != nil {
		return 0, fmt.Errorf("error processing response body: %v %s (%s)", err, string(body), string(reqJSON))
	}
	var respWebhook Webhook
	err = json.Unmarshal(body, &respWebhook)
	if err != nil {
		return 0, err
	}
	return respWebhook.ID, nil
}
