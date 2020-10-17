package control_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/matryer/is"
	"github.com/vshn/control-go-sdk"
)

const (
	testToken = "helpimtrappedinatokengenerator"
)

func TestClientNewRequest(t *testing.T) {
	is := is.New(t)
	c := control.NewClient(nil)

	req, err := c.NewRequest("GET", "foo/bar", nil)
	is.NoErr(err)
	is.Equal(req.URL.String(), "https://control.vshn.net/foo/bar")            // URL is expanded
	is.Equal(req.Header.Get("User-Agent"), "control-go-sdk/"+control.Version) // User-Agent header is set
}

func TestClientSetBaseURL(t *testing.T) {
	is := is.New(t)
	c := control.NewClient(nil)
	is.Equal(c.BaseURL.String(), "https://control.vshn.net/") // Default base URL

	// Error case
	err := c.SetBaseURL(":")
	is.Equal(err.Error(), `parse ":": missing protocol scheme`)

	// OK case
	err = c.SetBaseURL("http://127.0.0.1:1234/")
	is.NoErr(err)
	is.Equal(c.BaseURL.String(), "http://127.0.0.1:1234/") // customized base url
}

func TestClientSetUserAgent(t *testing.T) {
	is := is.New(t)
	c := control.NewClient(nil)
	is.Equal(c.UserAgent, "control-go-sdk/"+control.Version) // default UA

	c.SetUserAgent("myapp/12")
	is.Equal(c.UserAgent, "myapp/12 control-go-sdk/"+control.Version) //customized UA
}

func TestClientDoHTTPError(t *testing.T) {
	is := is.New(t)
	c := control.NewClientFromToken("")

	// With error message in body
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "Omg noes auth error", 403)
	}))
	is.NoErr(c.SetBaseURL(srv.URL))
	_, err := c.Do("GET", "", nil)
	is.Equal(err.Error(), "Omg noes auth error") // returns the HTTP error message

	// No error message in body
	srv.Config.Handler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(418)
	})
	is.NoErr(c.SetBaseURL(srv.URL))
	_, err = c.Do("GET", "", nil)
	is.Equal(err.Error(), "HTTP error: 418") // returns the HTTP error message

}

func TestClientDoInvalidURL(t *testing.T) {
	is := is.New(t)
	c := control.NewClientFromToken("")
	res, err := c.Do("GET", ":bar", nil)
	is.Equal(err.Error(), `parse ":bar": missing protocol scheme`) // error from NewRequest
	is.Equal(res, nil)                                             // no response returned
}

func TestClientDoBadRequest(t *testing.T) {
	is := is.New(t)
	c := control.NewClient(&http.Client{Timeout: 1})
	is.NoErr(c.SetBaseURL("http://192.0.2.42/")) // from TEST-NET-1
	res, err := c.Do("GET", "", nil)

	is.Equal(err.Error(), `Get "http://192.0.2.42/": context deadline exceeded (Client.Timeout exceeded while awaiting headers)`) // error from http.Do
	is.Equal(res, nil)
}
