package pkg

import (
	"context"
	"log"
	"net/http"
)

type Client struct {
	context    context.Context
	logger     log.Logger
	httpClient http.Client
	hostname   string
	token      string
	apiVer     string
}

func NewClient(hostname string) Client {
	return Client{
		context:    context.Background(),
		logger:     *log.Default(),
		httpClient: *http.DefaultClient,
		apiVer:     "/api/v2",
		hostname:   hostname,
	}
}
