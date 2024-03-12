FROM golang:1.22-alpine as builder

ARG CI_COMMIT_SHA
ARG CI_PROJECT_NAME

# Build directories
RUN mkdir -p /go/src/github.com/felipecoppola/go-boilerplate/app
WORKDIR /go/src/github.com/felipecoppola/go-boilerplate/app

# Copy app
COPY . .

# Compile
RUN GOOS=linux go build \
      -o api \
      -ldflags "-X main.appVersion=${CI_COMMIT_SHA} -X main.appName=${CI_PROJECT_NAME}" \
      ./cmd/app/

# Build 2
FROM alpine:3.17

# Copy binary from builder
COPY --from=builder /go/src/github.com/felipecoppola/go-boilerplate/app/api /usr/bin/
