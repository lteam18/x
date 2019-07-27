export PATH=$PATH:`pwd`/dist

install(){
    GOPATH="`pwd`" go get -d ./src/*
}

build(){
    GOPATH="`pwd`" go build -o build/xmain src/xmain/*.go
}

# https://blog.filippo.io/shrink-your-go-binaries-with-this-one-weird-trick/

dist(){
    GOPATH="`pwd`" go build -o dist/xmain -ldflags="-s -w" src/xmain/*.go
}

dist.linux(){
    GOPATH="`pwd`" go build -o dist/xmain -ldflags="-s -w" src/xmain/*.go
}

distall(){
    GOPATH="`pwd`" go build -o dist/hi -ldflags="-s -w" src/hi/*.go
    GOPATH="`pwd`" go build -o dist/x -ldflags="-s -w" src/x/*.go
    GOPATH="`pwd`" go build -o dist/xmain -ldflags="-s -w" src/xmain/*.go
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
