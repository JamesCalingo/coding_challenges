package server

import (
	"fmt"
	"log"
	"net/http"
)

func ServeHTML(port string) int {
	http.Handle("/q", http.FileServer(http.Dir("./public")))

	err := http.ListenAndServe(port, nil)
	if err != nil {
		fmt.Print("An error occurred")
		log.Fatal(err)
		return 1
	}
	return 0
}
