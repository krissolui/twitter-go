#!/bin/bash

print_helper_menu() {
    echo "$0 [<options...>] <file path> <output file>"
    echo "Options:"
    echo "  -a : linux, arm64"
    echo "       Build architecture."
    echo "  -o : yes / no"
    echo "       Override output file if exists"
}

override=''
arch=''

while getopts "a:ho" flag; do
    case "${flag}" in
        a) 
            arch="${OPTARG}"
            ;;
        h) 
            print_helper_menu
            exit 0
            ;;
        o) 
            override="yes"
            ;;
        *) 
            print_helper_menu
            exit 1 ;;
    esac
done

entry=${@:$OPTIND:1}
output=${@:$OPTIND+1:1}

if [[ $(($# - $OPTIND + 1)) -lt 2 ]];
then
    echo "required at least 2 arguments!"
    print_helper_menu
    exit 1
fi

# check file path exists
if [ ! -d "$entry" ];
then
    echo "$entry directory not exists!"
    exit 1
fi

# check output file not exists
if [ -f "$output" ];
then
    if [ "$override" == '' ];
    then
        read -p "$output file already exists! Override file? (y/n): " override
    fi

    case $override in
        "y" | "Y" | "yes" | "YES" | "Yes")
            ;;
        *)
            exit 0
            ;;
    esac
fi

# get machine architecture if not passed in
if [[ $arch == '' ]];
then
    arch="$(uname | tr '[:upper:]' '[:lower:]')-$(arch)"
fi

case $arch in
    "linux")
        env GOOS=linux CGO_ENABLED=0 go build -o $output $entry
        ;;
    "amd64" | "mac" | "darwin" | "mac-amd" | "darwin-amd" | "mac-amd64" | "darwin-amd64")
        env GOOS=darwin GOARCH=amd64 CGO_ENABLED=0 go build -o $output $entry
        ;;
    "m1" | "mac-m1" | "darwin-m1" | "mac-arm" |  "darwin-arm" | "mac-arm64" | "darwin-arm64")
        env GOOS=darwin GOARCH=arm64 CGO_ENABLED=0 go build -o $output $entry
        ;;
    *)
        go build -o $output $entry
        ;;
esac

echo "built $output for $arch"