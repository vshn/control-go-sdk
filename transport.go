package control

import "net/http"

// Transport is an implementation of http.RoundTripper that injects the control
// API access token into each request.
type Transport struct {
	token string
	T     http.RoundTripper
}

var _ http.RoundTripper = &Transport{}

// NewTransport returns a new Transport for the given token
func NewTransport(token string) *Transport {
	return &Transport{token: token, T: http.DefaultTransport}
}

// RoundTrip implements http.RoundTripper
func (t *Transport) RoundTrip(req *http.Request) (*http.Response, error) {
	r := *req
	r.Header.Set("X-AccessToken", t.token)
	return t.T.RoundTrip(&r)
}
