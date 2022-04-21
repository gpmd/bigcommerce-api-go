package bigcommerce

import (
	"io"
	"net/http"
	"time"
)

type HTTPClient interface {
	Do(req *http.Request) (res *http.Response, err error)
	Get(url string) (res *http.Response, err error)
	Post(urstring, bodyType string, body io.Reader) (res *http.Response, err error)
}

// BigCommerce is the BigCommerce API client object for BigCommerce Apps
// holds no client specific information
type App struct {
	Hostname        string
	AppClientID     string
	AppClientSecret string
	HTTPClient      HTTPClient
	MaxRetries      int
	ChannelID       int
}

// New returns a new BigCommerce API object with the given hostname, client ID, and client secret
// The client ID and secret are the App's client ID and secret from the BigCommerce My Apps dashboard
// The hostname is the domain name of the app from the same page (e.g. app.exampledomain.com)
func NewApp(hostname, appClientID, appClientSecret string) *App {
	return &App{
		Hostname:        hostname,
		AppClientID:     appClientID,
		AppClientSecret: appClientSecret,
		MaxRetries:      1,
		HTTPClient: &http.Client{
			Timeout: time.Second * 10,
		},
	}
}

func (a *App) NewClient(storeHash, xAuthToken string) *Client {
	return &Client{
		StoreHash:  storeHash,
		XAuthToken: xAuthToken,
		MaxRetries: 1,
		HTTPClient: a.HTTPClient,
		ChannelID:  1,
	}
}
