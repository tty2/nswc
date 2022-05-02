// Package nswc is a simple http notification library that takes messages and notifies about them by sending HTTP POST
// requests to the configured URL with the message content in the request body.
package nswc

import (
	"github.com/gammazero/workerpool"
)

// New creates and returns a new nswc client.
// The client is configured with the connection string.
// The connection string is a valid URL that specifies the target endpoint.
// options is a variadic list of options that can be used to configure the client.
// Available options are:
// - WithTransportType sets the transport type (http, grpc, mqp etc.). Http is default (and only available for now).
// - WithMaxWorkers sets the maximum number of workers.
// - WithRequestTimeout sets the notification send timeout.
func New(connStr string, opts ...options) (*Client, error) {
	cfg := &config{}
	for _, o := range opts {
		o(cfg)
	}
	cfg.validate()

	transport, err := getTransportClient(cfg.ConnectionType, connStr, cfg.RequestTimeout)
	if err != nil {
		return nil, err
	}

	client := Client{
		transportClient: transport,
		workers:         workerpool.New(cfg.MaxWorkers),
		ErrChan:         make(chan error),
	}

	return &client, nil
}
