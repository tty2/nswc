package nswc

import (
	"runtime"
	"time"
)

const defaultRequestTimeout = time.Second * 5

type config struct {
	ConnectionType TransportType
	MaxWorkers     int
	RequestTimeout time.Duration
}

type options func(*config)

// WithTransportType sets the transport type to use.
func WithTransportType(ct TransportType) options {
	return func(c *config) {
		c.ConnectionType = ct
	}
}

// WithMaxWorkers sets the maximum number of workers to use.
func WithMaxWorkers(n int) options {
	return func(c *config) {
		c.MaxWorkers = n
	}
}

// WithRequestTimeout sets the notification request timeout.
func WithRequestTimeout(d time.Duration) options {
	return func(c *config) {
		c.RequestTimeout = d
	}
}

// validate not only checks if the config is valid, but also sets default values
// if the config is invalid.
func (c *config) validate() {
	if c.MaxWorkers < 1 {
		c.MaxWorkers = runtime.NumCPU()
	}
	if c.RequestTimeout < time.Second {
		c.RequestTimeout = defaultRequestTimeout
	}
}
