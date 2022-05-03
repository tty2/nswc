package nswc

import (
	"context"

	"github.com/gammazero/workerpool"
)

type transportClient interface {
	Notify(ctx context.Context, msg string) error
}

// Client is a nswc client representing a pool of zero or more
// underlying connections.
type Client struct {
	transportClient transportClient
	workers         *workerpool.WorkerPool
	errChan         chan error
}

// Notify prepares notification to send to specified url.
func (c *Client) Send(ctx context.Context, msg string) {
	c.workers.Submit(func() {
		err := c.transportClient.Notify(ctx, msg)
		if err != nil {
			c.errChan <- err
		}
	})
}

func (c *Client) ReadErrors() <-chan error {
	return c.errChan
}

// Close stops all workers and closes ErrChan. After this the client is considered to be closed and
// no longer usable.
func (c *Client) Close() {
	c.workers.StopWait()
	close(c.errChan)
}
