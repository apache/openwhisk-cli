#!/usr/bin/env bash

HOMEDIR="$(dirname "$TRAVIS_BUILD_DIR")"
cd $HOMEDIR

sudo gpasswd -a travis docker
sudo -E bash -c 'echo '\''DOCKER_OPTS="-H tcp://0.0.0.0:4243 -H unix:///var/run/docker.sock --api-enable-cors --storage-driver=aufs"'\'' > /etc/default/docker'

# Docker
sudo apt-get -y update -qq
sudo apt-get -o Dpkg::Options::="--force-confold" --force-yes -y install docker-engine=1.12.0-0~trusty
sudo service docker restart
echo "Docker Version:"
docker version
echo "Docker Info:"
docker info

# Ansible
pip install --user ansible==2.1.2.0

# OpenWhisk stuff
git clone --depth 3 https://github.com/openwhisk/openwhisk.git

# Build script for Travis-CI.
WHISKDIR="$HOMEDIR/openwhisk"

ANSIBLE_CMD="ansible-playbook -i environments/local"

cd $WHISKDIR/ansible
$ANSIBLE_CMD setup.yml
$ANSIBLE_CMD prereq.yml
$ANSIBLE_CMD couchdb.yml
$ANSIBLE_CMD initdb.yml

cd $WHISKDIR
./gradlew distDocker

cd $WHISKDIR/ansible
$ANSIBLE_CMD wipe.yml
$ANSIBLE_CMD openwhisk.yml
