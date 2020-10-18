package control

import (
	"fmt"
	"net/http"
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
	return s.client.GetStringList(serversBasePath + "/")
}

// ListFQDNs returns a list of FQDNs the provided token has access to.
// To return ALL FQDN's the token has access to, leave customerID empty or set
// it to "_"
func (s *Servers) ListFQDNs(customerID string) ([]string, *http.Response, error) {
	if customerID == "" {
		customerID = "_"
	}

	path := fmt.Sprintf("%s/%s/", serversBasePath, customerID)
	return s.client.GetStringList(path)
}

// GetDefinition returns a server definition.
func (s *Servers) GetDefinition(fqdn string) (*Server, *http.Response, error) {
	var srv Server
	res, err := s.client.GetJSON(fmt.Sprintf("%s/_/%s", serversBasePath, fqdn), &srv)
	if err != nil {
		return nil, res, err
	}
	return &srv, res, nil
}

// GetFacts returns a server definition.
func (s *Servers) GetFacts(fqdn string) (map[string]string, *http.Response, error) {
	var facts map[string]string
	res, err := s.client.GetJSON(fmt.Sprintf("%s/_/%s/facts", serversBasePath, fqdn), &facts)
	if err != nil {
		return nil, res, err
	}
	return facts, res, nil
}
