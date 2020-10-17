package control_test

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/matryer/is"
	"github.com/vshn/control-go-sdk"
)

func TestServersListCustomerIDs(t *testing.T) {
	is := is.New(t)
	c := control.NewClientFromToken(testToken)
	numRequests := 0

	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		numRequests++
		is.Equal(r.Header.Get("X-AccessToken"), testToken) // token is set
		is.Equal(r.RequestURI, "/api/servers/1/")          // correct request URL

		fmt.Fprint(w, "foo\nbar\nbaz\n\n")
	}))
	is.NoErr(c.SetBaseURL(srv.URL))

	customers, res, err := c.Servers.ListCustomerIDs()
	is.NoErr(err)       // no error
	is.True(res != nil) // response is set
	is.Equal(customers, []string{"foo", "bar", "baz"})

	is.Equal(numRequests, 1) // got one request
}

func TestServersListFQDNs(t *testing.T) {
	is := is.New(t)
	c := control.NewClientFromToken(testToken)
	numRequests := 0

	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		numRequests++
		is.Equal(r.Header.Get("X-AccessToken"), testToken) // token is set
		is.Equal(r.RequestURI, "/api/servers/1/qq1soft/")  // correct request URL

		fmt.Fprint(w, "\ndb0.qq1soft.com\ndb1.qq1soft.com\njira.dev.qq1soft.com\n\n")
	}))
	is.NoErr(c.SetBaseURL(srv.URL))

	customers, res, err := c.Servers.ListFQDNs("qq1soft")
	is.NoErr(err)       // no error
	is.True(res != nil) // response is set
	is.Equal(customers, []string{"db0.qq1soft.com", "db1.qq1soft.com", "jira.dev.qq1soft.com"})
	is.Equal(numRequests, 1) // got one request

	// Empty customer
	srv.Config.Handler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		numRequests++
		is.Equal(r.RequestURI, "/api/servers/1/_/") // correct request URL
	})
	_, _, err = c.Servers.ListFQDNs("")
	is.NoErr(err)            // no error
	is.Equal(numRequests, 2) // got one request

}

func TestServerGetDefinition(t *testing.T) {
	is := is.New(t)
	c := control.NewClientFromToken(testToken)
	numRequests := 0

	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		numRequests++
		is.Equal(r.Header.Get("X-AccessToken"), testToken)              // token is set
		is.Equal(r.RequestURI, "/api/servers/1/_/jira.dev.qq1soft.com") // correct request URL
		w.Header().Set("Content-Type", "application/json")

		fmt.Fprint(w, `{
  "fqdn" : "jira.dev.qq1soft.com",
  "customer" : "qq1soft",
  "environment" : "QQ1Prod",
  "project" : "dev",
  "role" : "jira",
  "location" : "cloudscale",
  "stage" : "prod",
  "modDate" : 1477493084029,
  "modUser" : "qq-jdoe1"
}`)
	}))
	is.NoErr(c.SetBaseURL(srv.URL))

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
	is.Equal(numRequests, 1) // got one request
}
func TestServerGetFacts(t *testing.T) {
	is := is.New(t)
	c := control.NewClientFromToken(testToken)
	numRequests := 0

	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		numRequests++
		is.Equal(r.Header.Get("X-AccessToken"), testToken)                    // token is set
		is.Equal(r.RequestURI, "/api/servers/1/_/db1.prod.qq1soft.com/facts") // correct request URL
		w.Header().Set("Content-Type", "application/json")

		fmt.Fprint(w, `{
  "lsbdistrocodename" : "xenial",
  "processorcount" : "2"
}`)
	}))
	is.NoErr(c.SetBaseURL(srv.URL))

	s, res, err := c.Servers.GetFacts("db1.prod.qq1soft.com")
	is.NoErr(err)       // no error
	is.True(res != nil) // response is set
	is.Equal(s, map[string]string{
		"lsbdistrocodename": "xenial",
		"processorcount":    "2",
	})
	is.Equal(numRequests, 1) // got one request
}
