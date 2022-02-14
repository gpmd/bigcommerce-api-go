package bigcommerce

import (
	"net/http"
	"net/url"
	"time"
)

type BigCommerce struct {
	Hostname        string
	AppClientID     string
	AppClientSecret string
	DefaultClient   *http.Client
	MaxRetries      int
}

type BigCommerceAPI interface {
	GetAuthContext(clientId, clientSecret string, q url.Values) (*AuthContext, error)
}

func New(hostname, appClientID, appClientSecret string) *BigCommerce {
	return &BigCommerce{
		Hostname:        hostname,
		AppClientID:     appClientID,
		AppClientSecret: appClientSecret,
		MaxRetries:      5,
		DefaultClient: &http.Client{
			Timeout: time.Second * 10,
		},
	}
}

func (bc *BigCommerce) getAPIRequest(method, url, client, token string) *http.Request {
	req, _ := http.NewRequest(method, "https://api.bigcommerce.com/"+url, nil)

	req.Header.Add("X-Auth-Client", client)
	req.Header.Add("X-Auth-Token", token)
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("User-Agent", "GPMD Blog Post App")
	req.Header.Add("Cache-Control", "no-cache")
	req.Header.Add("Host", "api.bigcommerce.com")
	req.Header.Add("Accept-Encoding", "none")
	req.Header.Add("Connection", "keep-alive")
	return req
}
