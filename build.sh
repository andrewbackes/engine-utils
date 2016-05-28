#!/bin/bash
cd "$( dirname "$0" )"

programs=$(find . -type d | tail -n +2 | grep -v "git\|releases\|bin")

IFS=$'\n'
for prog in $programs; do
    pushd $prog
    
    echo "Linux Build..."
    export GOOS=linux 
    export GOARCH=amd64
    arch="linux-amd64"
    mkdir -p ../releases/$arch
    go build -o ../releases/$arch/$prog .
    
    echo "Mac Build..."
    export GOOS=darwin 
    export GOARCH=amd64
    arch="mac-amd64"
    mkdir -p ../releases/$arch
    go build -o ../releases/$arch/$prog .
    
    echo "Windows32 Build..."
    export GOOS=windows 
    export GOARCH=386
    arch="windows-386"
    mkdir -p ../releases/$arch
    go build -o ../releases/$arch/$prog.exe .
    
    echo "Windows64 Build..."
    export GOOS=windows 
    export GOARCH=amd64
    arch="windows-amd64"
    mkdir -p ../releases/$arch
    go build -o ../releases/$arch/$prog.exe .
    
    popd
done