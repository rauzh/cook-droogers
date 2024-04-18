#!/bin/bash

./clean_mocks.sh

cd backend/

go generate ./...
go test ./... -v

cd ../

./clean_mocks.sh