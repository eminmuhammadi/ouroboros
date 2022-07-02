package main

import (
	"fmt"
	"net"
	"time"

	tcp "github.com/eminmuhammadi/ouroboros/tcp"
)

/**
 * Handles a request from a client
 */
func requestHandler(data string, conn net.Conn) error {
	host, port, error := net.SplitHostPort(conn.RemoteAddr().String())

	if error != nil {
		return error
	}

	// Send a response to the server
	println(fmt.Sprintf("%s\t%s\t%s", fmt.Sprint(time.Now().Unix()), host, port))

	return nil
}

/**
 * Handles a response to client
 */
func responseHandler(data string, conn net.Conn) error {
	return nil
}

/**
 * Main function
 */
func main() {
	// secureListener, err := tcp.CreateSecureListener(tcp.SecureEndpoint{
	// 	Host: "127.0.0.1",
	// 	Port: "8080",
	// 	Certificate: "",
	// 	Key: "",
	// 	Config: &tcp.TLSConfig{},
	// })

	listener, err := tcp.CreateInsecureListener(tcp.Endpoint{
		Host: "127.0.0.1",
		Port: "8080",
	})

	if err != nil {
		panic(err)
	}

	tcp.CreateChannel(listener, requestHandler, responseHandler)
}
