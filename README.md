[![Build Status](https://travis-ci.org/skywinder/ActionSheetPicker-3.0.svg?branch=master)](https://travis-ci.org/skywinder/ActionSheetPicker-3.0)
[![Quality Gate](https://sonar.app.cvc.com.br/api/project_badges/measure?project=shopping-cart-recomendations&metric=alert_status)](https://sonar.app.cvc.com.br/dashboard?id=shopping-cart-recomendations)
[![Coverage](https://sonar.app.cvc.com.br/api/project_badges/measure?project=shopping-cart-recomendations&metric=coverage)](https://sonar.app.cvc.com.br/component_measures?id=shopping-cart-recomendations&metric=Coverage)
[![Maintainnability](https://sonar.app.cvc.com.br/api/project_badges/measure?project=shopping-cart-recomendations&metric=sqale_rating)](https://sonar.app.cvc.com.br/component_measures?id=shopping-cart-recomendations&metric=Maintainability)
[![Security](https://sonar.app.cvc.com.br/api/project_badges/measure?project=shopping-cart-recomendations&metric=security_rating)](https://sonar.app.cvc.com.br/component_measures?id=shopping-cart-recomendations&metric=Security)

# SHOPPING-CART MICROSERVICE IN GO

Shopping cart microservice written in `GO` to perform CRUD operations.

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

### Docker

```bash
# Build the docker image first
$ make docker

# Run the application
$ make run

# check if the containers are running
$ docker ps -a

# Stop
$ make stop
```

## Coding Style

Commits: <https://www.conventionalcommits.org/en/v1.0.0/>

Branching Model: <https://www.atlassian.com/git/tutorials/comparing-workflows/gitflow-workflow>

## Swagger

not yet

## Authors

Mohsan Abbas
