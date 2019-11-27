#!/bin/bash
SCRIPT_NAME="do"

if [ $# -ne 0  ]; then
        flag="${1}"
        if [ "$flag" == "gen-pb" ]; then
                (set -x; protoc -I pkg/pb/ pkg/pb/*.proto --go_out=plugins=grpc:pkg/pb)
        elif [ "$flag" == "run" ]; then
                (set -x; go run cmd/main.go)
        elif [ "$flag" == "dummy" ]; then
                echo "dummy"
        fi
else
    echo "Usage: \"./${SCRIPT_NAME} gen-pb\"";
    exit 1
fi

#case $opt in
#while getopts ":a:" opt; do
#a)
#echo "-a was triggered, Parameter: $OPTARG" >&2
#;;
#\?)
        #echo "Invalid option: -$OPTARG" >&2
        #exit 1
        #;;
#:)
        #echo "Option -$OPTARG requires an argument." >&2
        #exit 1
        #;;
        #esac
        #done
