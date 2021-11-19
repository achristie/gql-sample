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
	CharacterStore map[string]model.Character
	OutageService  OutageService
}

type HTTPClient interface {
	Fetch(req *http.Request) (*http.Response, error)
}

type OutageService struct {
	baseUrl    *url.URL
	httpClient *http.Client
}

func NewOutageService() (*OutageService, error) {
	u, _ := url.Parse("https://api.platts.com/refinery-data/v1/outage-alerts")
	o := &OutageService{
		baseUrl:    u,
		httpClient: http.DefaultClient,
	}

	return o, nil
}

// outage.get(id=id, )

func (o *OutageService) Get() (*http.Response, error) {
	resp, err := o.httpClient.Get(o.baseUrl.String())

	if err != nil {
		return nil, err
	}
	// defer resp.Body.Close()

	return resp, nil
}
