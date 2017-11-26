#!/bin/bash

env GOOS="windows" GOARCH="386" go build -v -o bin/winprivcheck-x86.exe
env GOOS="windows" GOARCH="amd64" go build -v -o bin/winprivcheck-x64.exe
