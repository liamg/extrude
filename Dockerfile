FROM golang:alpine AS builder
RUN apk update && apk add --no-cache git
WORKDIR /build
COPY . .
RUN CGO_ENABLED=0 go build -o /extrude ./cmd/extrude
FROM scratch
COPY --from=builder /extrude /extrude
ENTRYPOINT ["/extrude"]
