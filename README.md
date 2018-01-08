# OpenWhisk Command Line Interface `wsk`
[![License](https://img.shields.io/badge/license-Apache--2.0-blue.svg)](http://www.apache.org/licenses/LICENSE-2.0)
[![Build Status](https://travis-ci.org/apache/incubator-openwhisk-cli.svg?branch=master)](https://travis-ci.org/apache/incubator-openwhisk-cli)

# Disclaimer

This project is currently on an experimental stage. We periodically synchronize
the source code and test cases of this repository with the [CLI
folder](https://github.com/apache/incubator-openwhisk/tree/master/tools/cli/go-whisk-cli)
and the [test
folder](https://github.com/apache/incubator-openwhisk/tree/master/tests) in
OpenWhisk. The framework of test cases is under construction for this
repository. Please contribute to the [CLI
folder](https://github.com/apache/incubator-openwhisk/tree/master/tools/cli/go-whisk-cli)
in OpenWhisk for any CLI changes, before we officially announce the separation
of OpenWhisk CLI from OpenWhisk.


The OpenWhisk Command Line Interface (OpenWhisk CLI) is a unified tool that
provides a consistent interface to interact with OpenWhisk services. With this
tool to download and configure, you are able to manage OpenWhisk services from
the command line and automate them through scripts.


# Where to download the binary of OpenWhisk CLI

The OpenWhisk CLI is available on the release page: [click here to
download](https://github.com/apache/incubator-openwhisk-cli/releases). We
currently have binaries available for Linux, Mac OS and windows under amd64
architecture. You can download the binary, which fits your local environment.


# How to build the binary locally

You can also choose to build the binaries locally from the source code with Go tool.

Make sure that you have Go installed [installing
Go](https://golang.org/doc/install), and `$GOPATH` is defined [Go development
environment](https://golang.org/doc/code.html).

Then download the source code of the OpenWhisk CLI and the dependencies by typing:

```
$ cd $GOPATH
$ go get github.com/apache/incubator-openwhisk-cli
```


## Build the binary with Go

Open an terminal, go to the directory of OpenWhisk CLI home directory, and build
the binary via the following command:

```
$ go build -o wsk
```

If you would like to build the binary for a specific operating system, you may
add the arguments GOOS and GOARCH into the Go build command. Since it is only
applicable under amd64 architecture, you have to set GOARCH to amd64. GOOS can
be set to "linux" "darwin" or "windows".

For example, run the following command to build the binary for Linux:

```
$ GOOS=linux GOARCH=amd64 go build -o wsk
```

If it is executed successfully, you can find your binary `wsk` directly under
OpenWhisk CLI home directory.

## Build the binary with Gradle

Open a terminal, go to the directory of OpenWhisk CLI home
directory, and build the binary via the following command under Linux or Mac:

```
$ ./gradlew build
```

or run the following command for Windows:

```
$ ./gradlew.bat build
```

Finally, you can find the binary `wsk` or `wsk.exe` in the bin folder under the
OpenWhisk CLI home directory. 

If you would like to build the binaries available for all the operating systems
and architectures, run the following command: 

```
$ ./gradlew build -PcrossCompile=true
```

Then, you will find the binaries and their compressed packages generated under
the folder bin/\<os\>/\<cpu arc\>/ for each operating system and CPU
architecture pair. We supports both amd64 and 386 for Linux, Mac and Windows
operating systems. We also support s390x, ppc64le, arm and arm64 for Linux.

# How to use the binary

When you have the binary, you can copy the binary to any folder, and add folder
into the system PATH in order to run the OpenWhisk CLI command. To get the CLI
command help, execute the following command: 

```
$ wsk --help
```

To get CLI command debug information, include the -d, or --debug flag when
executing this command. 

# Continuous Integration

In order to build OpenWhisk CLI binaries with good quality, OpenWhisk CLI uses
Travis CI as the continuous delivery service for Linux and Mac. OpenWhisk CLI is
a Go project. Currently Travis CI supports the environments of Linux and Mac,
but it is not available for Windows. We will add support of AppVeyor CI in
future to run the test cases and build the binary for Windows. 