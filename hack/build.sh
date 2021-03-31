#!/bin/sh
set -euo pipefail

GOOS=darwin GOARCH=amd64 go build -o prify-darwin-amd64
GOOS=linux GOARCH=amd64 go build -o prify-linux-amd64
GOOS=linux GOARCH=arm64 go build -o prify-linux-arm64

