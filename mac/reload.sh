#!/usr/bin/bash

[[ "$TRACE" ]] && set -x
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

PLUGIN="world"
PLUGIN1="world1"

echo "Remove olds builds.."
rm -r bin/linux/plugin/hello/tmp/*

plugin/$PLUGIN/build.sh
[[ $? -ne 0 ]] && exit 1
echo

plugin/$PLUGIN1/build.sh
[[ $? -ne 0 ]] && exit 1
echo

echo "Sending signal..."
kill -USR1 `cat "bin/$xOS/hello.pid"`
# send signal for reload to application
