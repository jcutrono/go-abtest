# Start from a Debian image with the latest version of Go installed
# and a workspace (GOPATH) configured at /go.
FROM golang

# Copy the local package files to the container's workspace.
ADD . /go/src/roller
ADD /src/. /go/src

WORKDIR /go/src/roller

# Build the outyet command inside the container.
# (You may fetch or manage dependencies here,
# either manually or with a tool like "godep".)
RUN go get github.com/gorilla/mux
RUN go get gopkg.in/mgo.v2

RUN go install

# Run the outyet command by default when the container starts.
ENTRYPOINT /go/bin/roller

# Document that the service listens on port 8080.
EXPOSE 8080
