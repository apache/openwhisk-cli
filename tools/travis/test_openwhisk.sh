#!/usr/bin/env bash

set -e

export OPENWHISK_HOME="$(dirname "$TRAVIS_BUILD_DIR")/incubator-openwhisk";
HOMEDIR="$(dirname "$TRAVIS_BUILD_DIR")"
cd $HOMEDIR

# Clone the OpenWhisk code
git clone --depth 3 https://github.com/apache/incubator-openwhisk.git

# Build script for Travis-CI.
WHISKDIR="$HOMEDIR/incubator-openwhisk"

cd $WHISKDIR
./tools/travis/setup.sh

ANSIBLE_CMD="ansible-playbook -i environments/local -e docker_image_prefix=testing"

cd $WHISKDIR/ansible
$ANSIBLE_CMD setup.yml
$ANSIBLE_CMD prereq.yml
$ANSIBLE_CMD couchdb.yml
$ANSIBLE_CMD initdb.yml
$ANSIBLE_CMD apigateway.yml

cd $WHISKDIR
GRADLE_PROJS_SKIP="-x :core:pythonAction:distDocker -x :core:python2Action:distDocker -x :core:swift3Action:distDocker -x :core:javaAction:distDocker"
TERM=dumb ./gradlew distDocker -PdockerImagePrefix=testing $GRADLE_PROJS_SKIP

cd $WHISKDIR/ansible
$ANSIBLE_CMD wipe.yml
$ANSIBLE_CMD openwhisk.yml

# Copy the binary generated into the OPENWHISK_HOME/bin, so that the test cases will run based on it.
mkdir -p $WHISKDIR/bin
cp $TRAVIS_BUILD_DIR/wsk $WHISKDIR/bin

# Run the test cases under openwhisk to ensure the quality of the binary.
cd $TRAVIS_BUILD_DIR
./gradlew :tests:test -Dtest.single=Wsk*Tests*
./gradlew tests:test -Dtest.single=*ApiGwRoutemgmtActionTests*
sleep 30
./gradlew tests:test -Dtest.single=*ApiGwTests*
sleep 30
./gradlew tests:test -Dtest.single=*ApiGwEndToEndTests*
sleep 30
make integration_test
