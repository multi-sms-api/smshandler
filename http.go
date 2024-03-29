package smshandler

import (
	"bytes"
	"context"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httputil"
	"net/url"
	"os"
	"strings"
)

// DoHTTP sends an HTTP Request for sending an SMS
func DoHTTP(client *http.Client, method, contentType, address string,
	fields url.Values, body []byte, onResponse Response) (resp *http.Response, err error) {

	var request *http.Request
	var bodyReader *bytes.Reader

	fullAddress := address

	if body != nil {
		bodyReader = bytes.NewReader(body)
	} else {
		bodyReader = bytes.NewReader([]byte{})
	}

	switch method {
	case http.MethodGet:
		fullAddress = fmt.Sprintf("%s?%s", fullAddress, fields.Encode())
	}
	if strings.Contains(os.Getenv("SMSHTTPDEBUG"), "dump=true") {
		fmt.Printf(">>>> method: %s, fullAddress: %s, bodyReader.len %d, contentType: %s\n",
			method, fullAddress, bodyReader.Len(), contentType)
	}

	request, err = http.NewRequest(method, fullAddress, bodyReader)

	if err != nil {
		return nil, err
	}

	if contentType != "" {
		request.Header.Set("Content-Type", contentType)
	}
	request.Close = true

	ctx, cancel := context.WithTimeout(request.Context(), client.Timeout)
	defer cancel()
	defer client.CloseIdleConnections()

	resp, err = client.Do(request.WithContext(ctx))

	if err != nil {
		if strings.Contains(os.Getenv("SMSHTTPDEBUG"), "dump=true") {
			fmt.Printf(">>>> Error was given: %s", err)
		}
		return
	}

	if resp == nil {
		err = fmt.Errorf("resp is nil")
		if strings.Contains(os.Getenv("SMSHTTPDEBUG"), "dump=true") {
			fmt.Println(">>>> resp is nil")
		}
		return
	}

	if strings.Contains(os.Getenv("SMSHTTPDEBUG"), "dump=true") {
		dump, err := httputil.DumpRequestOut(request, true)
		fmt.Printf(">>>> dump request: %s \nerr: %s\n", dump, err)

		dump, err = httputil.DumpResponse(resp, true)
		fmt.Printf(">>>> dump response: %s \nerr: %s\n", dump, err)
	}

	if resp.Body == nil {
		err = fmt.Errorf(">>>> resp.body is nil")
		return
	}

	var respBody []byte
	respBody, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}

	if onResponse == nil {
		return
	}

	if resp.StatusCode == http.StatusOK {
		err = onResponse.FromResponse(respBody)
		if err != nil {
			return
		}
		if !onResponse.IsOK() {
			err = onResponse.ToError()
		}
	}

	return
}

// OnGettingSMS is an HTTP server handler when incoming SMS arrives.
// If mux exists, it will use it for a server, otherwise it will
// use http.HandleFunc.
func OnGettingSMS(path string, mux *http.ServeMux, httpHandler http.HandlerFunc) {
	if mux != nil {
		mux.HandleFunc(path, httpHandler)
		return
	}

	http.HandleFunc(path, httpHandler)
}
