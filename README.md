
# OpenWhisk Command Line Interface `wsk`

The OpenWhisk Command Line Interface (OpenWhisk CLI) is a unified tool that provides a consistent interface to
interact with OpenWhisk services. With this tool to download and configure, you are able to manage OpenWhisk services
from the command line and automate them through scripts.


# Where to download the binary of OpenWhisk CLI

The OpenWhisk CLI is available on the release page: [click here to download](https://github.com/openwhisk/openwhisk-cli/releases).
We currently have binaries available for Linux, Mac OS and windows under amd64 architecture. You can download the
binary, which fits your local environment.


# How to build the binary locally

You can also choose to build the binary locally based on the source code. First, install the prerequisites to 
download and build OpenWhisk CLI: [installing Git](https://git-scm.com/book/en/v2/Getting-Started-Installing-Git).

Then, download the source code via the Git command:

```
$ git clone https://github.com/openwhisk/openwhisk-cli.git
```

OpenWhisk CLI(`wsk`) is produced in a Docker container during the build process which is copied from the
Docker container to the local file system in the following directory: bin. This binary will be platform
specific, it will only run on the operating system, and CPU architecture that matches the build machine.

## Build the binary with Go

The binary can be built by Go build command. Make sure that you have Go installed: [installing Go](https://golang.org/doc/install).

After that, open an terminal, go to the directory of OpenWhisk CLI home directory, and build the binary via
the following command:

```
$ go build -o wsk
```

If you would like to build the binary for a specific operating system, you may add the arguments GOOS and
GOARCH into the Go build command. Since it is only applicable under amd64 architecture, you have to set GOARCH
to amd64. GOOS can be set to "linux" "darwin" or "windows".

For example, run the following command to build the binary for Linux:

```
$ GOOS=linux GOARCH=amd64 go build -o wsk
```

If it is executed successfully, you can find your binary `wsk` directly under OpenWhisk CLI home directory.

## Build the binary with Docker and Gradle

This is the second choice for you to build the binary. Make sure that you have Docker and gradle on your machine:
[installing Docker](https://docs.docker.com/engine/installation/) and [installing Gradle](https://gradle.org/install) for your local machine.

After that, open an terminal, go to the directory of OpenWhisk CLI home directory, and
build the binary via the following command under Linux or Mac:

```
$ ./gradlew distDocker
```

or run the following command for Windows:

```
$ ./gradlew.bat distDocker
```

Finally, you can find the binary `wsk` or `wsk.exe` in the bin folder under the OpenWhisk CLI home directory.


# How to use the binary

When you have the binary, you can copy the binary to any folder, and add folder into the system PATH in order to
run the OpenWhisk CLI command. To get the CLI command help, execute the following command:

```
$ wsk --help
```

To get CLI command debug information, include the -d, or --debug flag when executing this command.


# Continuous Integration

In order to build OpenWhisk CLI binaries with good quality, OpenWhisk CLI uses Travis CI as the continuous
delivery service for Linux and Mac, and AppVeyor CI for Windows. OpenWhisk CLI is a Go project. Currently Travis
CI supports the environments of Linux and Mac, but it is not available for Windows, which is the reason why we
uses AppVeyor CI to run the test cases and build the binary for Windows.
