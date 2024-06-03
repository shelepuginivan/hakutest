#!/bin/sh

BUILD_DIR=./build
TARGET_DIR=./target


build() {
    printf "[1] Compiling Hakutest binaries... "

    PKG_CONFIG_PATH=/usr/x86_64-w64-mingw32/sys-root/mingw/lib/pkgconfig GOARCH=amd64 GOOS=windows CGO_ENABLED=1 CC=x86_64-w64-mingw32-gcc go build -ldflags "-H=windowsgui" -trimpath -o "$TARGET_DIR/windows/hakutest.exe" ./cmd/hakutest-gtk

    printf "Done\n"
}

package_zip() {
    printf "[2] Packaging zip archive...       "

    ZIP_DIR="$TARGET_DIR/hakutest"

    mkdir -p "$ZIP_DIR/share/glib-2.0"

    cp "$TARGET_DIR/windows/hakutest.exe" "$ZIP_DIR"
    cp "$BUILD_DIR/vendor/dll/"* "$ZIP_DIR"
    cp -r "$BUILD_DIR/vendor/schemas" "$ZIP_DIR/share/glib-2.0"
    glib-compile-schemas "$ZIP_DIR/share/glib-2.0/schemas"

    cd "$TARGET_DIR" && zip -q hakutest-windows-x86_64.zip -r hakutest && cd ..
    rm -r "$ZIP_DIR"

    printf "Done\n"
}

cleanup() {
    printf "[3] Cleaning up...                 "
    rm -r "$TARGET_DIR/windows"
    printf "Done\n"
}

echo -e "Building \e[1;92mHakutest\e[0m for Windows (x86_64)\n"

build
if [ $? -ne 0 ]; then
    echo -e "\e[1;91mFailed!\e[0m"
    exit 1
fi

package_zip
cleanup
