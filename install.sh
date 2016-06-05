#!/usr/bin/sh
go build httpserver.go
rm /data/website/homelink/httpserver -rf
cp httpserver /data/website/homelink/

