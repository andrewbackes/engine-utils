#!/bin/bash
cd "$( dirname "$0" )"

programs=$(find . -type d | tail -n +2 | grep -v "git\|bin")

IFS=$'\n'
mkdir bin
for prog in $programs; do
    pushd $prog
    go build -o ../bin/$prog .
    popd
done