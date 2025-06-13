FROM golang:1.24 as builder

RUN go version
ENV GOPATH=/

COPY ./ ./

RUN apt-get update
RUN apt-get -y install postgresql-client

RUN chmod +x wait-for-postgres.sh

RUN go mod download
RUN go build -o test-task ./cmd/main.go


CMD ["./test-task"]