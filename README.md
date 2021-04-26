[![Build](https://github.com/mohsanabbas/cart-microservice/actions/workflows/build.yml/badge.svg)](https://github.com/mohsanabbas/cart-microservice/actions/workflows/build.yml)
[![Go Report Card](https://goreportcard.com/badge/github.com/mohsanabbas/cart-microservice)](https://goreportcard.com/report/github.com/mohsanabbas/cart-microservice)

# SHOPPING-CART MICROSERVICE IN GO

Shopping cart microservice written in [GO](https://golang.org/) to perform CRUD operations.

## .env examle

```bash
export DB_DRIVER=
export DB_NAME=shopping-cart
export DB_SOURCE=
export DB_USER=
export DB_PASSWORD=
export COLLECTION=
export SERVER_ADDRESS=
export VAULT_SCHEME=
export VAULT_HOST=
export VAULT_PORT=
export VAULT_SECURITY_TOKEN=
export CONSUL_HOST=
export CONSUL_PORT=
```

## Stacks

- [GO](https://golang.org/) - The Go programming language
- [GIN GONIC](https://github.com/gin-gonic/gin) - Gin is a web framework written in Go (Golang).
- [MongoDB](https://github.com/mongodb/mongo-go-driver) - The MongoDB supported driver for Go.
- [TESTIFY](https://github.com/stretchr/testify) - Go testing lib
- [Gomock](https://github.com/golang/mock)- Gomock is a mocking framework for the Go programming language

## Installation

### Useful commands

```bash
# To download dependencies and buid app
$ make build
```

```bash
# To run test
$ make tests
```

```bash
# To get test coverage
$ make coverage
```

```bash
# To clean cache and download fresh app dependencies
$ make setup
```

```bash
# Run service locally (it executes "go run src/main.go")
$ make run
```

### Docker

```bash

# Build the docker image first
$ make docker

# check if the containers are running
$ docker ps -a

# Run the build image with docker compose
$ make run-dc

# Stop docker compose
$ make stop

```

## Coding Style

Commits: <https://www.conventionalcommits.org/en/v1.0.0/>

Branching Model: <https://www.atlassian.com/git/tutorials/comparing-workflows/gitflow-workflow>

## Swagger

Not yet

## Authors

Mohsan Abbas
