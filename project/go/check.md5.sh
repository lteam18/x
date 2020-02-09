#! /usr/bin/env bash

md5t(){
    if md5sum --help 1>/dev/null 2>&1; then
        echo md5sum
        return 0
    fi
    if md5 -s "123" 1>/dev/null 2>&1; then
        echo md5
        return 0
    fi
    # TODO: Download md5 program from web
    return 1
}

MD5=$(md5t)
if [ -z "$MD5" ]; then
    echo "md5 program is not available." >&2
    exit 1
fi

TMP_FILEPATH=".x-cmd.com.tmp.release.txt"

curl https://x-cmd.github.io/release.txt >$TMP_FILEPATH

check(){
    grep "$($(md5t) "${1:?Please Provie filepath}")" tmp.release.txt
}

x_path=$(command -v x)
if [ -z "$x_path" ]; then
    echo x does not exists
else
    if check "$x_path"; then
        echo "$x_path" is valid 
    else
        echo WARNING: md5 of "$x_path" is Invalid
    fi
fi

rm $TMP_FILEPATH


