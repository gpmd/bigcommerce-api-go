package bigcommerce

import (
	"bytes"
	"encoding/json"
	"net/http"
)

type Adjustment struct {
	Reason string `json:"reason"`
	Items  []Item `json:"items"`
}

type Item struct {
	LocationId int    `json:"location_id"`
	VariantId  int    `json:"variant_id,omitempty"`
	Quantity   int    `json:"quantity"`
	Sku        string `json:"sku,omitempty"`
	ProductId  int    `json:"product_id,omitempty"`
}

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
