#! /usr/bin/env bash

if [ -e "/usr/bin/vvsh" ]; then
    echo "Alread have vvsh installed";
    echo "If you want to update, please remove original vvsh in /usr/bin/vvsh"
    exit 0
fi

RAN=v1df323dvfdafasf
NODESHELL_PATH=/tmp/vvsh.$RAN

SYS=$(uname)

# TODO: consider, curl the zip file and unzip. If unzip not installed. Curl the binary.

if [ ! "$SYS" == "Linux" ] && [ ! "$SYS" == "Darwin" ]; then
    echo "System NOT supported. $SYS"
    echo "If you are window user. Please download the exe visiting:"
    echo "   https://github.com/lteam18/auto/releases/download/v0/vvsh.win.exe.zip"
    echo "If you are npm user, you could just install it by"
    echo "   npm install "
    echo "ns is the command that you could enter vvsh"
    exit 1
fi

unzip -v >> /dev/null

if [ $? -eq 0 ]; then
    if [ "Linux" == "$SYS"  ]; then
        curl -L https://github.com/lteam18/auto/releases/download/v0/vvsh.linux.exe.zip > $NODESHELL_PATH.zip
        unzip $NODESHELL_PATH.zip -d $NODESHELL_PATH.folder
        mv $NODESHELL_PATH.folder/vvsh.linux.exe $NODESHELL_PATH
        rm -rf $NODESHELL_PATH.folder
        rm $NODESHELL_PATH.zip
    else
        curl -L https://github.com/lteam18/auto/releases/download/v0/vvsh.mac.exe.zip > $NODESHELL_PATH.zip
        unzip $NODESHELL_PATH.zip -d $NODESHELL_PATH.folder
        mv $NODESHELL_PATH.folder/vvsh.mac.exe $NODESHELL_PATH
        rm -rf $NODESHELL_PATH.folder
        rm $NODESHELL_PATH.zip
    fi
else
    if [ "Linux" == "$SYS"  ]; then
        curl -L https://github.com/lteam18/auto/releases/download/v0/vvsh.linux.exe > $NODESHELL_PATH
    else
        # so, Darwin
        #if [ "Darwin" == "$SYS" ]; then
        curl -L https://github.com/lteam18/auto/releases/download/v0/vvsh.mac.exe > $NODESHELL_PATH
    fi
fi

chmod 555 $NODESHELL_PATH
mv $NODESHELL_PATH /usr/local/bin/vvsh
