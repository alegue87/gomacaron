#!/usr/bin/bash

source conf.sh

[[ "$TRACE" ]] && set -x
OWD=`pwd`
pushd `dirname "$0"` > /dev/null
trap __EXIT EXIT

colorful=false
tput setaf 7 > /dev/null 2>&1
if [[ $? -eq 0 ]]; then
    colorful=true
fi

function __EXIT() {
    popd > /dev/null
}

function printError() {
    $colorful && tput setaf 1
    >&2 echo "Error: $@"
    $colorful && tput setaf 7
}

function printImportantMessage() {
    $colorful && tput setaf 3
    >&2 echo "$@"
    $colorful && tput setaf 7
}

function printUsage() {
    $colorful && tput setaf 3
    >&2 echo "$@"
    $colorful && tput setaf 7
}

xOS="linux"
if [[ $OSTYPE == darwin* ]]; then
    xOS="darwin"
fi

PROGRAM="main"
PLUGIN1="plugin1"
PLUGIN2="plugin2"
PROGRAM_BUILD_OUTPUT_DIR="bin/$xOS/$PROGRAM"
PROGRAM_EXE="bin/$xOS/$PROGRAM"

printf "Building $PROGRAM...\n"
echo

CGO_ENABLED=1 GOARCH=amd64 go build -trimpath -o "$PROGRAM_BUILD_OUTPUT_DIR" -buildvcs=false
[[ $? -ne 0 ]] && exit 1

plugin/$PLUGIN1/build.sh
[[ $? -ne 0 ]] && exit 1
echo

plugin/$PLUGIN2/build.sh
[[ $? -ne 0 ]] && exit 1
echo

printf "Starting $PROGRAM...\n\n"
"$PROGRAM_EXE" --pluginDir="bin/$xOS/plugin/$PROGRAM" --pidFile="bin/$xOS/$PROGRAM.pid"
