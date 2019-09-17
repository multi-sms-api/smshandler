package smshandler

import (
	"net/http"
	"net/url"
)

// HTTPHandler is an interface to send and post SMS messages
type HTTPHandler interface {
	SendSMS(method, contentType, address string, fields url.Values, body []byte) (resp *http.Response, err error)
	OnGettingSMS(server http.Server, path string, httpHandler http.HandlerFunc)
}
