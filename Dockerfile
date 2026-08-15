FROM golang:alpine AS builder

WORKDIR /src

RUN apk add --no-cache git

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 go build -ldflags "-s -w" -o ./out/infrared ./cmd/infrared

FROM alpine:latest

RUN apk add --no-cache ca-certificates \
    && addgroup -S infrared \
    && adduser -S infrared -G infrared

WORKDIR /app

COPY --from=builder /src/out/infrared /app/infrared

USER infrared

ENTRYPOINT [ "/app/infrared" ]
