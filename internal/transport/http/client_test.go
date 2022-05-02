package http

import (
	"context"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func Test_New(t *testing.T) {
	t.Parallel()

	t.Run("ok", func(t *testing.T) {
		t.Parallel()
		validURL := "http://some.url"
		c, err := New(validURL, time.Second)

		rq := require.New(t)

		rq.NoError(err)
		rq.NotNil(c)
		rq.Equal(time.Second, c.httpClient.Timeout)
		rq.Equal(validURL, c.url)
	})
	t.Run("error: invalid url", func(t *testing.T) {
		t.Parallel()
		invalidURL := "invalid url"
		c, err := New(invalidURL, time.Second)

		rq := require.New(t)

		rq.Error(err)
		rq.Nil(c)
	})
}

func TestNotify(t *testing.T) {
	t.Parallel()
	t.Run("ok", func(t *testing.T) {
		t.Parallel()
		msg := "some message"
		rq := require.New(t)
		ctx := context.Background()

		svr := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			rq.Equal(http.MethodPost, r.Method)
			defer r.Body.Close()

			body, err := io.ReadAll(r.Body)
			rq.NoError(err)
			rq.Equal(msg, string(body))
		}))
		defer svr.Close()
		c, err := New(svr.URL, time.Second)
		rq.NoError(err)

		err = c.Notify(ctx, msg)
		rq.NoError(err)
	})

	t.Run("err: timeout", func(t *testing.T) {
		t.Parallel()
		msg := "some message"
		rq := require.New(t)
		ctx := context.Background()

		svr := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			rq.Equal(http.MethodPost, r.Method)
			defer r.Body.Close()

			time.Sleep(time.Millisecond * 2)

			_, err := io.ReadAll(r.Body)
			rq.Error(err)
		}))
		defer svr.Close()
		c, err := New(svr.URL, time.Microsecond)
		rq.NoError(err)

		err = c.Notify(ctx, msg)
		rq.Error(err)
	})

	t.Run("err: status 500", func(t *testing.T) {
		t.Parallel()
		msg := "some message"
		rq := require.New(t)
		ctx := context.Background()

		svr := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			rq.Equal(http.MethodPost, r.Method)

			w.WriteHeader(http.StatusInternalServerError)
		}))
		defer svr.Close()
		c, err := New(svr.URL, time.Second)
		rq.NoError(err)

		err = c.Notify(ctx, msg)
		rq.Error(err)
	})
}
