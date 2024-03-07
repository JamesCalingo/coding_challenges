package server

import (
	"fmt"
	"log"
	"net"
	"net/url"
)

func handleConnection(c net.Conn) {
	fmt.Println("Incoming request from", c.RemoteAddr())
	_, err := c.Write([]byte("200"))
	if err != nil {
		fmt.Println(err)
	}
	defer c.Close()
}

func Connect(input string) {
	netUrl, netErr := url.Parse(input)
	if netErr != nil {
		log.Fatal(netErr)
	}
	fmt.Println("Scheme:", netUrl.Scheme, "Opaque:", netUrl.Opaque, "Host:", netUrl.Host, "Path:", netUrl.Path)
	ln, lnErr := net.Listen("tcp", input)
	if lnErr != nil {
		fmt.Println("Could not establish connection")
		log.Fatal(lnErr)
	}
	fmt.Println("Listening @", ln.Addr())
	for {
		conn, connErr := ln.Accept()
		if connErr != nil {
			fmt.Println("Error connecting")
			ln.Close()
			log.Fatal(connErr)
		}
		go handleConnection(conn)
	}

}
