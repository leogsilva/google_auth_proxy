#!/bin/bash

# build binary distributions for linux/amd64 and darwin/amd64

set -e 
export PATH=$PATH:$HOME/go/bin
export GOROOT=$HOME/go
export GOPATH=$HOME/google_auth_proxy
DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
echo "working dir $DIR"
mkdir -p $DIR/dist
mkdir -p $DIR/.godeps
export GOPATH=$DIR/.godeps:$GOPATH
gpm install

os=$(go env GOOS)
arch=$(go env GOARCH)
version=$(cat $DIR/version.go | grep "const VERSION" | awk '{print $NF}' | sed 's/"//g')
goversion=$(go version | awk '{print $3}')

echo "... running tests"
./test.sh || exit 1

for os in linux darwin; do
    echo "... building v$version for $os/$arch"
    BUILD=$(mktemp -d -t google_auth_proxy)
    TARGET="google_auth_proxy-$version.$os-$arch.$goversion"
    GOOS=$os GOARCH=$arch CGO_ENABLED=0 go build -o $BUILD/$TARGET/google_auth_proxy || exit 1
    pushd $BUILD
    tar czvf $TARGET.tar.gz $TARGET
    mv $TARGET.tar.gz $DIR/dist
    popd
done
