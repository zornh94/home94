FROM golang:1.14.1 AS builder 
RUN mkdir $GOPATH/src/project
WORKDIR $GOPATH/src/project
COPY . .

RUN git config --global url."git@github.com".insteadOf "https://github.com"
RUN go mod tidy
RUN go env -w GOPRIVATE=github.com/tndevr/packages

RUN go build -o app cmd/main.go
