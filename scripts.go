package bigcommerce

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

type Script struct {
	ID              string    `json:"uuid"`
	DateCreated     time.Time `json:"date_created"`
	DateModified    time.Time `json:"date_modified"`
	Name            string    `json:"name"`
	Description     string    `json:"description"`
	HTML            string    `json:"html"`
	Src             string    `json:"src"`
	AutoUninstall   bool      `json:"auto_uninstall"`
	LoadMethod      string    `json:"load_method"`
	Location        string    `json:"location"`
	Visibility      string    `json:"visibility"`
	Kind            string    `json:"kind"`
	APIClientID     string    `json:"api_client_id"`
	ConsentCategory string    `json:"consent_category"`
	Enabled         bool      `json:"enabled"`
	ChannelID       int64     `json:"channel_id"`
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
	b, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return s, err
	}
	err = json.Unmarshal(b, &sRes)
	if sRes.Data.ID == "" {
		return s, fmt.Errorf("error creating script: %s", string(b))
	}
	return &sRes.Data, err
}

func (bc *Client) GetScriptByID(uuid string) (*Script, error) {
	req := bc.getAPIRequest(http.MethodGet, fmt.Sprintf("/v3/content/scripts/%s", uuid), nil)
	res, err := bc.HTTPClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	var sRes struct {
		Data Script `json:"data"`
	}
	b, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(b, &sRes)
	if sRes.Data.ID == "" {
		return nil, fmt.Errorf("error getting script: %s", string(b))
	}
	return &sRes.Data, err
}

func (bc *Client) GetScripts() ([]Script, error) {
	req := bc.getAPIRequest(http.MethodGet, "/v3/content/scripts", nil)
	res, err := bc.HTTPClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	var sRes struct {
		Data []Script `json:"data"`
	}
	b, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(b, &sRes)
	if err != nil {
		return nil, fmt.Errorf("error getting scripts: %v %s", err, string(b))
	}
	return sRes.Data, err
}
