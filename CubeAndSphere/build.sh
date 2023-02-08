#!/bin/bash

set -e

cp $(go env GOROOT)/misc/wasm/wasm_exec.js public/
GOOS=js GOARCH=wasm go build -o public/game.wasm main.go
zip -r cubeAndSphere.zip public
