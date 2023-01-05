#!/usr/bin/env bash
go test ./... -count=1 -cover -p 4 -coverprofile=coverage.out
go tool cover -func coverage.out
go tool cover -html=coverage.out -o cx-installer.html
echo "Code-Coverage in Summe beträgt: $(go tool cover -func coverage.out | grep total | awk '{print $3}')"