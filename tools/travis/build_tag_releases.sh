#!/usr/bin/env bash

arc=amd64
build_file_name=${1:-"wsk"}
zip_file_name=${2:-"OpenWhisk_CLI"}
os=$TRAVIS_OS_NAME

if [ "$os" == "osx" ]; then
    os="darwin"
fi
wsk=$build_file_name
if [ "$os" == "windows" ]; then
    wsk="$wsk.exe"
fi
cd $TRAVIS_BUILD_DIR
GOOS=$os GOARCH=$arc go build -o build/$os/$wsk
cd build/$os
zip -r "$TRAVIS_BUILD_DIR/$zip_file_name-$TRAVIS_TAG-$os-$arc.zip" $wsk