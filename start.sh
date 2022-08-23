#!/bin/bash

GOOS=`go env GOOS`
GOARCH=`go env GOARCH`

nohup ./bin/$GOARCH/$GOOS/cbshare client --addr=$1 --log-file=./log/cbshare.log 2>&1>/dev/null &