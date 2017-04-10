// server.go
package main

import (
	"fmt"
	"log"
	"net"
)

func main() {
	ln, err := net.Listen("tcp", ":6000")
	if err != nil {
		log.Fatal(err)
	}
	for {

		fmt.Println("waiting for the client")
		conn, err := ln.Accept()
		if err != nil {
			log.Println(err)
			continue
		}
		go handleConnection(conn)

	}
}
func handleConnection(c net.Conn) {
	fmt.Println("A client has connected", c.RemoteAddr())
	c.Write([]byte("Hello pakistan"))
}
