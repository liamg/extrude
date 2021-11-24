FROM golang:alpine AS builder
RUN apk update && apk add --no-cache git
WORKDIR /src
COPY . .
RUN GO111MODULE=off CGO_ENABLED=0 go build -o /extrude ./cmd/extrude
FROM scratch
COPY --from=builder /extrude /extrude
ENTRYPOINT ["/extrude"]
