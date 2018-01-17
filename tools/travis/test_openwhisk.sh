#!/usr/bin/env bash

set -e

#
#  At this point, the Travis build should already have built the binaries and
#  the release.  If you're running manually, this command should get you to
#  the same place:
#
#    ./gradlew buildBinaries release --PcrossCompile=true
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
travis_build_dir="$( cd "${TRAVIS_BUILD_DIR:-$scriptdir/../..}" && pwd )"

#
#  These are the basic tests
#
cd $travis_build_dir
./tools/travis/scancode.sh

#  Run separate test scopes as separate
./gradlew --console=plain goLint
./gradlew --console=plain goTest -PgoTags=unit
export PATH=$PATH:$travis_build_dir
./gradlew --console=plain goTest -PgoTags=native

HOMEDIR="$(dirname "$travis_build_dir")"
OPENWHISK_HOME="$HOMEDIR/incubator-openwhisk"
export OPENWHISK_HOME

# Clone the OpenWhisk code
# TODO: Probably not necessary because the build now requires the
#       incubator-openwhisk directory as well.
if [ ! -e "$OPENWHISK_HOME" ]; then
    cd $HOMEDIR
    git clone --depth 3 https://github.com/apache/incubator-openwhisk.git
fi

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

#  A quick nested build -- probably not needed.
#( cd $travis_build_dir && ./gradlew --console=plain buildBinaries )

$ANSIBLE_CMD wipe.yml
# TODO -- Some flag might be needed to get CLI from local directory (?)
$ANSIBLE_CMD openwhisk.yml -e openwhisk_cli_home=$travis_build_dir

# Copy the binary generated into the OPENWHISK_HOME/bin, so that the test cases will run based on it.
# TODO - if the ansible above were correctly configured, this wouldn't be necessary
mkdir -p $OPENWHISK_HOME/bin
cp -f $travis_build_dir/bin/wsk $OPENWHISK_HOME/bin

# Run the test cases under openwhisk to ensure the quality of the binary.
cd $travis_build_dir

./gradlew --console=plain :tests:test -Dtest.single=*ApiGwCliTests*
#sleep 30
./gradlew --console=plain :tests:test -Dtest.single=*ApiGwCliRoutemgmtActionTests*
#sleep 30
./gradlew --console=plain :tests:test -Dtest.single=*ApiGwCliEndToEndTests*
#sleep 30
./gradlew --console=plain :tests:test -Dtest.single=*Wsk*Tests*

./gradlew --console=plain goTest -PgoTags=integration
