# syntax=docker/dockerfile:1

# Build the application from source
FROM golang:1.23 AS build-stage

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY *.go ./

RUN CGO_ENABLED=0 GOOS=linux go build -o /gin_server

# Deploy the application binary into a lean image
FROM gcr.io/distroless/base-debian11 AS build-release-stage

WORKDIR /

COPY --from=build-stage /gin_server /gin_server

EXPOSE 8080

USER nonroot:nonroot

ENTRYPOINT ["/gin_server"]