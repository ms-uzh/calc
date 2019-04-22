# Base build image
FROM golang:1.12-alpine AS builder
 
# Install some dependencies needed to build the project
RUN apk add ca-certificates git
WORKDIR /go/src/github.com/fforootd/calc
 
# Force the go compiler to use modules
ENV GO111MODULE=on

# Here we copy the rest of the source code
COPY . .
# And compile the project
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o app -mod=vendor
 
#In this last stage, we start from a fresh Alpine image, to reduce the image size and not ship the Go compiler in our production artifacts.
FROM alpine AS app
# We add the certificates to be able to verify remote weaviate instances
RUN apk add ca-certificates
# Finally we copy the statically compiled Go binary.
COPY /config /config
COPY /templates /templates
COPY /templates/html /templates/html
COPY /html /html
COPY --from=builder /go/src/github.com/fforootd/calc/app /app
ENTRYPOINT ["/app"]