package server

import (
	"fmt"
	"log"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, r.Proto)
}

func Serve(port string) int {
	http.HandleFunc("/", handler)
	err := http.ListenAndServe(port, nil)
	if err != nil {
		fmt.Print("An error occurred")
		log.Fatal(err)
		return 1
	}
	return 0
}
