#! /usr/bin/env bash

if [ -e "/usr/local/bin/x" ]; then
    echo "Alread have X installed";
    echo "If you want to update, please remove original X in /usr/local/bin/x"
    exit 0
fi

RAN=v1df323dvfdafasf
NODESHELL_PATH=/tmp/vvsh.$RAN

SYS=$(uname)

# TODO: consider, curl the zip file and unzip. If unzip not installed. Curl the binary.

if [ ! "$SYS" == "Linux" ] && [ ! "$SYS" == "Darwin" ]; then
    echo "System NOT supported. $SYS"
    echo "If you are window user. Please download the exe visiting:"
    echo "   https://github.com/lteam18/x/releases"
    echo "x is the command placed in /usr/local/bin/x"
    exit 1
fi

if [ "Linux" == "$SYS"  ]; then
    curl -L https://github.com/lteam18/x/releases/download/latest-dev/x-installer.linux-amd64 > $NODESHELL_PATH
else
    # so, Darwin
    #if [ "Darwin" == "$SYS" ]; then
    curl -L https://github.com/lteam18/x/releases/download/latest-dev/x-installer.darwin-amd64 > $NODESHELL_PATH
fi

chmod +x $NODESHELL_PATH
$NODESHELL_PATH install
rm $NODESHELL_PATH
# mv $NODESHELL_PATH /usr/local/bin/vvsh
