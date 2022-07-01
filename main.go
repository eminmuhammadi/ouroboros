package main

import (
	"log"
	"net"

	tcp "github.com/eminmuhammadi/ouroboros/tcp"
)

/**
 * Handles a request from a client
 */
func requestHandler(data string, conn net.Conn) error {
	log.Printf("%s: %s", conn.RemoteAddr().String(), data)

	return nil
}

/**
 * Handles a response from a client
 */
func responseHandler(data string, conn net.Conn) error {
	conn.Write([]byte("OK\n"))

	return nil
}

/**
 * Main function
 */
func main() {
	listener, err := tcp.CreateInsecureListener(tcp.Endpoint{
		Host: "127.0.0.1",
		Port: "8080",
	})

	if err != nil {
		panic(err)
	}

	tcp.CreateChannel(listener, requestHandler, responseHandler)
}
