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

[v] HttpTransport
[] GRPCTransport
[] RedisTransport
[] KafkaTransport
[] RabbitMQTransport
[] SNSTransport
[] SMSTransport
[] TelegramTransport