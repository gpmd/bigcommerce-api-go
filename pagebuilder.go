package bigcommerce

import (
	"bytes"
	"encoding/json"
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
	err = json.NewDecoder(res.Body).Decode(&ptRes)
	return &ptRes.Data, err
}
