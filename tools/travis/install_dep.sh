#!/usr/bin/env bash
set -euo pipefail
platform=""
arch=""
tempdir=""
version="v0.3.2"
[ "Linux" = "$(uname)" ] && platform="linux" || platform="darwin"
[ "x86_64" = "$(uname -m)" ] && arch="amd64" || arch="386"
[ "linux" = "${platform}" ] && tempdir=$(mktemp -d asdf-go-dep.XXXX) || tempdir=$(mktemp -dt go-dep)
mkdir -p $(pwd)/golang-dep
download_url=$(curl --silent https://api.github.com/repos/golang/dep/releases/tags/${version} |grep -oE "\"browser_download_url\"\:\s?\"(.*)\""| cut -d "\"" -f 4 | grep "\-${platform}\-${arch}$")
curl --silent -L "${download_url}" -o "${tempdir}/dep"
chmod +x "${tempdir}/dep"
mv "${tempdir}/dep" $(pwd)/golang-dep
rm -rf "${tempdir}"
