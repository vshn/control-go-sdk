package control_test

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/matryer/is"
)

type mockAPI struct {
	*httptest.Server
	is *is.I

	NumRequests int

	ExpectedAccessToken string
	ExpectedRequestURI  string

	ResponseBody        string
	ResponseCode        int
	ResponseContentType string
}

func newMockAPI(t *testing.T, expectedURI string) *mockAPI {
	srv := &mockAPI{
		is:                  is.New(t),
		NumRequests:         0,
		ExpectedAccessToken: testToken,
		ExpectedRequestURI:  expectedURI,
		ResponseBody:        "",
		ResponseCode:        200,
		ResponseContentType: "",
	}
	srv.Server = httptest.NewServer(srv)

	return srv
}

func (m *mockAPI) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	m.NumRequests++
	m.is.Equal(r.Header.Get("X-AccessToken"), m.ExpectedAccessToken)
	m.is.Equal(r.RequestURI, m.ExpectedRequestURI)

	if m.ResponseContentType != "" {
		w.Header().Set("Content-Type", m.ResponseContentType)
	}
	w.WriteHeader(m.ResponseCode)
	if m.ResponseBody != "" {
		fmt.Fprint(w, m.ResponseBody)
	}
}
