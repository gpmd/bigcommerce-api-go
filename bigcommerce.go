package bigcommerce

import (
	"errors"
	"io"
	"net/http"
	"net/url"
	"time"
)

// BigCommerce is the BigCommerce API client object for BigCommerce Apps
// holds no client specific information
type BigCommerce struct {
	Hostname        string
	AppClientID     string
	AppClientSecret string
	DefaultClient   *http.Client
	MaxRetries      int
}

var ErrNoContent = errors.New("no content 204 from BigCommerce API")

// AuthContexter interface for GetAuthContext
type AuthContexter interface {
	GetAuthContext(clientID, clientSecret string, q url.Values) (*AuthContext, error)
}

// New returns a new BigCommerce API object with the given hostname, client ID, and client secret
// The client ID and secret are the App's client ID and secret from the BigCommerce My Apps dashboard
// The hostname is the domain name of the app from the same page (e.g. app.exampledomain.com)
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

func (bc *BigCommerce) getAPIRequest(method, url, xAuthToken string, body io.Reader) *http.Request {
	req, _ := http.NewRequest(method, "https://api.bigcommerce.com/"+url, body)

	req.Header.Add("X-Auth-Client", bc.AppClientID)
	req.Header.Add("X-Auth-Token", xAuthToken)
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("User-Agent", "GPMD Blog Post App")
	req.Header.Add("Cache-Control", "no-cache")
	req.Header.Add("Host", "api.bigcommerce.com")
	req.Header.Add("Accept-Encoding", "none")
	req.Header.Add("Connection", "keep-alive")
	return req
}
