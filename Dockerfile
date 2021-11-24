FROM golang:alpine AS builder
RUN apk update && apk add --no-cache git
WORKDIR /build
COPY . .
FROM scratch
COPY --from=builder /build/extrude /extrude
ENTRYPOINT ["/extrude"]
