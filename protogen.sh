#!/bin/bash
# Script FOr generating rpc equivalent code for proto DSL
protoc -I. --go_out=plugins=grpc:. proto/database/database.proto
mv github.com/oyamoh-brian/tv-service-database/database.pb.go proto/database
rm -rf github.com