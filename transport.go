package nswc

import (
	"errors"
	"time"

	"github.com/tty2/nswc/internal/transport/http"
)

// TransportType is a type of transport.
// HTTP client, GRPC client, Message Queue client, etc.
// HTTPClient is the default transport type and only available for now.
type TransportType int

// Transport types enum.
const (
	HTTPClient TransportType = iota
	// GRPCClient
	// MQClient
	// ...etc.
)

func getTransportClient(tp TransportType, connStr string, timeout time.Duration) (transportClient, error) {
	switch tp {
	case HTTPClient:
		return http.New(connStr, timeout)
	default:
		return nil, errors.New("undefined transport type")
	}
}
