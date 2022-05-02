package nswc

import (
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func Test_New(t *testing.T) {
	t.Parallel()
	t.Run("ok", func(t *testing.T) {
		t.Parallel()
		validURL := "http://some.url"
		c, err := New(validURL, WithMaxWorkers(2), WithRequestTimeout(time.Second), WithTransportType(HTTPClient))

		rq := require.New(t)

		rq.NoError(err)
		rq.NotNil(c)
		rq.False(c.workers.Stopped())

		c.Close()
		rq.True(c.workers.Stopped())
	})
}
