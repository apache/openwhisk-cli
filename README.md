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

## Getting started

Here are some quick links to help you get started:

- [Downloading released binaries](#downloading-released-binaries) for Linux, macOS and Windows
- [Running the `wsk` CLI](#running-the-wsk-cli) executable
- [Building the project](#building-the-project) - download and build the GoLang source code
- [Contributing to the project](#contributing-to-the-project) - join us!

---

## Downloading released binaries

Executable binaries of the OpenWhisk CLI are available for download on the project's GitHub [releases page](https://github.com/apache/openwhisk-cli/releases).

We currently provide binaries for the following Operating Systems (OS) and architecture combinations:

Operating System | Architectures
--- | ---
Linux | i386, AMD64, ARM, ARM64, PPC64 (Power), S/390 and IBM Z
macOS (Darwin) | 386<sup>[1](#1)</sup>, AMD64
Windows | 386, AMD64

1. macOS, 32-bit (386) released versions are not available for builds using Go lang version 1.15 and greater.

We also provide instructions on how to build your own binaries from source code. See [Building the project](#building-the-project).

---

## Running the `wsk` CLI

You can copy the `wsk` binary to any folder, and add the folder to your system `PATH` in order to run the OpenWhisk CLI command from anywhere on your system. To get the CLI command help, execute the following:

```sh
$ wsk --help
```

To get CLI command debug information, include the `-d`, or `--debug` flag when executing this command.

---

## Building the project

### GoLang setup

The Openwhisk CLI is a GoLang program, so you will first need to [Download and install GoLang](https://golang.org/dl/) onto your local machine.

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
git clone --branch 1.1.0 git@github.com:apache/openwhisk-cli
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

If successful, an executable named `wsk` will be created in the project directory compatible with your current operating system and architecture.

Supported value combinations include:

`GOOS` | `GOARCH`
--- | ---
linux | 386 (32-bit), amd64 (64-bit), s390x (S/390, Z), ppc64le (Power), arm (32-bit), arm64 (64-bit)
darwin (macOS) | amd64
windows | 386 (32-bit), amd64 (64-bit)

### Build using Gradle

The project includes its own packaged version of Gradle called Gradle Wrapper which is invoked using the `./gradlew` command on Linux/Unix/Mac or `gradlew.bat` on Windows.

1. Gradle requires you to [install Java JDK version 8](https://gradle.org/install/) or higher

1. Clone the `openwhisk-cli` repo:

    ```sh
    git clone https://github.com/apache/openwhisk-cli
    ```

    and change into the project directory.

1. Cross-compile binaries for all supported Operating Systems and Architectures:

    ```sh
    ./gradlew goBuild
    ```

    Upon a successful build, the `wsk` binaries can be found under the corresponding `build/<os>-<architecture>/` folder of your project:

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

> **Note** You may use the `compile` Gradle task to build a subset of the supported platforms using the `buildPlatforms` parameter and supplying a comma-separated list, for example:
`-PbuildPlatforms=linux-amd64,mac-amd64,windows-amd64`

#### Using your own local Gradle to build

Alternatively, you can choose to [Install Gradle](https://gradle.org/install/) and use it instead of the project's Gradle Wrapper.  If so, you would use the `gradle` command instead of `gradlew`. If you do elect to use your own Gradle, verify its version is `6.8.1` or higher:

```sh
gradle -version
```

> **Note** If using your own local Gradle installation, use the `gradle` command instead of the `./gradlew` command in the build instructions below.

### Building for internationalization (i18n)

The CLI internationalization is generated dynamically using the `bindata` tool as part of the gradle build. If you need to install it manually, you may use:

```sh
$ go get -u github.com/jteeuwen/go-bindata/...
$ go-bindata -pkg wski18n -o wski18n/i18n_resources.go wski18n/resources
```

> **Note**: the `go-bindata` package will automatically be installed if the `go build` command is used in the project as it is listed in the `go.mod` dependency file.

### Running unit tests

##### Using Go

```sh
$ cd commands
$ go test -tags=unit -v
```

> **Note** A large number of CLI tests today are not yet available as Go tests.

##### Using gradle

All tests can be run using the Gradle script:

```sh
$ ./gradlew goTest -PgoTags=unit
$ ./gradlew goTest -PgoTags=native
```

### Running integration tests

Integration tests are best left to the Travis build as they depend on a fully functional OpenWhisk environment.

---

## Contributing to the project

### Git repository setup

1. [Fork](https://docs.github.com/en/github/getting-started-with-github/fork-a-repo) the Apache repository

    If you intend to contribute code, you will want to fork the `apache/openwhisk-cli` repository into your github account and use that as the source for your clone.

1. Clone the repository from your fork:

    ```sh
    git clone git@github.com:${GITHUB_ACCOUNT_USERNAME}/openwhisk-cli.git
    ```

1. Add the Apache repository as a remote with the `upstream` alias:

    ```sh
    git remote add upstream git@github.com:apache/openwhisk-cli
    ```

    You can now use `git push` to push local `commit` changes to your `origin` repository and submit pull requests to the `upstream` project repository.

1. Optionally, prevent accidental pushes to `upstream` using this command:

    ```sh
    git remote set-url --push upstream no_push
    ```

> Be sure to [Sync your fork](https://docs.github.com/en/github/collaborating-with-issues-and-pull-requests/syncing-a-fork) before starting any contributions to keep it up-to-date with the upstream repository.

### Adding new dependencies

Please use `go get` to add new dependencies to the `go.mod` file:

```sh
go get -u github.com/project/libname@v1.2.0
```

> Please avoid using commit hashes for referencing non-OpenWhisk libraries.

### Removing unused dependencies

Please us `go tidy` to remove any unused dependencies after any significant code changes:

```sh
go mod tidy
```

### Updating dependency versions

Although you might be tempted to edit the go.mod file directly, please use the recommended method of using the `go get` command:

```sh
go get -u github.com/project/libname  # Using "latest" version
go get -u github.com/project/libname@v1.1.0 # Using tagged version
go get -u github.com/project/libname@aee5cab1c  # Using a commit hash
```

### Updating Go version

Although you could edit the version directly in the go.mod file, it is better to use the `go edit` command:

```sh
go mod edit -go=1.15
```

---

## Continuous Integration

Travis CI is used as a continuous delivery service for Linux and Mac. Currently, Travis CI supports the environments of Linux and Mac, but it is not available for Windows. The project would like to add AppVeyor CI in the future to run test cases for Windows.
