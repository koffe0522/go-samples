FROM golang:latest

WORKDIR /go
ADD . /go

ENV GO111MODULE=on
ENV GOPATH=

RUN go mod download

EXPOSE 8000

CMD ["go", "run", "main.go"]
