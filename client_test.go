package nswc

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/gammazero/workerpool"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
	"github.com/tty2/nswc/mocks"
)

func Test_Client_Notify(t *testing.T) {
	t.Parallel()

	t.Run("ok", func(t *testing.T) {
		t.Parallel()

		ctx, cancel := context.WithCancel(context.Background())
		defer cancel()

		validMsg := "valid message"
		invalidMsg := "invalid message"

		ctrl := gomock.NewController(t)
		transportMock := mocks.NewMocktransportClient(ctrl)

		transportMock.EXPECT().Notify(ctx, validMsg).Return(nil).Times(1)
		transportMock.EXPECT().Notify(ctx, invalidMsg).Return(errors.New("some error")).Times(1)

		c := Client{
			transportClient: transportMock,
			workers:         workerpool.New(2),
			errChan:         make(chan error),
		}

		rq := require.New(t)
		errs := []error{}

		go func(ctx context.Context) {
			for {
				select {
				case err := <-c.errChan:
					errs = append(errs, err)
				case <-ctx.Done():
					c.workers.Stop()
					close(c.errChan)

					return
				}
			}
		}(ctx)

		c.Send(ctx, validMsg)
		time.Sleep(time.Millisecond * 100)
		rq.Len(errs, 0)
		c.Send(ctx, invalidMsg)
		time.Sleep(time.Millisecond * 100)
		rq.Len(errs, 1)
	})
}
