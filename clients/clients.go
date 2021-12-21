package clients

import (
	"net"
	"net/http"
	"time"

	"github.com/iAmPlus/microservice/tracing"
)

var defaultHTTPTransport = &http.Transport{
	Proxy: http.ProxyFromEnvironment,
	DialContext: (&net.Dialer{
		Timeout:   30 * time.Second,
		KeepAlive: 30 * time.Second,
	}).DialContext,
	MaxIdleConns:          100,
	MaxIdleConnsPerHost:   100,
	IdleConnTimeout:       90 * time.Second,
	TLSHandshakeTimeout:   10 * time.Second,
	ExpectContinueTimeout: 1 * time.Second,
}

var defaultHTTPClient = &tracing.HTTPClient{Client: http.Client{Transport: defaultHTTPTransport}}

// HTTP returns the default HTTP client.
func HTTP() *tracing.HTTPClient {
	return defaultHTTPClient
}

// HTTPTransport returns default HTTP transport.
func HTTPTransport() *http.Transport {
	return defaultHTTPTransport
}
