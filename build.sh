#!/usr/bin/env bash
version=`cat buildcounter.txt`
version=$((version+1))
echo "$version" > buildcounter.txt
go clean -cache -testcache -modcache ; go mod vendor &> /dev/null
GOOS=windows GOARCH=amd64 go build -ldflags "-X main.buildcount=%VERSION%" -o check_ciscoswitch.exe
GOOS=linux GOARCH=amd64 go build -ldflags "-X main.buildcount=`echo $version`" -o check_ciscoswitch
