FROM golang:1.19-buster

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download && go mod verify

COPY . .
RUN go build -v -o /usr/local/bin
