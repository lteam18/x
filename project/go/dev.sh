export PATH=$PATH:`pwd`/dist

install(){
    GOPATH="`pwd`" go get -d ./src/*
}

build(){
    GOPATH="`pwd`" go build -o build/xmain src/xmain/*.go
}

# https://blog.filippo.io/shrink-your-go-binaries-with-this-one-weird-trick/

dist(){
    GOPATH="`pwd`" go build -o dist/xmain src/xmain/*.go
    # GOPATH="`pwd`" go build -o dist/xmain -ldflags="-s -w" src/xmain/*.go
}

dist.linux(){
    GOPATH="`pwd`" go build -o dist/xmain -ldflags="-s -w" src/xmain/*.go
}

dist.arch(){
    rm -rf src/github.com src/golang.org
    GOOS=${1:?"GOOS?"}
    GOARCH=${2:?"GOARCH?"}
    echo Building $GOOS-$GOARCH
    GOPATH=`pwd` go get -d ./src/*
    local OUTPUT=x-installer.$GOOS-$GOARCH
    if [ "windows" == "$GOOS" ]; then
        OUTPUT=x-installer.$GOOS-$GOARCH.exe
    fi
    GOPATH=`pwd` go build -o ./dist/$OUTPUT ./src/xmain/*.go
}

dist.all(){
    echo "dist.all with GOPATH: $GOPATH"
    dist.arch darwin amd64
    dist.arch windows amd64
    dist.arch linux arm
    dist.arch linux arm64
    dist.arch linux amd64 # Must be the last one. For travis build test
}

dist.all.mac(){
    GOPATH="`pwd`" go build -o dist/xmain.mac -ldflags="-s -w" src/xmain/*.go
    GOPATH="`pwd`" CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o dist/xmain.linux -ldflags="-s -w" src/xmain/*.go
    # CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -o dist/xmain -ldflags="-s -w" src/xmain/*.go
}

timeall(){
    ls -alh ./dist
    time ./dist/hi
    time ./dist/x hi
}

[ -n "$1" ] && eval "$1"
