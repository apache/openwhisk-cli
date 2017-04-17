#!/usr/bin/env bash

declare -a os_list=()
arc=amd64
build_file_name=${1:-"wsk"}
zip_file_name=${2:-"OpenWhisk_CLI"}
os=$TRAVIS_OS_NAME

if [[ $TRAVIS_OS_NAME == 'linux' ]]; then
    # Currently we have not set up the CI designated to build windows binaries, so we tentatively
    # add the windows build into the linux CI environment.
    os_list=("linux" "windows")
elif [[ $TRAVIS_OS_NAME == 'osx' ]]; then
    os_list=("darwin")
fi

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
