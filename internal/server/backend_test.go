package server_test

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func testHandler(wr http.ResponseWriter, r *http.Request) {
	wr.Write([]byte("HELLO"))
}

func testConnection(t *testing.T) {
	httptest.NewServer(nil)
	t.Error()
}
