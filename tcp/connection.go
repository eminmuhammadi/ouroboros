package tcp

import (
	"log"
	"net"
	"time"
)

type Client struct {
	Connection net.Conn
	Time       time.Time
}

// Accepts a client connection
func CreateChannel(listener net.Listener,
	requestHandler RequestHandler,
	responseHandler ResponseHandler) {
	for {
		connection, err := listener.Accept()

		// Handle error from accepted connection
		if err != nil {
			log.Println(err.Error()) // Error handling

			connection.Close()
			listener.Close()
			break
		}

		// Create client
		client := Client{
			Connection: connection,
			Time:       time.Now(),
		}

		// Handler
		go client.Handler(requestHandler, responseHandler)
	}
}
