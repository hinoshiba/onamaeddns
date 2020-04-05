#!/bin/sh
export GOPATH
GOPATH="`pwd`"
cd src/onamaeddns
dep ensure
dep status
