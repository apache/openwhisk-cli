#!/usr/bin/env bash

set -e

#
#  At this point, the Travis build should already have built the binaries and
#  the release.  If you're running manually, this command should get you to
#  the same place:
#
#    ./gradlew buildBinaries release --PcrossCompile=true
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
#  Figure out default directories, etc., so we're not beholden to Travis
#  when running tests of the script.
#
scriptdir="$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )"

TRAVIS_BUILD_DIR="$( cd "${TRAVIS_BUILD_DIR:-$scriptdir/../..}" && pwd )"
export TRAVIS_BUILD_DIR

# For the gradle builds.
HOMEDIR="$(dirname "$TRAVIS_BUILD_DIR")"
OPENWHISK_HOME="$( cd "${OPENWHISK_HOME:-$HOMEDIR/incubator-openwhisk}" && pwd )"
export OPENWHISK_HOME

#
#  These are the basic tests
#
../incubator-openwhisk-utilities/scancode/scanCode.py $TRAVIS_BUILD_DIR

#  Run separate test scopes as separate
./gradlew --console=plain goLint
./gradlew --console=plain goTest -PgoTags=unit
export PATH=$PATH:$TRAVIS_BUILD_DIR
./gradlew --console=plain goTest -PgoTags=native

#
#  Set up the OpenWhisk environment ( TODO: reusable script for incubtor-openwhisk? )
cd $OPENWHISK_HOME
./tools/travis/setup.sh

ANSIBLE_CMD="ansible-playbook -i environments/local -e docker_image_prefix=testing"
./gradlew --console=plain distDocker -PdockerImagePrefix=testing

cd $OPENWHISK_HOME/ansible
$ANSIBLE_CMD setup.yml
$ANSIBLE_CMD prereq.yml
$ANSIBLE_CMD couchdb.yml
$ANSIBLE_CMD initdb.yml
$ANSIBLE_CMD apigateway.yml
$ANSIBLE_CMD wipe.yml
# TODO -- Some flag might be needed to get CLI from local directory (?)
$ANSIBLE_CMD openwhisk.yml -e openwhisk_cli_home=$TRAVIS_BUILD_DIR

# Copy the binary generated into the OPENWHISK_HOME/bin, so that the test cases will run based on it.
# TODO - if the ansible above were correctly configured, this wouldn't be necessary
mkdir -p $OPENWHISK_HOME/bin
cp -f $TRAVIS_BUILD_DIR/bin/wsk $OPENWHISK_HOME/bin

# Run the test cases under openwhisk to ensure the quality of the binary.
cd $TRAVIS_BUILD_DIR

./gradlew --console=plain :tests:test -Dtest.single=*ApiGwCliTests*
#sleep 30
./gradlew --console=plain :tests:test -Dtest.single=*ApiGwCliRoutemgmtActionTests*
#sleep 30
./gradlew --console=plain :tests:test -Dtest.single=*ApiGwCliEndToEndTests*
#sleep 30
./gradlew --console=plain :tests:test -Dtest.single=*Wsk*Tests*

./gradlew --console=plain goTest -PgoTags=integration
