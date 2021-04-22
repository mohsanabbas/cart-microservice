# Build Stage
FROM golang:1.16.0-alpine3.13 as builder

WORKDIR /src

RUN apk update && apk upgrade && \
  apk --update add git make && \
  apk add --no-cache openssh

COPY . .

RUN make setup && make engine

# Final Stage
FROM alpine:latest

RUN apk update && apk upgrade && \
  apk --update --no-cache add tzdata && \
  mkdir /src

WORKDIR /src

EXPOSE 8090

COPY --from=builder /src/engine /src

CMD /src/engine

LABEL Name=cart-microservice Version=0.0.1
