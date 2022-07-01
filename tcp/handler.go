package tcp

import (
	"bufio"
	"log"
	"net"
)

type RequestHandler func(string, net.Conn) error
type ResponseHandler func(string, net.Conn) error

// Handles a request from a client
func (client *Client) Handler(requestHandler RequestHandler, responseHandler ResponseHandler) {
	reader := bufio.NewReader(client.Connection)

	for {
		data, err := reader.ReadString('\n')

		// Read stream
		if err != nil {
			log.Println(err.Error()) // Error handling

			client.Connection.Close()
			break
		}

		// Handle request
		if err := client.Request(data, requestHandler); err != nil {
			log.Println(err.Error()) // Error handling

			client.Connection.Close()
			break
		}

		// Handle response
		if err := client.Response(data, responseHandler); err != nil {
			log.Println(err.Error()) // Error handling

			client.Connection.Close()
			break
		}
	}
}

// Request a message from a client
func (client *Client) Request(data string, requestHandler RequestHandler) error {
	return requestHandler(data, client.Connection)
}

// Response a message to a client
func (client *Client) Response(data string, responseHandler ResponseHandler) error {
	return responseHandler(data, client.Connection)
}
