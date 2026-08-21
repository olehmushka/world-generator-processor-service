FROM golang:1.26-alpine

# NOTE: this is intentionally a single-stage build, not multi-stage. The
# world-generator-engine dependency locates its data files (languages,
# cultures, religions, ...) at runtime via an absolute path baked in at
# compile time (runtime.Caller), pointing into the Go module cache. A
# multi-stage build that copies only the compiled binary into a slim runtime
# image would leave that module cache behind and break data loading at
# startup. Keeping build and run in the same filesystem avoids that.

RUN apk add --no-cache ca-certificates

WORKDIR /usr/src/app

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN go build -o app .

EXPOSE 8080

# HTTP_SERVER_ADDRESS controls the actual bind address/port at runtime (see
# config/http_server.go); 8080 above is just the documented default.
CMD ["./app", "http_server_run"]
