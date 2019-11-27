#!/bin/bash
SCRIPT_NAME="do"

if [ $# -ne 0  ]; then
        if [ "${1}" == "gen-pb" ]; then
                (set -x; protoc -I pkg/pb/ pkg/pb/*.proto --go_out=plugins=grpc:pkg/pb)

        elif [ "${1}" == "dummy" ]; then
                echo "dummy"
        fi
else
    echo "Usage: \"./${SCRIPT_NAME}.sh gen-pb\"";
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
