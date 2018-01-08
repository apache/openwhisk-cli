#!/bin/bash

set +x
set -e
set -f

get_bin_name () {
  local os=$1
  local bin="wsk"

  if [ $os = "windows" ]; then
    bin="${bin}.exe";
  fi

  echo $bin;
};

build_cli () {
  local os=$1
  local arch=$2
  local bin=$3

  echo "Building for OS '$os' and architecture '$arch'"

  if [ $os = "mac" ]; then
    export GOOS=darwin;
  else
    export GOOS=$os;
  fi

  export GOARCH=$arch

  cd /src/github.com/apache/incubator-openwhisk-cli
  go build -ldflags "-X main.CLI_BUILD_TIME=`date -u '+%Y-%m-%dT%H:%M:%S%:z'`" \
      -v -o build/$os/$arch/$bin main.go;
};

get_compressed_name() {
  local os=$1
  local arch=$2
  local product_name="OpenWhisk_CLI"

  if [ $arch = amd64 ]; then
      comp_name="$product_name-$os";
  elif [ $arch = 386 ]; then
      comp_name="$product_name-$os-32bit";
  else
      comp_name="$product_name-$os-$arch";
  fi

  echo $comp_name;
};

compress_binary() {
    local comp_name=$1
    local bin=$2
    local os=$3
    local arch=$4

    cd build/$os/$arch

    if [ $os = "linux" ]; then
      comp_name="$comp_name.tgz"
      tar -cvzf $comp_name $bin >/dev/null 2>&1;
    else
      comp_name="$comp_name.zip"
      zip $comp_name $bin >/dev/null 2>&1;
    fi

    cd ../../..
    echo $os/$arch/$comp_name;
};

create_cli_packages() {
  local dirIndex='{"cli":{'

  IFS="," echo "$platforms" | tr " " "\n" | sort | while read -r platform arch
  do
      #  Control-break processing for platform changes
      if [ "$platform" != "$old_platform" ]; then
        if [ "$old_platform" != "" ]; then
          dirIndex="${dirIndex/%','/'},'/}" # Replace trailing comma with end-brace comma
        fi
        dirIndex+='"$platform":{'
      fi
      old_platform=$platform

      bin=$(get_bin_name $platform)
      build_cli $platform $arch $bin
      comp_name=$(get_compressed_name $platform $arch)
      comp_path=$(compress_binary $comp_name $bin $platform $arch)

      if [ $arch = $default_arch ]; then
          dirIndex+="\"default\":{\"path\":\"${comp_path}\"},";
      fi
      dirIndex+="\"${arch}\":{\"path\":\"${comp_path}\"},";
  done

  dirIndex="${dirIndex/%','/'}}}'/}"   # Replace trailing comma with end-braces
  echo $dirIndex > ./build/content.json
};

platforms="${CLI_OS_ARCH}"  # In format 'linux,386 linux,amd64 linux,s390x mac,386 ...'
default_arch="amd64"
create_cli_packages
