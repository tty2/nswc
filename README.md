# nswc


nswc is a simple client that provides you to send notifications to set address by several transports.

## Installation

        go get github.com/tty2/nswc

## Quickstart

        import (
                "context"
                "github.com/tty2/nswc"
        )

        func main() {
                nswcClient := nswc.New("http://someurl.com/notify")

                ctx := context.Background()
                nswcClient.Send(ctx, "Hello World!")
                nswc.Close()
        }
                       

## Development

### Helpful tools.

Use `make` to lint, make mocks or test project.

- Show available make commands:

        make

- Lint project:

        make lint

- Test project:

        make test

- Generate mocks:

        make mocks

## Road map

- [x] HttpTransport
- [ ] GRPCTransport
- [ ] RedisTransport
- [ ] KafkaTransport
- [ ] RabbitMQTransport
- [ ] SNSTransport
- [ ] SMSTransport
- [ ] TelegramTransport

## Playground

In order to play with the library you can use `./cmd/notifier` and `./cmd/listener`.

Check  [notifier readme](./cmd/notifier/README.md)
and  [listener readme](./cmd/listener/README.md) for more information.