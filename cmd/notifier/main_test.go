package main

import (
	"bufio"
	"bytes"
	"context"
	"errors"
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
	"github.com/tty2/nswc/cmd/notifier/mocks"
)

func Test_notify(t *testing.T) {
	t.Parallel()
	t.Run("ok", func(t *testing.T) {
		t.Parallel()
		ctx := context.Background()
		rq := require.New(t)

		ctrl := gomock.NewController(t)
		notifierMock := mocks.NewMocknotifier(ctrl)

		errChan := make(chan error)
		defer close(errChan)

		notifierMock.EXPECT().Send(ctx, "message1\n").Return().Times(1)
		notifierMock.EXPECT().Send(ctx, "message2\n").Return().Times(1)
		notifierMock.EXPECT().ReadErrors().Return(errChan).AnyTimes()

		reader := bufio.NewReader(bytes.NewReader(
			[]byte("message1\nmessage2\n"),
		))

		err := notify(ctx, notifierMock, reader, time.Millisecond)
		rq.NoError(err)
	})
	t.Run("err: interrupted", func(t *testing.T) {
		t.Parallel()
		ctx, cancel := context.WithTimeout(context.Background(), time.Second)
		defer cancel()
		rq := require.New(t)

		ctrl := gomock.NewController(t)
		notifierMock := mocks.NewMocknotifier(ctrl)

		errChan := make(chan error)
		defer close(errChan)

		notifierMock.EXPECT().ReadErrors().Return(errChan).AnyTimes()

		reader := bufio.NewReader(bytes.NewReader(
			[]byte("message1\nmessage2\n"),
		))

		// make ticker > than context deadline time
		err := notify(ctx, notifierMock, reader, 10*time.Second)
		rq.Error(err)
	})
	t.Run("err: sent by sender", func(t *testing.T) {
		t.Parallel()
		ctx := context.Background()
		rq := require.New(t)

		ctrl := gomock.NewController(t)
		notifierMock := mocks.NewMocknotifier(ctrl)

		errChan := make(chan error)
		defer close(errChan)

		notifierMock.EXPECT().ReadErrors().Return(errChan).AnyTimes()

		reader := bufio.NewReader(bytes.NewReader(
			[]byte("message1\nmessage2\n"),
		))

		go func() {
			errChan <- errors.New("some error")
		}()

		// make ticker > than context deadline time
		err := notify(ctx, notifierMock, reader, 10*time.Second)
		rq.Error(err)
	})
	t.Run("err: errChan has been closed", func(t *testing.T) {
		t.Parallel()
		ctx := context.Background()
		rq := require.New(t)

		ctrl := gomock.NewController(t)
		notifierMock := mocks.NewMocknotifier(ctrl)

		errChan := make(chan error)

		notifierMock.EXPECT().ReadErrors().Return(errChan).AnyTimes()

		reader := bufio.NewReader(bytes.NewReader(
			[]byte("message1\nmessage2\n"),
		))

		go func() {
			time.Sleep(time.Second)
			close(errChan)
		}()

		err := notify(ctx, notifierMock, reader, 10*time.Second)
		rq.Error(err)
	})
}
