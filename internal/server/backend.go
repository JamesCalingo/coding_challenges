package server

import (
	"fmt"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Your request: %s", r.URL.Path[1:])
}

func Serve() int {
	http.HandleFunc("/", handler)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Print("An error occurred")
		return 1
	}
	return 0
}
