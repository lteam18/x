# shellcheck shell=bash

# shellcheck disable=SC2155
export PATH=$PATH:$(pwd)/dist

install(){
    GOPATH=$(pwd) go get -d ./src/*
}

build(){
    # GOPATH="`pwd`" go build -o build/xmain src/xmain/*.go
    GOPATH=$(pwd) go build -o build/xmain  -ldflags="-s -w" src/xmain/*.go
}

# https://blog.filippo.io/shrink-your-go-binaries-with-this-one-weird-trick/

dist(){
    # GOPATH="`pwd`" go build -o dist/xmain src/xmain/*.go
    GOPATH=$(pwd) go build -o dist/xmain -ldflags="-s -w" src/xmain/*.go
}

dist.linux(){
    GOPATH=$(pwd) go build -o dist/xmain -ldflags="-s -w" src/xmain/*.go
}

dist.arch(){
    rm -rf src/github.com
    rm -rf src/golang.org
    export GOPATH=$(pwd)
    export GOOS=${1:?"GOOS?"}
    export GOARCH=${2:?"GOARCH?"}
    echo "---------------"
    echo Building "$GOOS-$GOARCH" "$GOPATH"
    go get -d ./src/*
    local OUTPUT=x-installer.$GOOS-$GOARCH
    if [ "windows" == "$GOOS" ]; then
        OUTPUT=x-installer.$GOOS-$GOARCH.exe
    fi
    go build -ldflags="-s -w" -o "./dist/$OUTPUT" ./src/xmain/*.go
    upx "./dist/$OUTPUT"
}

dist.all(){
    command -v upx || (
        apt update && apt install upx -y
    )
    echo "dist.all with GOPATH: $GOPATH"
    dist.arch windows amd64
    dist.arch linux arm
    dist.arch linux arm64
    dist.arch linux amd64 # Must be the last one. For travis build test
    dist.arch darwin amd64
}

dist.all.mac(){
    GOPATH=$(pwd) go build -o dist/xmain.mac -ldflags="-s -w" src/xmain/*.go
    GOPATH=$(pwd) CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o dist/xmain.linux -ldflags="-s -w" src/xmain/*.go
    # CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -o dist/xmain -ldflags="-s -w" src/xmain/*.go
}

timeall(){
    ls -alh ./dist
    time ./dist/hi
    time ./dist/x hi
}

[ -n "$1" ] && eval "$1"
