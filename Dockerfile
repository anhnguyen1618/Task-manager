FROM alpine:latest

ENV GO_VERSION go-1.9.2-r1

ENV GOPATH /go

ENV PROJECT_DIR github.com/anhnguyen300795/Task-manager

ENV WORKDIR = ${GOPATH}/src/${PROJECT_DIR}/src

WORKDIR ${WORKDIR}

# Copy the local package files to the container's workspace.
COPY ./src .

# Update list of softwares that need updating & upgrade those softwares to latest version
RUN apk update && apk upgrade && \
    # add extra package for installation 
    apk add git go=$GO_VERSION musl-dev && \
    # remove cached info
    rm -rf /var/cache/apk/*

# Recursively add excecution permission for files under scripts folder
RUN chmod -R +x ./scripts

# Generate log file text
RUN cd scripts && ./log.sh


# Install golang project's dependencies
RUN go get ${PROJECT_DIR}/...

# Document that the service listens on port 8080.
EXPOSE 8080

ENTRYPOINT [ "./scripts/start.sh" ]