package bigcommerce

import (
	"errors"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"time"
)

// BigCommerce is the BigCommerce API client object for BigCommerce Apps
// holds no client specific information
type BigCommerceApp struct {
	Hostname        string
	AppClientID     string
	AppClientSecret string
	HTTPClient      *http.Client
	MaxRetries      int
}

type BigCommerce struct {
	StoreHash  string `json:"store-hash"`
	XAuthToken string `json:"x-auth-token"`
	MaxRetries int
	HTTPClient *http.Client
}

var ErrNoContent = errors.New("no content 204 from BigCommerce API")
var ErrNoMainThumbnail = errors.New("no main thumbnail")
var ErrNotFound = errors.New("404 not found")

// AuthContexter interface for GetAuthContext
type AuthContexter interface {
	GetAuthContext(clientID, clientSecret string, q url.Values) (*AuthContext, error)
}

// New returns a new BigCommerce API object with the given hostname, client ID, and client secret
// The client ID and secret are the App's client ID and secret from the BigCommerce My Apps dashboard
// The hostname is the domain name of the app from the same page (e.g. app.exampledomain.com)
func NewApp(hostname, appClientID, appClientSecret string) *BigCommerceApp {
	return &BigCommerceApp{
		Hostname:        hostname,
		AppClientID:     appClientID,
		AppClientSecret: appClientSecret,
		MaxRetries:      1,
		HTTPClient: &http.Client{
			Timeout: time.Second * 10,
		},
	}
}

func NewClient(storeHash, xAuthToken string) *BigCommerce {
	return &BigCommerce{
		StoreHash:  storeHash,
		XAuthToken: xAuthToken,
		MaxRetries: 1,
		HTTPClient: &http.Client{
			Timeout: time.Second * 10,
		},
	}
}

func (bc *BigCommerce) getAPIRequest(method, url string, body io.Reader) *http.Request {
	if !strings.HasPrefix(url, "/") {
		url = "/" + url
	}
	fullURL := "https://api.bigcommerce.com/stores/" + bc.StoreHash + url

	req, _ := http.NewRequest(method, fullURL, body)

	req.Header.Add("X-Auth-Token", bc.XAuthToken)
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("User-Agent", "BigCommerce-Go-SDK")
	req.Header.Add("Cache-Control", "no-cache")
	req.Header.Add("Host", "api.bigcommerce.com")
	req.Header.Add("Accept-Encoding", "none")
	req.Header.Add("Connection", "keep-alive")
	return req
}

func processBody(res *http.Response) ([]byte, error) {
	if res.StatusCode == http.StatusNoContent {
		return nil, ErrNoContent
	}
	if res.StatusCode == http.StatusNotFound {
		return nil, ErrNotFound
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	return body, nil
}
