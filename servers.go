package control

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

const (
	serversBasePath = "api/servers/1"
)

// ServersService is an interface for communicating with the Servers endpoint
// of the Control API.
// See: https://kb.vshn.ch/kb/api_servers.html
type ServersService interface {
	ListCustomerIDs() ([]string, *http.Response, error)
	ListFQDNs(customer string) ([]string, *http.Response, error)
	GetDefinition(fqdn string) (*Server, *http.Response, error)
	GetFacts(fqdn string) (map[string]string, *http.Response, error)
}

// Servers implements the ServersService interface
type Servers struct {
	client *Client
}

var _ ServersService = &Servers{}

// ListCustomerIDs returns a list of customer ID the provided token has access
// to.
func (s *Servers) ListCustomerIDs() ([]string, *http.Response, error) {
	res, err := s.client.Do(http.MethodGet, serversBasePath+"/", nil)
	if err != nil {
		return nil, res, err
	}

	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, res, err
	}

	ids := strings.Split(strings.TrimSpace(string(body)), "\n")
	return ids, res, nil
}

// ListFQDNs returns a list of FQDNs the provided token has access to.
// To return ALL FQDN's the token has access to, leave customerID empty or set
// it to "_"
func (s *Servers) ListFQDNs(customerID string) ([]string, *http.Response, error) {
	if customerID == "" {
		customerID = "_"
	}

	res, err := s.client.Do(http.MethodGet, fmt.Sprintf("%s/%s/", serversBasePath, customerID), nil)
	if err != nil {
		return nil, res, err
	}

	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, res, err
	}

	ids := strings.Split(strings.TrimSpace(string(body)), "\n")
	return ids, res, nil
}

// GetDefinition returns a server definition.
func (s *Servers) GetDefinition(fqdn string) (*Server, *http.Response, error) {
	res, err := s.client.Do(http.MethodGet, fmt.Sprintf("%s/_/%s", serversBasePath, fqdn), nil)
	if err != nil {
		return nil, res, err
	}

	defer res.Body.Close()
	var srv Server
	if err := json.NewDecoder(res.Body).Decode(&srv); err != nil {
		return nil, res, err
	}
	return &srv, res, nil
}

// GetFacts returns a server definition.
func (s *Servers) GetFacts(fqdn string) (map[string]string, *http.Response, error) {
	res, err := s.client.Do(http.MethodGet, fmt.Sprintf("%s/_/%s/facts", serversBasePath, fqdn), nil)
	if err != nil {
		return nil, res, err
	}

	defer res.Body.Close()
	var facts map[string]string
	if err := json.NewDecoder(res.Body).Decode(&facts); err != nil {
		return nil, res, err
	}
	return facts, res, nil
}
