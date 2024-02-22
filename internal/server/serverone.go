package server

import (
	"fmt"
	"net/http"
)

func handlerone(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Your link is here: %s", r.URL.Path)
}

func ServeBasic() int {
	http.HandleFunc("/", handlerone)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Print("An error occurred")
		return 1
	}
	return 0
}