#!/usr/bin/env bash
set -euo pipefail
#set -x
install_target=$(pwd)/tools/golang-dep
tempdir=""
clean_up() {
    rm -rf "${tempdir}" 
    exit -1
}
trap clean_up EXIT
platform=""
arch=""
version="v0.3.2"
[ "Linux" = "$(uname)" ] && platform="linux" || platform="darwin"
[ "x86_64" = "$(uname -m)" ] && arch="amd64" || arch="386"
[ "linux" = "${platform}" ] && tempdir=$(mktemp -d asdf-go-dep.XXXX) || tempdir=$(mktemp -dt go-dep)

mkdir -p ${install_target}
printf "Getting link for ${version}\n"
# This suffers from API limits which hurt when you have a hide behind address. So just hard coding it, sorry.
# download_url=$(curl --silent https://api.github.com/repos/golang/dep/releases/tags/${version}|grep -oE "\"browser_download_url\"\:\s?\"(.*)\""| cut -d "\"" -f 4 | grep "\-${platform}\-${arch}$")
download_url=https://github.com/golang/dep/releases/download/${version}/dep-${platform}-${arch}
printf "Getting binary from ${download_url}\n"
curl -L "${download_url}" -o "${tempdir}/dep"
chmod +x "${tempdir}/dep"
mv "${tempdir}/dep" ${install_target}
