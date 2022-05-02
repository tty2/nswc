package nswc

import (
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func Test_getTransportClient(t *testing.T) {
	t.Parallel()
	t.Run("ok", func(t *testing.T) {
		t.Parallel()

		tp := HTTPClient
		connStr := "http://some.url"
		tc, err := getTransportClient(tp, connStr, time.Second)

		rq := require.New(t)

		rq.NoError(err)
		rq.NotNil(tc)
	})
	t.Run("error: invalid type", func(t *testing.T) {
		t.Parallel()

		tp := TransportType(99)
		connStr := "http://some.url"
		tc, err := getTransportClient(tp, connStr, time.Second)

		rq := require.New(t)

		rq.Error(err)
		rq.Nil(tc)
	})
}
