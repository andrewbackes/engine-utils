#!/bin/bash


cd "$( dirname "$0" )"

programs=$(find . -type d | tail -n +2 | grep -v "git\|releases\|bin")
name="engine-utils"

IFS=$'\n'
for prog in $programs; do
    pushd $prog
    ls *.go >/dev/null
    if [ "$?" -eq  "0" ]; then
        echo "Linux Build..."
        export GOOS=linux 
        export GOARCH=amd64
        arch="linux-amd64"
        mkdir -p ../releases/${arch}
        go build -o ../releases/${arch}/$prog . || exit 1
        
        echo "Mac Build..."
        export GOOS=darwin 
        export GOARCH=amd64
        arch="mac-amd64"
        mkdir -p ../releases/${arch}
        go build -o ../releases/${arch}/$prog . || exit 1
        
        echo "Windows32 Build..."
        export GOOS=windows 
        export GOARCH=386
        arch="windows-386"
        mkdir -p ../releases/${arch}
        go build -o ../releases/${arch}/${prog}.exe . || exit 1
        
        echo "Windows64 Build..."
        export GOOS=windows 
        export GOARCH=amd64
        arch="windows-amd64"
        mkdir -p ../releases/${arch}
        go build -o ../releases/${arch}/${prog}.exe . || exit 1
    fi
    popd
done

echo "Adding EPDs"
ls -d releases/*/ | xargs -n 1 cp -R testsuite/epds || exit 1

cd releases
echo "Compressing"
zip -r "${name}-win64.zip" "windows-amd64" || exit 1
zip -r "${name}-win386.zip" "windows-386" || exit 1
tar -zcvf "${name}-mac64.tar.gz" "mac-amd64" || exit 1
tar -zcvf "${name}-linux64.tar.gz" "linux-amd64" || exit 1