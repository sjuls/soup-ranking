package utils

import "net/http"

type (
	// HTTPClient defines a means to invoke http requests
	HTTPClient interface {
		Do(req *http.Request) (*http.Response, error)
	}
)
