package server

import (
	"fmt"
	"net/http"
)

func ServeHTML() int {
	http.Handle("/", http.FileServer(http.Dir("./public")))
	err := http.ListenAndServe(":3000", nil)
	if err != nil {
		fmt.Print("An error occurred")
		return 1
	}
	return 0
}
