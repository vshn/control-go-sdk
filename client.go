package control

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"time"
)

const (
	defaultBaseURL = "https://control.vshn.net/"
	userAgent      = "control-go-sdk/" + Version
)

// Client for interacting with the VSHN Control API
type Client struct {
	// Base URL for API requests
	BaseURL *url.URL

	// User Agent string for client
	UserAgent string

	// HTTP client used to communicate
	client *http.Client

	// Services used for communicating with the APIs
	Servers ServersService
}

// NewClient returns a new Client with the given HTTP client in place
func NewClient(httpClient *http.Client) *Client {
	if httpClient == nil {
		httpClient = http.DefaultClient
	}

	baseURL, _ := url.Parse(defaultBaseURL)

	c := &Client{client: httpClient, BaseURL: baseURL, UserAgent: userAgent}
	c.Servers = &Servers{client: c}

	return c
}

// NewClientFromToken constructs a new client with a HTTP client that will
// inject the access token into each request
func NewClientFromToken(token string) *Client {
	httpClient := &http.Client{
		Transport: NewTransport(token),
		Timeout:   10 * time.Second,
	}
	return NewClient(httpClient)
}

// NewRequest prepares a new HTTP request with its URL properly expanded (using
// the clients base URL) and the User-Agent header set.
func (c *Client) NewRequest(method, url string, body io.Reader) (*http.Request, error) {
	u, err := c.BaseURL.Parse(url)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest(method, u.String(), body)
	if err != nil {
		return nil, err
	}

	req.Header.Set("User-Agent", c.UserAgent)

	return req, nil
}

// Do the given HTTP request
func (c *Client) Do(method, url string, body io.Reader) (*http.Response, error) {
	req, err := c.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}

	res, err := c.client.Do(req)
	if err != nil {
		return nil, err
	}

	if err := checkResponse(res); err != nil {
		return res, err
	}

	return res, nil
}

// SetBaseURL customizes the default base URL used by the client
func (c *Client) SetBaseURL(newURL string) error {
	u, err := url.Parse(newURL)
	if err != nil {
		return err
	}

	c.BaseURL = u
	return nil
}

// SetUserAgent sets an (optional) additional user agent string. Format will
// always be `<your UA> <default UA>`
func (c *Client) SetUserAgent(ua string) {
	c.UserAgent = ua + " " + userAgent
}

func checkResponse(res *http.Response) error {
	code := res.StatusCode
	if code >= 200 && code <= 299 {
		return nil
	}
	body, err := ioutil.ReadAll(res.Body)
	if err == nil && len(body) > 0 {
		return errors.New(strings.TrimSpace(string(body)))
	}

	return fmt.Errorf("HTTP error: %d", code)
}

// GetStringList does a HTTP GET request to the API with the given path, and
// returns the response as a slice of strings.
func (c *Client) GetStringList(path string) ([]string, *http.Response, error) {
	res, err := c.Do(http.MethodGet, path, nil)
	if err != nil {
		return nil, res, err
	}

	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, res, err
	}

	body = bytes.TrimSpace(body)
	lines := strings.Split(string(body), "\n")

	return lines, res, nil
}

// GetJSON sends a HTTP GET request to the API and decodes the (JSON) response
// body into the passed struct.
func (c *Client) GetJSON(path string, v interface{}) (*http.Response, error) {
	res, err := c.Do(http.MethodGet, path, nil)
	if err != nil {
		return res, err
	}

	defer res.Body.Close()
	if err := json.NewDecoder(res.Body).Decode(&v); err != nil {
		return res, err
	}
	return res, nil
}
