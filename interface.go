package smshandler

import "net/http"

// XMLResponse holds interface for XMLResponse that are implemented on each provider
type XMLResponse struct{}

// HTTPHandler holds interface for HTTP client usage that implemented by each provider
type HTTPHandler struct {
	Client *http.Client
}

// XMLHandler is an interface to send and post SMS messages
type XMLHandler interface {
	SendSMS(h HTTPHandler) (*http.Response, error)
}

// Response holds interface for
type Response interface {
	FromXMLResponse(status []byte) error
	ToError() error
	IsOK() bool
}
