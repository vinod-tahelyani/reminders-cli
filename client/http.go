package client

import (
	"net/http"
	"time"
)

type HTTPClient struct {
	client *http.Client
	BackendURI string
}

func NewHTTPClient(uri string) HTTPClient {
	return HTTPClient{
		BackendURI: uri,
		client: &http.Client{},
	}
}

func (client HTTPClient) Create(title string, message string, duration time.Duration) ([]byte, error) {
	return []byte{}, nil
}

func (client HTTPClient) Edit(id int, title string, message string, duration time.Duration) ([]byte, error) {
	return []byte{}, nil
}

func (client HTTPClient) Fetch(id int) ([]byte, error) {
	return []byte{}, nil
}

func (client HTTPClient) Delete(id int) (error) {
	return nil
}

func (client HTTPClient) Healthy(host string) bool {
	return true
}