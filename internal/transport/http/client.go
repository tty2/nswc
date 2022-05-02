// Package http provides http client.
// This http client has limited set of methods that are needed for nswc client only.
package http

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"strings"
	"time"
)

type client struct {
	url        string
	httpClient *http.Client
}

// New creates a new http client.
func New(connStr string, timeout time.Duration) (*client, error) {
	_, err := url.ParseRequestURI(connStr)
	if err != nil {
		return nil, fmt.Errorf("invalid connection string: %v", err)
	}

	c := client{
		url: connStr,
		httpClient: &http.Client{
			Timeout: timeout,
		},
	}

	return &c, nil
}

// Notify sends notification for specified url.
func (c *client) Notify(ctx context.Context, msg string) error {
	r := strings.NewReader(msg)

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, c.url, r)
	if err != nil {
		return fmt.Errorf("can't create a new request: %v", err)
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return fmt.Errorf("can't make the request for message `%s`: %v", msg, err)
	}
	defer func() {
		closeErr := resp.Body.Close()
		if closeErr != nil && err == nil {
			err = closeErr
		}
	}()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("http status code is not ok: %d, message: %s", resp.StatusCode, msg)
	}

	return nil
}
