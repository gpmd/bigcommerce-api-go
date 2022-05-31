package bigcommerce

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

type PageBuilderTemplate struct {
	ChannelID          int64         `json:"channel_id"`
	ClientRerender     bool          `json:"client_rerender"`
	CurrentVersionUUID string        `json:"current_version_uuid"`
	DateCreated        time.Time     `json:"date_created"`
	DateModified       time.Time     `json:"date_modified"`
	IconName           string        `json:"icon_name"`
	Kind               string        `json:"kind"`
	Name               string        `json:"name"`
	Schema             []interface{} `json:"schema"`
	StorefrontAPIQuery string        `json:"storefront_api_query"`
	Template           string        `json:"template"`
	TemplateEngine     string        `json:"template_engine"`
	UUID               string        `json:"uuid"`
}

func (bc *Client) CreateWidgetTemplate(pt *PageBuilderTemplate) (*PageBuilderTemplate, error) {
	ptJSON, err := json.Marshal(pt)
	if err != nil {
		return nil, err
	}
	req := bc.getAPIRequest(http.MethodPost, "/v3/content/widget-templates", bytes.NewReader(ptJSON))
	res, err := bc.HTTPClient.Do(req)
	if err != nil {
		return pt, err
	}
	defer res.Body.Close()
	var ptRes struct {
		Data PageBuilderTemplate `json:"data"`
	}
	b, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return pt, err
	}
	err = json.Unmarshal(b, &ptRes)
	if ptRes.Data.UUID == "" {
		return pt, fmt.Errorf("error creating widget template: %s", string(b))
	}
	return &ptRes.Data, err
}

func (bc *Client) GetWidgetTemplates() ([]PageBuilderTemplate, error) {
	req := bc.getAPIRequest(http.MethodGet, "/v3/content/widget-templates", nil)
	res, err := bc.HTTPClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	var ptRes struct {
		Data []PageBuilderTemplate `json:"data"`
	}
	b, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(b, &ptRes)
	if ptRes.Data == nil {
		return nil, fmt.Errorf("error getting widget templates: %s", string(b))
	}
	return ptRes.Data, err
}

func (bc *Client) DeleteWidgetTemplate(uuid string) error {
	req := bc.getAPIRequest(http.MethodDelete, fmt.Sprintf("/v3/content/widget-templates/%s", uuid), nil)
	res, err := bc.HTTPClient.Do(req)
	if err != nil {
		return err
	}
	defer res.Body.Close()
	if res.StatusCode != http.StatusNoContent {
		return fmt.Errorf("error deleting widget template: %d", res.StatusCode)
	}
	return nil
}
