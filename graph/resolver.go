package graph

import (
	"net/http"
	"net/url"

	"github.com/achristie/gql-sample/graph/model"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	CharacterStore   map[string]model.Character
	PlattsApiService HTTPClient
}

type HTTPClient interface {
	Do(req *http.Request) (*http.Response, error)
}

type Client struct {
	baseURL    *url.URL
	httpClient HTTPClient
	apiKey     string
}

func NewClient(baseURL string, apiKey string) (*Client, error) {
	url, err := url.Parse(baseURL)

	if err != nil {
		return &Client{}, err
	}
	c := &Client{
		baseURL:    url,
		httpClient: http.DefaultClient,
		apiKey:     apiKey,
	}
	return c, nil
}

func (c *Client) Get() (resp *http.Response, err error) {
	req, err := http.NewRequest("GET", c.baseURL.String(), nil)

	if err != nil {
		return nil, err
	}

	return c.Do(req)
}

func (c *Client) Do(req *http.Request) (*http.Response, error) {
	req.Header.Add("appkey", c.apiKey)
	return c.httpClient.Do(req)
}
