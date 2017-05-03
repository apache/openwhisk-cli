FROM golang:1.7

# Install zip
RUN apt-get -y update && \
    apt-get -y install zip

ENV GOPATH=/

# Download and install tools
RUN echo "Installing the godep tool"
RUN go get github.com/tools/godep

ADD . /src/github.com/openwhisk/openwhisk-cli

# Load all of the dependencies from the previously generated/saved godep generated godeps.json file
RUN echo "Restoring Go dependencies"
RUN cd /src/github.com/openwhisk/openwhisk-cli && /bin/godep restore -v

# Collect all translated strings into single .go module
RUN echo "Packaging i18n Go module"
RUN cd /src/github.com/openwhisk/openwhisk-cli && /bin/go-bindata -pkg wski18n -o wski18n/i18n_resources.go wski18n/resources

# wsk binary will be placed under a build folder
RUN mkdir /src/github.com/openwhisk/openwhisk-cli/build

ARG CLI_OS
ARG CLI_ARCH

# Build the Go wsk CLI binaries and compress resultant binaries
RUN chmod +x /src/github.com/openwhisk/openwhisk-cli/build.sh
RUN cd /src/github.com/openwhisk/openwhisk-cli && ./build.sh
