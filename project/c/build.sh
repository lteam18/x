build(){
    gcc vvx.c -o vvx
}

test(){
    time gcc vvx.c -o vvx.test
    echo "Run"
    time A=3 ./vvx.test bash test.bash
}

dist.mac(){
    mkdir -p dist
    # gcc vvx.c -O3 -o dist/vvx.mac
    gcc vvx.c -o dist/vvx.mac
}

dist.linux(){
    mkdir -p dist
    gcc vvx.c -o dist/vvx.linux
}

dist.linux-docker(){
    mkdir -p dist
    docker run -v `pwd`:/code -it gcc gcc /code/vvx.c -o /code/dist/vvx.linux
}

dist.windows(){
    mkdir -p dist
    gcc vvx.c -o dist/vvx.windows
}
