# STAGE 1: Build
FROM golang:1.25-alpine AS builder
ARG service

WORKDIR /src

# Packages for the go download step
RUN apk add --no-cache ca-certificates git

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN go build -o ./app ./$service/cmd/main.go

# STAGE 2: Runtime
FROM alpine:3.20

# TLS-Certificates for HTTPS calls
RUN apk add --no-cache ca-certificates && update-ca-certificates

# Running the binary as a non-root user
RUN adduser -D -H -s /sbin/nologin appuser

USER appuser

EXPOSE 8080
ENTRYPOINT ["./app"]

