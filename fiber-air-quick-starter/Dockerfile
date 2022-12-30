FROM golang:latest

COPY ./ /go/app

WORKDIR /go/app/

ENV GO111MODULE=on
ENV GOPATH=/

RUN go install github.com/cosmtrek/air@latest
RUN go mod download

CMD ["air", "-c", ".air.toml"]**