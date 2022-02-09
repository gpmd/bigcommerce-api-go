package bigcommerce

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/url"
	"strings"
)

// GetAuthContext returns an AuthContext object from the BigCommerce API
// Call it with r.URL.Query() - will return BigCommerce Auth Context or error
func (bc *BigCommerce) GetAuthContext(requestURLQuery url.Values) (*AuthContext, error) {

	req := AuthTokenRequest{
		ClientID:     bc.AppClientID,
		ClientSecret: bc.AppClientSecret,
		RedirectURI:  "https://" + bc.Hostname + "/auth",
		GrantType:    "authorization_code",
		Code:         requestURLQuery.Get("code"),
		Scope:        requestURLQuery.Get("scope"),
		Context:      requestURLQuery.Get("context"),
	}
	reqb, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}

	res, err := bc.DefaultClient.Post("https://login.bigcommerce.com/oauth2/token",
		"application/json",
		bytes.NewReader(reqb),
	)
	if err != nil {
		return nil, err
	}

	bytes, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	if strings.Contains(string(bytes), "invalid_") {
		return nil, fmt.Errorf("%s", string(bytes))
	}

	var ac AuthContext
	err = json.Unmarshal(bytes, &ac)
	if err != nil {
		return nil, err
	}
	return &ac, nil
}
