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
}

type BigCommerceAPI interface {
	GetAuthContext(clientId, clientSecret string, q url.Values) (*AuthContext, error)
}

func New(hostname, appClientID, appClientSecret string) *BigCommerce {
	return &BigCommerce{
		Hostname:        hostname,
		AppClientID:     appClientID,
		AppClientSecret: appClientSecret,
		DefaultClient: &http.Client{
			Timeout: time.Second * 10,
		},
	}
}
