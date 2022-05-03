# notifier

notifier cli tool that allows you to send notifications to url with intervals from stdin.

## Build executable
### Helpful tools.

Use `make`.

All the commands below must be run from the project root.

- Build executable::

        make build

It will create an executable in the root directory with name `notify`.


You can with several flags or set environment variables.

- Run executable example:

        source .env.example

        ./notify < fixtures/precepts_of_zote.txt

- or customize run:

        ./notify -i 1 -u example.com/notify < fixtures/precepts_of_zote.txt

List of flags:

| Short Flag | Long Flag | Environment Variable| Is Required | Default Value | Type | Available values |
| ---   | --- | --- | --- | --- | --- | --- |
| -u | --url | NSWC_URL | True | | string | |
| -i | --interval | NSWC_INTERVAL | False | 5 | integer | any > 0|