package smshandler

import (
	"net/http"
	"net/url"
)

// XMLHandler is an interface to send and post SMS messages
type XMLHandler interface {
	SendSMS(client *http.Client) (*http.Response, error)
}

// RESTHandler is an interface to send and post SMS messages using HTTP REST
type RESTHandler interface {
	SendSMS(method string, client *http.Client, onResponse Response) (*http.Response, error)
}

// HTTPHandler interface for working with HTTP client interface
type HTTPHandler interface {
	DoHTTP(method, contentType, address string, fields url.Values, body []byte) (resp *http.Response, err error)
	OnGettingSMS(path string, mux *http.ServeMux, httpHandler http.HandlerFunc)
}

// Response holds interface for
type Response interface {
	FromXMLResponse(status []byte) error
	FromJSONResponse(status []byte) error
	ToError() error
	IsOK() bool
}
