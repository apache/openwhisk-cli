#!/usr/bin/env bash

declare -a os_list=()
declare -a arc_list=("amd64" "386")
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
    for arc in "${arc_list[@]}"
    do
        wsk=$build_file_name
        os_name=$os
        if [ "$os" == "windows" ]; then
            wsk="$wsk.exe"
        fi
        if [ "$os" == "darwin" ]; then
            os_name="mac"
        fi
        cd $TRAVIS_BUILD_DIR
        GOOS=$os GOARCH=$arc go build -o build/$os/$arc/$wsk
        cd build/$os/$arc
        zip -r "$TRAVIS_BUILD_DIR/$zip_file_name-$TRAVIS_TAG-$os_name-$arc.zip" $wsk
    done
done
