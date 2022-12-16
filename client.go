package redmineclientgo

import (
	"net/http"
	"time"
)

// HostURL - Default Redmine URL
const HostURL string = "http://localhost:3000"

type Client struct {
	HostURL    string
	HTTPClient *http.Client
	Auth       AuthStruct
}

type AuthStruct struct {
	ApiKey string `json:"api_key"`
}

func NewClient(host, apiKey *string) (*Client, error) {
	c := Client{
		HTTPClient: &http.Client{Timeout: 10 * time.Second},
		HostURL:    HostURL,
	}

	if host != nil {
		c.HostURL = *host
	}

	if apiKey == nil {
		return &c, nil
	}

	// Add the Auth struct to the client, so we can make authenticated requests
	c.Auth = AuthStruct{
		ApiKey: *apiKey,
	}

	return &c, nil
}
