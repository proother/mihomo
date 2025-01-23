#!/bin/bash

# Exit on error
set -e

# Set up environment variables
export CGO_ENABLED=1
export GOOS=darwin
export GOARCH=arm64
export CC=$(xcrun --sdk iphoneos --find clang)
export CXX=$(xcrun --sdk iphoneos --find clang++)
export CFLAGS="-arch arm64 -isysroot $(xcrun --sdk iphoneos --show-sdk-path) -mios-version-min=12.0"
export LDFLAGS="-arch arm64 -isysroot $(xcrun --sdk iphoneos --show-sdk-path)"

# Create build directory
mkdir -p build

# Build iOS framework
echo "Building Mihomo.xcframework..."
gomobile bind \
    -target=ios \
    -iosversion=12.0 \
    -prefix=Mihomo \
    -o build/Mihomo.xcframework \
    -ldflags="-s -w" \
    -tags="ios,cgo" \
    ./experimental/libbox

echo "Build completed: build/Mihomo.xcframework"

# Build iOS binary
echo "Building mihomo-ios..."
go build -tags ios -trimpath \
  -ldflags "-X 'github.com/metacubex/mihomo.Version=$(git describe --tags)' \
  -X 'github.com/metacubex/mihomo.BuildTime=$(date -u)'" \
  -o build/mihomo-ios

echo "Build completed: build/mihomo-ios" 