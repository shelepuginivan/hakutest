#!/bin/sh

BUILD_DIR=./build
TARGET_DIR=./target


build() {
    printf "[1] Compiling Hakutest binaries... "
    GOARCH=amd64 GOOS=darwin go build -trimpath -o "$TARGET_DIR/macos/hakutest" ./cmd/hakutest 
    printf "Done\n"
}

package_tarball() {
    printf "[2] Packaging tarball...           "
    tar -czf "$TARGET_DIR/hakutest-macos-arm64.tar.gz" --transform 's/^./hakutest/' -C "$TARGET_DIR/macos" .
    printf "Done\n"
}

cleanup() {
    printf "[3] Cleaning up...                 "
    rm -r "$TARGET_DIR/macos"
    printf "Done\n"
}

echo -e "Building \e[1;92mHakutest\e[0m for macOS (arm64)\n"

build
if [ $? -ne 0 ]; then
    echo -e "\e[1;91mFailed!\e[0m"
    exit 1
fi

package_tarball
cleanup
