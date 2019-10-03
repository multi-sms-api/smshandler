package smshandler

import "net/http"

// XMLResponse holds interface for XMLResponse that are implemented on each provider
type XMLResponse struct{}

// HTTPHandler holds interface for HTTP client usage that implemented by each provider
type HTTPHandler struct{}

// XMLHandler is an interface to send and post SMS messages
type XMLHandler interface {
	SendSMS(h HTTPHandler) (*http.Response, error)
}

// Response holds interface for
type Response interface {
	FromXMLResponse(status []byte) (XMLResponse, error)
	ToError(status XMLResponse) error
	IsOK(response XMLResponse) bool
}
