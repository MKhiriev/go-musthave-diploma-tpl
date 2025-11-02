package utils

import (
	"github.com/go-resty/resty/v2"
)

type HTTPClient struct {
	*resty.Client
}

func NewHTTPClient() *HTTPClient {
	return &HTTPClient{Client: resty.New()}
}
