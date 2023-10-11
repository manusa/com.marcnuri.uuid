FROM golang:1.21 AS builder
WORKDIR /app

# Modules
COPY go.mod go.sum ./
RUN go mod download

# Build
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build

RUN ls -la

FROM builder AS tester
RUN go test

FROM gcr.io/distroless/static-debian12 AS app

MAINTAINER Marc Nuri <marc@marcnuri.com>
LABEL MAINTAINER="Marc Nuri <marc@marcnuri.com>"

COPY --from=builder /app/com.marcnuri.uuid /app

EXPOSE 8080
USER 1000

ENTRYPOINT ["/app"]
