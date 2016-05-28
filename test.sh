#!/bin/bash
cd "$( dirname "$0" )"

packages=$(find . -type d | tail -n +2 |grep -v "git\|bin")

IFS=$'\n'
for pkg in $packages; do
    pushd $pkg
    go test -coverprofile=.coverprofile -test.v -covermode=count || exit 1
    popd
done
echo "mode: count" > total.coverprofile
cat $(find . -name .coverprofile) | grep -v "mode: count" >> total.coverprofile