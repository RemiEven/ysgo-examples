#!/bin/bash

set -e

# Clean and recreate dist folder
rm -rf dist
mkdir dist

# Copy static files to dist folder
cp $(go env GOROOT)/misc/wasm/wasm_exec.js dist/
cp index.html dist/
cp -r assets/files dist/assets

# Build wasm executable
GOOS=js GOARCH=wasm go build -o dist/game.wasm main.go game.go

# Create a zip archive
zip -r game_wasm.zip dist
