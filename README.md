# Go Challenge

## Overview

A server program which exposes APIs to create and get user transaction.

## Uncompleted features

- [ ] API create transaction return with http status 201.
- [ ] Authenticate user.

## API document

Check the [swagger](./docs)

## Getting Started

### Prerequisites

- [Golang](https://go.dev/doc/install)
- [Docker with Compose plugin](https://docs.docker.com/compose/install/)
- [GNU make](https://www.gnu.org/software/make/)

### Starting the Server

```sh
make compose-up
```

Note: you might need to run with `sudo` in Linux computer.

Run unit tests

```sh
make unit-test
```

To explore more useful commands for developing the project.

```sh
make help
```

### How to test

Start the local server.
Have a look at [API document](#api-document). Prepare your data to generate `curl` command. Then run the curl.
Note: please use account id from 1 to 10. I couldn't complete the API to create account. Sorry for the inconvenience.

## What's next

- [ ] Implement authentication
- [ ] API to create account by user
- [ ] Improve unit test coverage
- [ ] Integration tests
