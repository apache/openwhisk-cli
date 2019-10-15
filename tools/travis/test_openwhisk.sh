#!/bin/bash
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

set -e

#
#  At this point, the Travis build should already have built the binaries and
#  the release.  If you're running manually, this command should get you to
#  the same place:
#
#    ./gradlew releaseBinaries
#
#  Also at this point, you should already have the openwhisk main repo. pulled down
#  from gradle in the parent directory, using a command such as:
#
#    git clone --depth 3 https://github.com/apache/openwhisk.git
#
#  To be clear, your directory structure will look something like...
#
#      $HOMEDIR
#       |- openwhisk
#       |- openwhisk-cli (This project)
#       |- openwhisk-utilities (For scancode)
#
#  The idea is to only build once and to be transparent about building in
#  the Travis script.  To that end, some of the other builds that had been
#  done in this script will be moved into Travis.yml.
#

#
#  Determine default directories, etc., so we're not beholden to Travis
#  when running tests of the script during the development cycle.
#
openwhisk_cli_tag=${1:-"latest"}
scriptdir="$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )"

TRAVIS_BUILD_DIR="$( cd "${TRAVIS_BUILD_DIR:-$scriptdir/../..}" && pwd )"
export TRAVIS_BUILD_DIR

# For the gradle builds.
HOMEDIR="$(dirname "$TRAVIS_BUILD_DIR")"
OPENWHISK_HOME="$( cd "${OPENWHISK_HOME:-$HOMEDIR/openwhisk}" && pwd )"
export OPENWHISK_HOME

#
#  Run scancode using the ASF Release configuration
#
UTILDIR="$( cd "${UTILDIR:-$HOMEDIR/openwhisk-utilities}" && pwd )"
export UTILDIR
cd $UTILDIR
scancode/scanCode.py --config scancode/ASF-Release.cfg $TRAVIS_BUILD_DIR

#
#  Run Golint
#
cd $TRAVIS_BUILD_DIR
./gradlew --console=plain goLint

#
#  Run Unit and native tests
#
./gradlew --console=plain --info goTest -PgoTags=unit
./gradlew --console=plain --info goTest -PgoTags=native

#
#  Set up the OpenWhisk environment for integration testing
#

cd $OPENWHISK_HOME

#  Fire up the cluster
cd $OPENWHISK_HOME/ansible
ANSIBLE_CMD="ansible-playbook -i environments/local -e docker_image_prefix=openwhisk -e docker_image_tag=nightly"
$ANSIBLE_CMD setup.yml
$ANSIBLE_CMD prereq.yml
$ANSIBLE_CMD couchdb.yml
$ANSIBLE_CMD initdb.yml
$ANSIBLE_CMD apigateway.yml
$ANSIBLE_CMD wipe.yml
$ANSIBLE_CMD openwhisk.yml -e cli_tag=$openwhisk_cli_tag -e cli_installation_mode=local -e openwhisk_cli_home=$TRAVIS_BUILD_DIR -e controllerProtocolForSetup=http
$ANSIBLE_CMD properties.yml
$ANSIBLE_CMD routemgmt.yml

#  Run the test cases under openwhisk to ensure the quality of the runnint API.
cd $TRAVIS_BUILD_DIR
./gradlew --console=plain :tests:test --tests=*ApiGwCliTests*
sleep 30
./gradlew --console=plain :tests:test --tests=*ApiGwCliRoutemgmtActionTests*
sleep 30
./gradlew --console=plain :tests:test --tests=*ApiGwCliEndToEndTests*
sleep 30
./gradlew --console=plain :tests:test --tests=*Wsk*Tests*

#
#  Finally, run the integration test for the CLI
#
./gradlew --console=plain --info goTest -PgoTags=unit
./gradlew --console=plain --info goTest -PgoTags=integration
