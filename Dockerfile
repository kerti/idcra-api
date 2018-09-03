# Start from a Debian image with the latest version of Go installed
# and a workspace (GOPATH) configured at /go.
FROM golang:latest

ENV SRC_DIR=/go/src/github.com/kerti/idcra-api/
# Add the source code:
COPY . $SRC_DIR
WORKDIR $SRC_DIR

# Build it:
RUN go get -v github.com/jteeuwen/go-bindata/...
RUN export PATH=$PATH:$GOPATH/bin
RUN go generate ./schema
RUN go get -v ./
RUN go build
ENTRYPOINT ["./idcra-api"]
EXPOSE 3000