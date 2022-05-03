package main

import (
	"bufio"
	"context"
	"errors"
	"fmt"
	"io"
	"log"
	"os"
	"os/signal"
	"time"

	"github.com/tty2/nswc"
	"github.com/urfave/cli/v2"
)

type notifier interface {
	Send(ctx context.Context, msg string)
	ReadErrors() <-chan error
}

// nolint forbidigo: printf here is on purpose
func main() {
	app := &cli.App{
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:     "url",
				Aliases:  []string{"u"},
				EnvVars:  []string{"NSWC_URL"},
				Required: true,
				Usage:    "notification url",
			},
			&cli.IntFlag{
				Name:    "interval",
				Aliases: []string{"i"},
				Value:   5,
				EnvVars: []string{"NSWC_INTERVAL"},
				Usage:   "send interval",
			},
		},
		Action: func(c *cli.Context) error {
			url := c.String("url")
			intervalValue := c.Int("interval")
			if intervalValue < 1 {
				return errors.New("interval must be greater than 0")
			}
			interval := time.Duration(intervalValue) * time.Second

			ctx, cancel := signal.NotifyContext(c.Context, os.Interrupt)
			defer cancel()

			reader := bufio.NewReader(os.Stdin)

			client, err := nswc.New(url)
			if err != nil {
				return err
			}
			defer client.Close()

			return notify(ctx, client, reader, interval)
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}

// notify reads stdin line by line and sends it to the notification url
func notify(ctx context.Context, client notifier, reader *bufio.Reader, interval time.Duration) error {
	ticker := time.NewTicker(interval)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			line, err := reader.ReadString('\n')
			if err != nil {
				if errors.Is(err, io.EOF) {
					fmt.Println("file reading completed")
					return nil
				}
				return err
			}
			client.Send(ctx, line)
		case <-ctx.Done():
			return errors.New("interrupted")
		case err, ok := <-client.ReadErrors():
			if !ok {
				return errors.New("error channel has been closed")
			}
			return fmt.Errorf("nswc library error: %v", err)
		}
	}
}
