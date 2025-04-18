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

PLUGIN1="plugin1"
PLUGIN2="plugin2"

echo "Remove olds builds.."
rm -r bin/linux/plugin/main/tmp/*

plugin/$PLUGIN1/build.sh
[[ $? -ne 0 ]] && exit 1
echo

plugin/$PLUGIN2/build.sh
[[ $? -ne 0 ]] && exit 1
echo

echo "Sending signal..."
kill -USR1 `cat "bin/$xOS/main.pid"`
# send signal for reload to application
