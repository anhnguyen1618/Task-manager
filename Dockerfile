# Start from a Debian image with the latest version of Go installed
# and a workspace (GOPATH) configured at /go.
FROM golang

WORKDIR /go/src/github.com/anhnguyen300795/Task-manager/src

# Copy the local package files to the container's workspace.
COPY ./src /go/src/github.com/anhnguyen300795/Task-manager/src

RUN go get ./...

# Document that the service listens on port 8080.
EXPOSE 8080

CMD [ "go", "run", "main.go" ]