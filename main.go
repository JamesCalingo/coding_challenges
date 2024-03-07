package main

import (
	"challenges/internal/server"
	"fmt"
	"os"
)

func main() {
	port := os.Args[1]
	fmt.Println("Up and running")
	server.Connect(port)
	// go func() {
	// 	server.ServeHTML(":5173")
	// }()
	// server.Serve(port)
}
