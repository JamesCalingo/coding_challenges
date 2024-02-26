package server_test

import (
	"fmt"
	"net/http"
	"net/http/httptest"
)

func testHandler(wr http.ResponseWriter, r *http.Request) {
	wr.Write([]byte("HELLO"))
}

func main() {
	req := httptest.NewRequest("GET", "localhost:8080", nil)
	recorder := httptest.NewRecorder()
	testHandler(recorder, req)
	res := recorder.Result()
	fmt.Println(res.StatusCode)
}
