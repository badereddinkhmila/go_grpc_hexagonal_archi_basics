FROM golang:1.20-alpine AS build

WORKDIR /app

COPY go.mod go.sum ./

COPY go.mod go.sum ./

COPY *.go ./

RUN CGO_ENABLED=0 GOOS=linux go build -o /docker-gs-ping

# Run
CMD ["/docker-gs-ping"]