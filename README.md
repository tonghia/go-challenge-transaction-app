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

To explore more useful commands for developing the project.

```sh
make help
```

## What's next

- [ ] Implement authentication
- [ ] Improve unit test coverage
- [ ] Integration tests
