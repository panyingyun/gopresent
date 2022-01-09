#!/bin/bash
CGO_ENABLED=0 GOOS=darwin  GOARCH=amd64 go build -o gopresent_mac
CGO_ENABLED=0 GOOS=linux  GOARCH=amd64 go build -o gopresent
CGO_ENABLED=0 GOOS=windows  GOARCH=amd64 go build -o gopresent.exe

