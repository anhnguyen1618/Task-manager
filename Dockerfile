# Start from a Debian image with the latest version of Go installed
# and a workspace (GOPATH) configured at /go.
FROM golang

WORKDIR /go/src/github.com/anhnguyen300795/Task-manager

# Copy the local package files to the container's workspace.
COPY . /go/src/github.com/anhnguyen300795/Task-manager

RUN go get ./...

# Document that the service listens on port 8080.
EXPOSE 8080

CMD [ "go", "run", "src/main.go" ]