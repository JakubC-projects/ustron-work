## Build server
FROM golang:1.21-bullseye as server
WORKDIR /app
COPY server/go.mod server/go.sum ./
RUN go mod download

COPY server/work work

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o main work/cmd/server/main.go

FROM node:18-buster as client

WORKDIR /app
COPY client/package.json client/package-lock.json client/index.html client/tsconfig.json client/tsconfig.node.json client/vite.config.ts client/tailwind.config.js client/postcss.config.js ./
RUN npm ci

COPY client/src src
RUN npm run build


## Deploy
FROM alpine:3.15
WORKDIR /
COPY --from=server /app/main /usr/bin/
COPY --from=client /app/dist /frontend
ENV "FRONTEND_LOCATION" "/frontend"
ENTRYPOINT ["main"]
