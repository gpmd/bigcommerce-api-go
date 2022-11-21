package bigcommerce

import (
	"encoding/json"
	"net/http"
)

type CustomerGroup struct {
	ID               int64          `json:"id"`
	Name             string         `json:"name"`
	IsDefault        bool           `json:"is_default"`
	CategoryAccess   CategoryAccess `json:"category_access"`
	DiscountRules    []DiscountRule `json:"discount_rules"`
	IsGroupForGuests bool           `json:"is_group_for_guests"`
}

type CategoryAccess struct {
	Type       string  `json:"type"`
	Categories []int64 `json:"categories"`
}

type DiscountRule struct {
	Type        string `json:"type"`
	Method      string `json:"method"`
	Amount      string `json:"amount"`
	PriceListID int64  `json:"price_list_id"`
}

func (bc *Client) GetCustomerGroups() ([]CustomerGroup, error) {
	req := bc.getAPIRequest(http.MethodGet, "/v2/customer_groups", nil)
	res, err := bc.HTTPClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	body, err := processBody(res)
	if err != nil {
		return nil, err
	}
	var ret []CustomerGroup
	err = json.Unmarshal(body, &ret)
	return ret, err
}
