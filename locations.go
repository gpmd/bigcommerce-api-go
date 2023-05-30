package bigcommerce

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

type Location struct {
	Code                    string                 `json:"code"`
	Label                   string                 `json:"label"`
	Description             string                 `json:"description"`
	ManagedByExternalSource bool                   `json:"managed_by_external_source"`
	TypeId                  string                 `json:"type_id"`
	Enabled                 bool                   `json:"enabled"`
	OperatingHours          LocationOpeningHours   `json:"operating_hours"`
	TimeZone                string                 `json:"time_zone"`
	Address                 LocationAddress        `json:"address"`
	StorefrontVisibility    bool                   `json:"storefront_visibility"`
	SpecialHours            []LocationSpecialHours `json:"special_hours"`
}

type LocationAddress struct {
	Address1       string `json:"address1"`
	Address2       string `json:"address2"`
	City           string `json:"city"`
	State          string `json:"state"`
	Zip            string `json:"zip"`
	Email          string `json:"email"`
	Phone          string `json:"phone"`
	GeoCoordinates struct {
		Latitude  float64 `json:"latitude"`
		Longitude float64 `json:"longitude"`
	} `json:"geo_coordinates"`
	CountryCode string `json:"country_code"`
}

type LocationOpeningHours struct {
	Sunday struct {
		Open    bool   `json:"open"`
		Opening string `json:"opening"`
		Closing string `json:"closing"`
	} `json:"sunday"`
	Monday struct {
		Open    bool   `json:"open"`
		Opening string `json:"opening"`
		Closing string `json:"closing"`
	} `json:"monday"`
	Tuesday struct {
		Open    bool   `json:"open"`
		Opening string `json:"opening"`
		Closing string `json:"closing"`
	} `json:"tuesday"`
	Wednesday struct {
		Open    bool   `json:"open"`
		Opening string `json:"opening"`
		Closing string `json:"closing"`
	} `json:"wednesday"`
	Thursday struct {
		Open    bool   `json:"open"`
		Opening string `json:"opening"`
		Closing string `json:"closing"`
	} `json:"thursday"`
	Friday struct {
		Open    bool   `json:"open"`
		Opening string `json:"opening"`
		Closing string `json:"closing"`
	} `json:"friday"`
	Saturday struct {
		Open    bool   `json:"open"`
		Opening string `json:"opening"`
		Closing string `json:"closing"`
	} `json:"saturday"`
}

type LocationSpecialHours struct {
	Label   string `json:"label"`
	Date    string `json:"date"`
	Open    bool   `json:"open"`
	Opening string `json:"opening"`
	Closing string `json:"closing"`
	AllDay  bool   `json:"all_day"`
	Annual  bool   `json:"annual"`
}

// GetLocations returns all locations using filters.
// filters: request query parameters for BigCommerce locations endpoint, for example {"is_active": true}
func (bc *Client) GetLocations(filters map[string]string) ([]Location, error) {
	var params []string
	for k, v := range filters {
		params = append(params, fmt.Sprintf("%s=%s", k, v))
	}
	url := "/v3/inventory/locations?" + strings.Join(params, "&")

	req := bc.getAPIRequest(http.MethodGet, url, nil)
	res, err := bc.HTTPClient.Do(req)
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()
	body, err := processBody(res)
	if err != nil {
		if res.StatusCode == http.StatusNoContent {
			return []Location{}, nil
		}
		return nil, err
	}

	var locations []Location
	err = json.Unmarshal(body, &locations)
	if err != nil {
		return nil, err
	}
	return locations, nil
}

// CreateLocation creates a new location based on the Location struct
func (bc *Client) CreateLocation(location *Location) error {
	url := "/v3/inventory/locations"

	reqJSON, _ := json.Marshal(location)
	req := bc.getAPIRequest(http.MethodPost, url, bytes.NewReader(reqJSON))
	_, err := bc.HTTPClient.Do(req)
	if err != nil {
		return err
	}

	return nil
}

// UpdateLocation alters the locations values
func (bc *Client) UpdateLocation(location *Location) error {
	url := "/v3/inventory/locations"

	reqJSON, _ := json.Marshal(location)
	req := bc.getAPIRequest(http.MethodPut, url, bytes.NewReader(reqJSON))
	_, err := bc.HTTPClient.Do(req)
	if err != nil {
		return err
	}

	return nil
}
