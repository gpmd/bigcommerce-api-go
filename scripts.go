package bigcommerce

import (
	"bytes"
	"encoding/json"
	"net/http"
)

type Script struct {
	Name            string `json:"name"`
	Description     string `json:"description"`
	HTML            string `json:"html"`
	Src             string `json:"src"`
	AutoUninstall   bool   `json:"auto_uninstall"`
	LoadMethod      string `json:"load_method"`
	Location        string `json:"location"`
	Visibility      string `json:"visibility"`
	Kind            string `json:"kind"`
	APIClientID     string `json:"api_client_id"`
	ConsentCategory string `json:"consent_category"`
	Enabled         bool   `json:"enabled"`
	ChannelID       int64  `json:"channel_id"`
}

func (bc *Client) CreateScript(s *Script) (*Script, error) {
	sJSON, err := json.Marshal(s)
	if err != nil {
		return nil, err
	}
	req := bc.getAPIRequest(http.MethodPost, "/v3/content/scripts", bytes.NewReader(sJSON))
	res, err := bc.HTTPClient.Do(req)
	if err != nil {
		return s, err
	}
	defer res.Body.Close()
	var sRes struct {
		Data Script `json:"data"`
	}
	err = json.NewDecoder(res.Body).Decode(&sRes)
	return &sRes.Data, err
}
