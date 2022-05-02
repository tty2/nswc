package nswc

import (
	"runtime"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_config_validate(t *testing.T) {
	t.Parallel()
	t.Run("ok", func(t *testing.T) {
		t.Parallel()

		c := config{}
		c.validate()

		rq := require.New(t)

		rq.Equal(runtime.NumCPU(), c.MaxWorkers)
		rq.Equal(defaultRequestTimeout, c.RequestTimeout)
	})
}
