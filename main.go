package main

import (
	"challenges/internal/loadtester"
	"challenges/internal/server"
	"fmt"
	"os"
)

func main() {
	fmt.Println("Up and running")
	fmt.Println(loadtester.CheckStatus(os.Args[1]))
	server.ServeHTML()
}
