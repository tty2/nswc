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
	"golang.org/x/sync/errgroup"
)

func Test_Client_Notify(t *testing.T) {
	t.Parallel()

	t.Run("ok", func(t *testing.T) {
		t.Parallel()

		ctx, cancel := context.WithTimeout(context.Background(), time.Millisecond*500)
		defer cancel()

		validMsg := "valid message"

		ctrl := gomock.NewController(t)
		transportMock := mocks.NewMocktransportClient(ctrl)

		transportMock.EXPECT().Notify(ctx, validMsg).Return(nil).Times(1)

		c := Client{
			transportClient: transportMock,
			workers:         workerpool.New(2),
			errChan:         make(chan error),
		}
		defer c.Close()

		rq := require.New(t)

		errOrNil := make(chan error)

		go func() {
			for {
				select {
				case err, ok := <-c.ReadErrors():
					if ok {
						errOrNil <- err

						return
					}

					errOrNil <- nil

					return
				case <-ctx.Done():
					errOrNil <- nil

					return
				}
			}
		}()

		c.Send(ctx, validMsg)
		rq.NoError(<-errOrNil)
	})

	t.Run("err", func(t *testing.T) {
		t.Parallel()

		ctx, cancel := context.WithTimeout(context.Background(), time.Millisecond*500)
		defer cancel()

		invalidMsg := "invalid message"

		ctrl := gomock.NewController(t)
		transportMock := mocks.NewMocktransportClient(ctrl)

		transportMock.EXPECT().Notify(ctx, invalidMsg).Return(errors.New("some error")).Times(1)

		c := Client{
			transportClient: transportMock,
			workers:         workerpool.New(2),
			errChan:         make(chan error),
		}
		defer c.Close()

		rq := require.New(t)

		g, _ := errgroup.WithContext(ctx)
		g.Go(func() error {
			for {
				select {
				case err, ok := <-c.ReadErrors():
					if ok {
						return err
					}
				case <-ctx.Done():
					return nil
				}
			}
		})

		c.Send(ctx, invalidMsg)
		err := g.Wait()
		rq.Error(err)
	})
}
