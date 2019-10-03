package smshandler

import "net/http"

// XMLHandler is an interface to send and post SMS messages
type XMLHandler interface {
	SendSMS(client *http.Client) (*http.Response, error)
}

// Response holds interface for
type Response interface {
	FromXMLResponse(status []byte) error
	ToError() error
	IsOK() bool
}
