#!/usr/bin/env bash
test -d gopath && chmod -R 777 gopath 
rm -rf go.mod go.sum main gopath
