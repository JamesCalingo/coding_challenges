package main

import (
	"challenges/internal/server"
	"fmt"
)

func main() {
	fmt.Println("Up and running")
	go func() {
		server.ServeHTML()
	}()
	server.Serve()
}
