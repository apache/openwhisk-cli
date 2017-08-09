FROM golang:1.8

# Install zip
RUN apt-get -y update && \
    apt-get -y install zip

ENV GOPATH=/

# Download and install tools
RUN echo "Installing the godep tool"
RUN go get github.com/tools/godep

ADD . /src/github.com/apache/incubator-openwhisk-cli

# Load all of the dependencies from the previously generated/saved godep generated godeps.json file
RUN echo "Restoring Go dependencies"
RUN cd /src/github.com/apache/incubator-openwhisk-cli && /bin/godep restore -v

# wsk binary will be placed under a build folder
RUN mkdir /src/github.com/apache/incubator-openwhisk-cli/build

ARG CLI_OS
ARG CLI_ARCH

# Build the Go wsk CLI binaries and compress resultant binaries
RUN chmod +x /src/github.com/apache/incubator-openwhisk-cli/build.sh
RUN cd /src/github.com/apache/incubator-openwhisk-cli && ./build.sh
