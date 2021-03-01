<!--
#
# Licensed to the Apache Software Foundation (ASF) under one or more
# contributor license agreements.  See the NOTICE file distributed with
# this work for additional information regarding copyright ownership.
# The ASF licenses this file to You under the Apache License, Version 2.0
# (the "License"); you may not use this file except in compliance with
# the License.  You may obtain a copy of the License at
#
#     http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.
#
-->

# OpenWhisk Command-line Interface `wsk`

[![Build Status](https://travis-ci.com/apache/openwhisk-cli.svg?branch=master)](https://travis-ci.com/apache/openwhisk-cli)
[![License](https://img.shields.io/badge/license-Apache--2.0-blue.svg)](http://www.apache.org/licenses/LICENSE-2.0)
[![Join Slack](https://img.shields.io/badge/join-slack-9B69A0.svg)](http://slack.openwhisk.org/)
[![Twitter](https://img.shields.io/twitter/follow/openwhisk.svg?style=social&logo=twitter)](https://twitter.com/intent/follow?screen_name=openwhisk)

OpenWhisk Command-line Interface (CLI) is a unified tool that provides a consistent interface to interact with OpenWhisk services.

---

## Downloading released binaries

Binaries of the OpenWhisk CLI are available for download on the project's GitHub [releases page](https://github.com/apache/openwhisk-cli/releases).

We currently provide binaries for the following Operating Systems (OS) and architecture combinations:

Operating System | Architectures
--- | ---
Linux | i386, AMD64, ARM, ARM64, PPC64 (Power), S/390 and IBM Z
Mac OS (Darwin) | 386<sup>[1](#1)</sup>, AMD64
Windows | 386, AMD64

1. Mac OS, 32-bit (386) released versions are not available for builds using Go lang version 1.15 and greater.

We also provide instructions on how to build your own binaries from source code using the `Go` tool.

- See [Building the project](#building-the-project).

---

## Building the project

### GoLang setup

The Openwhisk CLI is a GoLang program so you will first need to [Download and install GoLang](https://golang.org/dl/) onto your local machine.

> **Note** Go version 1.15 or higher is recommended

Make sure your `$GOPATH` is defined correctly in your environment. For detailed setup of your GoLang development environment, please read [How to Write Go Code](https://golang.org/doc/code.html).

### Download the source code from GitHub

As the code is managed using GitHub, it is easiest to retrieve the code using the `git clone` command.

if you just want to build the code and do not intend to be a Contributor, you can clone the latest code from the Apache repository:

```sh
git clone git@github.com:apache/openwhisk-cli
```

or you can specify a release (tag) if you do not want the latest code by using the `--branch <tag>` flag. For example, you can clone the source code for the tagged 1.1.0 [release](https://github.com/apache/openwhisk-cli/releases/tag/1.1.0)

```sh
git clone --branch 1.1.0 git@github.com:apache/openwhisk-clie
```

You can also pull the code from a fork of the repository. If you intend to become a Contributor to the project, read the section [Contributing to the project](#contributing-to-the-project) below on how to setup a fork.

### Build using `go build`

Use the Go utility to build the ```wsk`` binary.

Change into the cloned project directory and use `go build` with the target output name for the binary:

```sh
$ go build -o wsk
```

an executable named `wsk` will be created in the project directory compatible with your current operating system and architecture.

#### Building for other Operating Systems (GOOS) and Architectures (GOARCH)

If you would like to build the binary for a specific operating system and processor architecture, you may add the arguments `GOOS` and `GOARCH` into the Go build command (as inline environment variables).

For example, run the following command to build the binary for 64-bit Linux:

```sh
$ GOOS=linux GOARCH=amd64 go build -o wsk
```

Supported value combinations include:

`GOOS` | `GOARCH`
--- | ---
linux | 386 (32-bit), amd64 (64-bit), s390x (S/390, Z), ppc64le (Power), arm (32-bit), arm64 (64-bit)
darwin (Mac OS) | amd64
windows | 386 (32-bit), amd64 (64-bit)

### Build using Gradle

The project includes its own packaged version of Gradle called Gradle Wrapper which is invoked using the `./gradlew` command on Linux/Unix/Mac or `gradlew.bat` on Windows.

1. Gradle requires requires you to [install Java JDK version 8](https://gradle.org/install/) or higher

1. Clone the `openwhisk-cli` repo:

    ```sh
    git clone https://github.com/apache/openwhisk-cli
    ```

    and change into the project directory.

1. Cross-compile binaries for all supported Operating Systems and Architectures:

    ```sh
    ./gradlew goBuild
    ```

    Upon a successful build, the `wsk` binaries can be found within the under the corresponding `build/<os>-<architecture>/` folder of your project:

    ```sh
    $ ls build
    darwin-amd64  linux-amd64   linux-arm64   linux-s390x   windows-amd64
    linux-386     linux-arm     linux-ppc64le windows-386
    ```

#### Compiling for a single OS/ARCH

1. View gradle build tasks for supported Operating Systems and Architectures:

    ```sh
    ./gradlew tasks
    ```

    you will see build tasks for supported OS/ARCH combinations:

    ```sh
    Gogradle tasks
    --------------
    buildDarwinAmd64 - Custom go task.
    buildLinux386 - Custom go task.
    buildLinuxAmd64 - Custom go task.
    buildLinuxArm - Custom go task.
    buildLinuxArm64 - Custom go task.
    buildLinuxPpc64le - Custom go task.
    buildLinuxS390x - Custom go task.
    buildWindows386 - Custom go task.
    buildWindowsAmd64 - Custom go task.
    ```

    > **Note**: The `buildWindows386` option is only supported on Golang versions less than 1.15.

1. Build using one of these tasks, for example:

    ```sh
    $ ./gradlew buildDarwinAmd64
    ```

#### Using your own local Gradle to build

Alternatively, you can choose to [Install Gradle](https://gradle.org/install/), which you can use instead of Gradle Wrapper by using the `gradle` command instead of `gradlew`. If you wish to use you own Gradle, verify its version is `6.6` or higher:

```sh
gradle -version
```

> **Note** If using your own local Gradle installation, use the `gradle` command instead of the `./gradlew` command in the build instructions below.

<!-->
```
$ ./gradlew compile
```

The build script will place the binaries into the folder `build/<os>-<cpu arc>/`
for each operating system and CPU architecture pair. The build supports both
amd64 and 386 for Linux, Mac and Windows operating systems, as well as Power,
64-bit ARM, and S390X architectures for Linux.

A binary compatible with the local architecture will be placed at `build/wsk`
(`build\wsk.exe` on Windows).

To specify a build for specific architectures, you can provide a comma or
space-delimited list of hyphenated os-architecture pairs, like this:

```
$ ./gradlew compile -PbuildPlatforms=linux-amd64,mac-amd64,windows-amd64
```

The build library understands most representations of most Operating Systems.

-->

### Running unit tests

##### Using Go

```sh
$ cd commands
$ go test -tags=unit -v
```

> **Note** A large number of CLI tests today are not yet available as Go tests.

##### Using gradle

All tests can be run using the Gradle script:

```
$ ./gradlew goTest -PgoTags=unit
$ ./gradlew goTest -PgoTags=native
```

### Running integration tests

Integration tests are best left to the Travis build as they depend on a fully functional OpenWhisk environment.

<!-->
## Compile the binary using your local Go environment

Make sure that you have [Go installed](https://golang.org/doc/install), and `$GOPATH` is defined in your [Go development environment](https://golang.org/doc/code.html).

Then download the source code of the OpenWhisk CLI and the dependencies by typing:

```
$ go get github.com/apache/openwhisk-cli
$ cd $GOPATH/src/github.com/apache/openwhisk-cli
```

The CLI internationalization is generated dynamically using the `bindata` tool as part of the gradle build. If you need to install it manually, you may use:

```
$ go get -u github.com/jteeuwen/go-bindata/...
$ go-bindata -pkg wski18n -o wski18n/i18n_resources.go wski18n/resources
```

Now you can build the binary.
```
$ go build -o wsk
```

If you would like to build the binary for a specific operating system, you may add the arguments `GOOS` and `GOARCH` into the Go build command. `GOOS` can be set to `linux`, `darwin`, or `windows`.

For example, run the following command to build the binary for Linux:

```
$ GOOS=linux GOARCH=amd64 go build -o wsk-$GOOS-$GOARCH
```

If it is executed successfully, you can find your binary `wsk` directly under OpenWhisk CLI home directory.

-->

# How to use the binary

When you have the binary, you can copy the binary to any folder, and add folder into the system PATH in order to run the OpenWhisk CLI command. To get the CLI command help, execute the following command:

```
$ wsk --help
```

To get CLI command debug information, include the `-d`, or `--debug` flag when executing this command.

# Continuous Integration

Travis CI is used as a continuous delivery service for Linux and Mac. Currently Travis CI supports the environments of Linux and Mac, but it is not available for Windows. We will add support of AppVeyor CI in future to run test cases and build the binary for Windows.
