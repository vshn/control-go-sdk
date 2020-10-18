package control_test

import (
	"testing"

	"github.com/matryer/is"
	"github.com/vshn/control-go-sdk"
)

func TestServersListCustomerIDs(t *testing.T) {
	is := is.New(t)
	c := control.NewClientFromToken(testToken)
	api := newMockAPI(t, "/api/servers/1/")
	api.ResponseBody = "foo\nbar\nbaz\n\n"
	is.NoErr(c.SetBaseURL(api.URL))

	customers, res, err := c.Servers.ListCustomerIDs()
	is.NoErr(err)       // no error
	is.True(res != nil) // response is set
	is.Equal(customers, []string{"foo", "bar", "baz"})

	is.Equal(api.NumRequests, 1) // got one request
}

func TestServersListFQDNs(t *testing.T) {
	is := is.New(t)
	c := control.NewClientFromToken(testToken)
	api := newMockAPI(t, "/api/servers/1/qq1soft/")
	api.ResponseBody = "\ndb0.qq1soft.com\ndb1.qq1soft.com\njira.dev.qq1soft.com\n\n"

	is.NoErr(c.SetBaseURL(api.URL))

	customers, res, err := c.Servers.ListFQDNs("qq1soft")
	is.NoErr(err)       // no error
	is.True(res != nil) // response is set
	is.Equal(customers, []string{"db0.qq1soft.com", "db1.qq1soft.com", "jira.dev.qq1soft.com"})
	is.Equal(api.NumRequests, 1) // got one request

	// Empty customer
	api.ExpectedRequestURI = "/api/servers/1/_/"
	_, _, err = c.Servers.ListFQDNs("")
	is.NoErr(err)                // no error
	is.Equal(api.NumRequests, 2) // got another request
}

func TestServerGetDefinition(t *testing.T) {
	is := is.New(t)
	c := control.NewClientFromToken(testToken)
	api := newMockAPI(t, "/api/servers/1/_/jira.dev.qq1soft.com")
	api.ResponseContentType = "application/json"
	api.ResponseBody = `{
  "fqdn" : "jira.dev.qq1soft.com",
  "customer" : "qq1soft",
  "environment" : "QQ1Prod",
  "project" : "dev",
  "role" : "jira",
  "location" : "cloudscale",
  "stage" : "prod",
  "modDate" : 1477493084029,
  "modUser" : "qq-jdoe1"
}`

	is.NoErr(c.SetBaseURL(api.URL))

	s, res, err := c.Servers.GetDefinition("jira.dev.qq1soft.com")
	is.NoErr(err)       // no error
	is.True(res != nil) // response is set
	is.Equal(s, &control.Server{
		FQDN:        "jira.dev.qq1soft.com",
		Customer:    "qq1soft",
		Environment: "QQ1Prod",
		Project:     "dev",
		Role:        "jira",
		Location:    "cloudscale",
		Stage:       "prod",
		ModDate:     1477493084029,
		ModUser:     "qq-jdoe1",
	})
	is.Equal(api.NumRequests, 1) // got one request
}
func TestServerGetFacts(t *testing.T) {
	is := is.New(t)
	c := control.NewClientFromToken(testToken)
	api := newMockAPI(t, "/api/servers/1/_/db1.prod.qq1soft.com/facts")
	api.ResponseContentType = "application/json"
	api.ResponseBody = `{
  "lsbdistrocodename" : "xenial",
  "processorcount" : "2"
}`

	is.NoErr(c.SetBaseURL(api.URL))

	s, res, err := c.Servers.GetFacts("db1.prod.qq1soft.com")
	is.NoErr(err)       // no error
	is.True(res != nil) // response is set
	is.Equal(s, map[string]string{
		"lsbdistrocodename": "xenial",
		"processorcount":    "2",
	})
	is.Equal(api.NumRequests, 1) // got one request
}
