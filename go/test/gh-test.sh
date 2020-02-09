#! /usr/bin/env bash

SOURCE_DIR=$(dirname "${BASH_SOURCE[0]}")
cd "$SOURCE_DIR"
SOURCE_DIR=$(pwd)
cd -

export X_CMD_PATH=${X_CMD_PATH:-"x"}
eval "$(x bash/boot)"

@src std/@ std/assert

@run(){
    tput setaf 6
    tput bold
    echo ">>> $*"
    tput sgr0
    eval "$@"
    local CODE=$?
    if [ $CODE -eq 0 ]; then
        tput setaf 2
        tput bold
        echo "CODE: $?"
        echo
        tput sgr0
    else
        tput setaf 1
        tput bold
        echo "CODE: $?"
        echo
        tput sgr0
    fi
}

export GITHUB_USER="lteam-bot"
export GITHUB_TOKEN="cc1269732b75810a7f307df9149a394f3c35c462"

# export DEBUG=x:ghkv

CUR_FOLDER="$(pwd)"
TEST_WORKDIR="$CUR_FOLDER/abcdefg-test"
export VVKV_STORAGE_PATH="$TEST_WORKDIR/.x-cmd.com"

mkdir -p "$VVKV_STORAGE_PATH"; @defer 'rm -rf "$TEST_WORKDIR"'

cd "$TEST_WORKDIR" || exit 1; @defer 'cd "$PWD"'

@run x gh-repo-delete "$GITHUB_USER" "source.x-cmd.com"

@run x gh-setup "$GITHUB_USER" "$GITHUB_TOKEN"

@run x upload public @gh/hi/py "$SOURCE_DIR/testscripts/hi.py"  # "$CUR/testscripts/hi.py"

@run x ls @gh

@run x set-access private @gh/hi/py

@run x @gh/hi/py a b c

