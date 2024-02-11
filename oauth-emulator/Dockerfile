## Build
FROM golang:1.22-bullseye as build
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o main main.go

ENTRYPOINT ["./main"]

## Deploy
FROM debian:trixie-slim
WORKDIR /

RUN apt-get -y update; apt-get -y install curl

COPY --from=build /app/main /usr/bin/
ENTRYPOINT ["main"]
