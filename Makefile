BINARY=engine
VERSION = $(shell git describe --tags || echo "developer")


test:
	go test -v -cover -covermode=atomic ./...

engine:
	go build -o ${BINARY} src/*.go

coverage:
	go mod vendor
	go test ./... -coverprofile=coverage.out -covermode=count -mod=vendor
	go tool cover -func=coverage.out

setup:
	go clean -modcache
	go get -u ./...
	go mod vendor

build:
	go mod vendor
	go build -o bin/cart-microservice -mod=vendor

vendor:
	@go mod vendor

clean:
	go clean
	-find . -name ".out" -exec rm -f {} \;

docker:
	docker build -t cart-microservice .

run-dc:
	docker-compose up --build -d

run:
	go run src/main.go

stop:
	docker-compose down

lint-prepare:
	@echo "Installing golangci-lint"
	curl -sfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh| sh -s latest

lint:
	./bin/golangci-lint run ./...

.PHONY: clean install unittest build docker run stop vendor lint-prepare lint