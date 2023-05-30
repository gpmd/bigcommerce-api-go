package bigcommerce

import (
	"bytes"
	"encoding/json"
	"net/http"
)

type Adjustment struct {
	Reason string           `json:"reason"`
	Items  []AdjustmentItem `json:"items"`
}

type AdjustmentItem struct {
	LocationId int    `json:"location_id"`
	VariantId  int    `json:"variant_id,omitempty"`
	Quantity   int    `json:"quantity"`
	Sku        string `json:"sku,omitempty"`
	ProductId  int    `json:"product_id,omitempty"`
}

// AdjustInventoryRelative changes the stock value relative to it's current value
func (bc *Client) AdjustInventoryRelative(adjustment *Adjustment) error {
	url := "/v3/inventory/adjustments/relative"

	reqJSON, _ := json.Marshal(adjustment)
	req := bc.getAPIRequest(http.MethodPost, url, bytes.NewReader(reqJSON))
	_, err := bc.HTTPClient.Do(req)
	if err != nil {
		return err
	}

	return nil
}

// AdjustInventoryAbsolute sets the stock value to a specific value
func (bc *Client) AdjustInventoryAbsolute(adjustment *Adjustment) error {
	url := "/v3/inventory/adjustments/absolute"

	reqJSON, _ := json.Marshal(adjustment)
	req := bc.getAPIRequest(http.MethodPost, url, bytes.NewReader(reqJSON))
	_, err := bc.HTTPClient.Do(req)
	if err != nil {
		return err
	}

	return nil
}
