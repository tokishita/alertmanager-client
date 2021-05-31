FROM golang:1.14.4

WORKDIR /go/src

COPY go.mod main.go  ./

ARG CGO_ENABLED=0
ARG GOOS=linux
ARG GOARCH=amd64
RUN go build \
    -o /go/bin/main \
    -ldflags '-s -w'

ENTRYPOINT ["/go/bin/main"]

