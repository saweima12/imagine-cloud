FROM golang:1.20-buster


COPY . /app
WORKDIR /app

RUN go build -o ./build/imagine
RUN cp ./build/imagine /usr/bin

EXPOSE 8000

