FROM golang:1.20.0

WORKDIR /usr/src/camagru

COPY . .

RUN go mod tidy