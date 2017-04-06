#!/usr/bin/env bash

declare -a os_list=("linux" "darwin" "windows")
arc=amd64
build_file_name=${1:-"wsk"}
zip_file_name=${2:-"OpenWhisk_CLI"}

for os in "${os_list[@]}"
do
    wsk=$build_file_name
    if [ "$os" == "windows" ]; then
        wsk="$wsk.exe"
    fi
    cd $TRAVIS_BUILD_DIR
    GOOS=$os GOARCH=$arc go build -o build/$os/$wsk
    cd build/$os
    zip -r "$TRAVIS_BUILD_DIR/$zip_file_name-$TRAVIS_TAG-$os-$arc.zip" $wsk
done