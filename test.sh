#!/bin/bash
cd "$( dirname "$0" )"

packages=$(find . -type d | tail -n +2 | grep -v "git\|releases\|bin")

IFS=$'\n'
for pkg in $packages; do
    pushd $pkg
    ls *.go > /dev/null
    if [ "$?" -eq  "0" ]; then
        go test -coverprofile=.coverprofile -test.v -covermode=count || exit 1
    fi
    popd
done
echo "mode: count" > total.coverprofile
cat $(find . -name .coverprofile) | grep -v "mode: count" >> total.coverprofile