#!/usr/bin/env bash

set -e

#
#  At this point, the Travis build should already have built the binaries and
#  the release.  If you're running manually, this command should get you to
#  the same place:
#
#    ./gradlew releaseBinaries
#
#  Also at this point, you should already have incubator-openwhisk pulled down
#  from gradle in the parent directory, using a command such as:
#
#    git clone --depth 3 https://github.com/apache/incubator-openwhisk.git
#
#  To be clear, your directory structure will look something like...
#
#      $HOMEDIR
#       |- incubator-openwhisk
#       |- incubator-openwhisk-cli (This project)
#       |- incubator-openwhisk-utilities (For scancode)
#
#  The idea is to only build once and to be transparent about building in
#  the Travis script.  To that end, some of the other builds that had been
#  done in this script will be moved into Travis.yml.
#

#
#  Determine default directories, etc., so we're not beholden to Travis
#  when running tests of the script during the development cycle.
#
scriptdir="$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )"

TRAVIS_BUILD_DIR="$( cd "${TRAVIS_BUILD_DIR:-$scriptdir/../..}" && pwd )"
export TRAVIS_BUILD_DIR

# For the gradle builds.
HOMEDIR="$(dirname "$TRAVIS_BUILD_DIR")"
OPENWHISK_HOME="$( cd "${OPENWHISK_HOME:-$HOMEDIR/incubator-openwhisk}" && pwd )"
export OPENWHISK_HOME

#
#  Perform code validation using scanCode and Golint
#
../incubator-openwhisk-utilities/scancode/scanCode.py $TRAVIS_BUILD_DIR
./gradlew --console=plain goLint

#
#  Run Unit and native tests
#
./gradlew --console=plain --info goTest -PgoTags=unit,native

#
#  Set up the OpenWhisk environment for integration testing
#

#  Build docker images
cd $OPENWHISK_HOME
./gradlew --console=plain distDocker -PdockerImagePrefix=testing

#  Fire up the cluster
cd $OPENWHISK_HOME/ansible
ANSIBLE_CMD="ansible-playbook -i environments/local -e docker_image_prefix=testing"
$ANSIBLE_CMD setup.yml
$ANSIBLE_CMD prereq.yml
$ANSIBLE_CMD couchdb.yml
$ANSIBLE_CMD initdb.yml
$ANSIBLE_CMD apigateway.yml
$ANSIBLE_CMD wipe.yml
$ANSIBLE_CMD openwhisk.yml -e cli_installation_mode=local -e openwhisk_cli_home=$TRAVIS_BUILD_DIR -e controllerProtocolForSetup=http

#  Run the test cases under openwhisk to ensure the quality of the runnint API.
cd $TRAVIS_BUILD_DIR
./gradlew --console=plain :tests:test -Dtest.single=*ApiGwCliTests*
sleep 30
./gradlew --console=plain :tests:test -Dtest.single=*ApiGwCliRoutemgmtActionTests*
sleep 30
./gradlew --console=plain :tests:test -Dtest.single=*ApiGwCliEndToEndTests*
sleep 30
./gradlew --console=plain :tests:test -Dtest.single=*Wsk*Tests*

#
#  Finally, run the integration test for the CLI
#
./gradlew --console=plain --info goTest -PgoTags=integration
