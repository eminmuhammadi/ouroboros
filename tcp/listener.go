package tcp

import (
	"crypto/tls"
	"net"
)

type Endpoint struct {
	Host string
	Port string
}

type SecureEndpoint struct {
	Endpoint
	Certificate string
	Key         string
	Config      *tls.Config
}

// Create insecure listener on given endpoint
func CreateInsecureListener(endpoint Endpoint) (net.Listener, error) {
	// Create listener
	listener, err := net.Listen("tcp", net.JoinHostPort(endpoint.Host, endpoint.Port))

	return listener, err
}

// Create secure listener on given endpoint
func CreateSecureListener(endpoint SecureEndpoint) (net.Listener, error) {
	// Load certificate
	cert, err := tls.LoadX509KeyPair(endpoint.Certificate, endpoint.Key)
	if err != nil {
		return nil, err
	}

	// Create TLS config
	endpoint.Config.Certificates = []tls.Certificate{cert}

	// Create listener
	listener, err := tls.Listen("tcp", net.JoinHostPort(endpoint.Host, endpoint.Port), endpoint.Config)

	return listener, err
}
