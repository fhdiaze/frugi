FROM golang:1.23.0-alpine3.20 AS build

WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download

COPY ./ ./
RUN CGO_ENABLED=0 GOOS=linux go build -o ./frugi ./cmd/frugi


FROM alpine:edge

WORKDIR /app

COPY --from=build /app/frugi .

ENTRYPOINT ["/app/frugi"]