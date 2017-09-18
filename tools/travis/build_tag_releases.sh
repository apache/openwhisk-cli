#!/usr/bin/env bash

declare -a os_list=("linux" "windows" "darwin")
declare -a arc_list=("amd64" "386")
build_file_name=${1:-"wsk"}
zip_file_name=${2:-"OpenWhisk_CLI"}

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
        GOOS=$os GOARCH=$arc go build -ldflags "-X main.CLI_BUILD_TIME=`date -u '+%Y-%m-%dT%H:%M:%S%:z'`" -o build/$os/$arc/$wsk
        cd build/$os/$arc
        if [[ "$os" == "linux" ]]; then
            tar -czvf "$TRAVIS_BUILD_DIR/$zip_file_name-$TRAVIS_TAG-$os_name-$arc.tgz" $wsk
        else
            zip -r "$TRAVIS_BUILD_DIR/$zip_file_name-$TRAVIS_TAG-$os_name-$arc.zip" $wsk
        fi
    done
done
