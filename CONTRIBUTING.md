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

# Contributing to OpenWhisk CLI

## Set up the development environment

In order to develop OpenWhisk CLI on your local machine. First, install the prerequisites to
download and build OpenWhisk CLI: [installing Git](https://git-scm.com/book/en/v2/Getting-Started-Installing-Git).

Then, save the project in the location compliant with the Go standard naming convention, which means you need to
created a directory named $GOPATH/src/github.com/apache/ and download the source code via the following commands:

```
$ cd $GOPATH/src/github.com/apache/
$ git clone https://github.com/apache/openwhisk-cli.git
```

After cloning the source code, you need to install all the dependencies by running the command under openwhisk cli folder:

```
$ go get -d -t ./...
```

or

```
$ make deps
```

You should be able to build the binaries with either the go command or the Gradle command, which is available in [README](https://github.com/apache/openwhisk-cli/blob/master/README.md).


## Proposing new features

If you would like to implement a new feature, please [raise an issue](https://github.com/apache/openwhisk-cli/issues) before sending a pull request so the feature can be discussed.
This is to avoid you spending your valuable time working on a feature that the project developers are not willing to accept into the code base.

## Fixing bugs

If you would like to fix a bug, please [raise an issue](https://github.com/apache/openwhisk-cli/issues) before sending a pull request so it can be discussed.
If the fix is trivial or non controversial then this is not usually necessary.

## Merge approval

The project maintainers use LGTM (Looks Good To Me) in comments on the code review to
indicate acceptance. A change requires LGTMs from two of the maintainers of each
component affected.

## Communication

Please use [Slack channel #whisk-users](https://cloudplatform.slack.com/messages/whisk_cli).

## Setup

Project was written with `Go v1.9`. It has a dependency on [openwhisk-client-go](https://github.com/apache/openwhisk-client-go).

## Testing

This repository needs unit tests.

Please provide information that helps the developer test any changes they make before submitting.

## Coding style guidelines

Use idomatic go. Document exported functions.

# Publishing Tagged Release to Homebrew

 [Homebrew](https://brew.sh) is used to install `wsk` locally. Once we release a new version of `wsk` we should update its version in homebrew.

 Get the new release SHA256 checksum by downloading the Source Code (tar.gz) from the [releases page](https://github.com/apache/openwhisk-cli/releases) and running `shasum -a 256 X.Y.Z.tar.gz` on the tarball.

 Update brew formula with the automation command `brew bump-formula-pr`:
 ```bash
 $ brew bump-formula-pr \
   --url='https://github.com/apache/openwhisk-cli/archive/X.Y.Z.tar.gz' \
   --sha256='PASTE THE SHA256 CHECKSUM HERE' \
   wsk
 ```
