@echo off
set /p VERSION=<buildcounter.txt
set /a VERSION=%VERSION%+1
>"buildcounter.txt" echo(%VERSION%
set GOARCH=amd64
set GOOS=linux
go build -ldflags "-X main.buildcount=%VERSION%" -o check_ciscoswitch
set GOOS=windows
go build -ldflags "-X main.buildcount=%VERSION%" -o check_ciscoswitch.exe
