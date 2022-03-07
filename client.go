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

type Client struct {
	StoreHash  string `json:"store-hash"`
	XAuthToken string `json:"x-auth-token"`
	MaxRetries int
	HTTPClient *http.Client
	ChannelID  int
}

var ErrNoContent = errors.New("no content 204 from BigCommerce API")
var ErrNoMainThumbnail = errors.New("no main thumbnail")
var ErrNotFound = errors.New("404 not found")

// AuthContexter interface for GetAuthContext
type AuthContexter interface {
	GetAuthContext(clientID, clientSecret string, q url.Values) (*AuthContext, error)
}

func NewClient(storeHash, xAuthToken string) *Client {
	return &Client{
		StoreHash:  storeHash,
		XAuthToken: xAuthToken,
		MaxRetries: 1,
		HTTPClient: &http.Client{
			Timeout: time.Second * 10,
		},
		ChannelID: 1,
	}
}

func (bc *Client) getAPIRequest(method, url string, body io.Reader) *http.Request {
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
	res.Body.Close()
	if res.StatusCode > 299 {
		return body, errors.New(res.Status)
	}
	return body, nil
}
