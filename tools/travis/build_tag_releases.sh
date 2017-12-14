#!/usr/bin/env bash

<<<<<<< HEAD
=======
#declare -a os_list=("linux" "windows" "darwin")
#declare -a arc_list=("amd64" "386")

>>>>>>> 44586c8d33c670c39cbbe2de9884f234f04e96eb
#  Currently supported combinations of OS and Architecture
declare -a builds=(
    "linux amd64"
    "linux 386"
    "linux s390x"
    "darwin amd64"
    "darwin 386"
    "windows amd64"
    "windows 386"
)

build_file_name=${1:-"wsk"}
zip_file_name=${2:-"OpenWhisk_CLI"}

for build in "${builds[@]}"
do
    # A little bash foo to tokenize the build string
    IFS=' ' read os arc <<< "${build}"
    wsk=$build_file_name
    os_name=$os
    if [ "$os" == "windows" ]; then
        wsk="$wsk.exe"
    fi
    if [ "$os" == "darwin" ]; then
        os_name="mac"
    fi
    cd $TRAVIS_BUILD_DIR || exit
    GOOS=$os GOARCH=$arc go build -ldflags "-X main.CLI_BUILD_TIME=`date -u '+%Y-%m-%dT%H:%M:%S%:z'`" -o build/$os/$arc/$wsk
    cd build/$os/$arc || exit
    if [[ "$os" == "linux" ]]; then
        tar -czvf "$TRAVIS_BUILD_DIR/$zip_file_name-$TRAVIS_TAG-$os_name-$arc.tgz" $wsk
    else
        zip -r "$TRAVIS_BUILD_DIR/$zip_file_name-$TRAVIS_TAG-$os_name-$arc.zip" $wsk
    fi
done
