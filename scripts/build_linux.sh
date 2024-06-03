#!/bin/sh

BUILD_DIR=./build
TARGET_DIR=./target


build() {
    printf "[1] Compiling Hakutest binaries... "

    GOARCH=amd64
    GOOS=linux 

    go build -trimpath -o "$TARGET_DIR/linux/hakutest" ./cmd/hakutest
    CGO_ENABLED=1 go build -trimpath -o "$TARGET_DIR/linux/hakutest-gtk" ./cmd/hakutest-gtk

    printf "Done\n"
}

package_appimage() {
    printf "[2] Packaging AppImage...          "

    # Ensure build requirements.
    if [ ! -d "$BUILD_DIR/appimage" ]; then
        printf "Cancelled ($BUILD_DIR/appimage not found)!\n"
        return
    elif [ ! -f "$BUILD_DIR/vendor/linuxdeploy-x86_64.AppImage" ]; then
        printf "Cancelled (linuxdeploy not found)!\n"
        return
    elif [ ! -f "$BUILD_DIR/vendor/linuxdeploy-plugin-gtk.sh" ]; then
        printf "Cancelled (linuxdeploy-plugin-gtk not found)!\n"
        return
    fi

    mkdir -p "$TARGET_DIR/hakutest.AppDir/usr/bin"

    # Copy resources into AppDir.
    cp "$BUILD_DIR/appimage/AppRun" "$TARGET_DIR/hakutest.AppDir"
    cp "$BUILD_DIR/resources/hakutest-gtk.desktop" "$TARGET_DIR/hakutest.AppDir"
    cp "$BUILD_DIR/resources/hakutest.svg" "$TARGET_DIR/hakutest.AppDir"
    cp "$TARGET_DIR/linux/hakutest-gtk" "$TARGET_DIR/hakutest.AppDir/usr/bin"

    # Build AppImage using linuxdeploy.
    NO_STRIP=true "$BUILD_DIR/vendor/linuxdeploy-x86_64.AppImage" --appdir "$TARGET_DIR/hakutest.AppDir" --plugin gtk --output appimage > /dev/null 2>&1

    # Cleanup.
    mv ./Hakutest-x86_64.AppImage "$TARGET_DIR/hakutest.AppImage"
    rm -r "$TARGET_DIR/hakutest.AppDir"

    printf "Done\n"
}

package_deb() {
    printf "[3] Packaging deb...               "

    # Ensure dpkg is installed.
    if ! command -v dpkg > /dev/null; then
        printf "Cancelled (dpkg not found)!\n"
        return
    fi

    DPKG_DIR="$TARGET_DIR/hakutest"

    # Copy debian files.
    mkdir -p "$DPKG_DIR/DEBIAN"
    cp "$BUILD_DIR/deb/control" "$DPKG_DIR/DEBIAN"
    cp "$BUILD_DIR/deb/preinst" "$DPKG_DIR/DEBIAN"

    # Copy binaries.
    mkdir -p "$DPKG_DIR/usr/bin"
    cp "$TARGET_DIR/linux/hakutest" "$DPKG_DIR/usr/bin"
    cp "$TARGET_DIR/linux/hakutest-gtk" "$DPKG_DIR/usr/bin"

    # Copy .desktop entries.
    mkdir -p "$DPKG_DIR/usr/share/applications"
    cp "$BUILD_DIR/resources/hakutest-gtk.desktop" "$DPKG_DIR/usr/share/applications"

    # Copy icons.
    mkdir -p "$DPKG_DIR/usr/share/icons"
    cp "$BUILD_DIR/resources/hakutest.svg" "$DPKG_DIR/usr/share/icons"
    
    # Build deb package.
    dpkg --build "$DPKG_DIR" > /dev/null

    # Cleanup.
    rm -r "$DPKG_DIR"

    printf "Done\n"
}

package_tarball() {
    printf "[4] Packaging tarball...           "
    tar -czf "$TARGET_DIR/hakutest-linux-x86_64.tar.gz" --transform 's/^./hakutest/' -C "$TARGET_DIR/linux" .
    printf "Done\n"
}

cleanup() {
    printf "[5] Cleaning up...                 "
    rm -r "$TARGET_DIR/linux"
    printf "Done\n"
}

echo -e "Building \e[1;92mHakutest\e[0m for Linux (x86_64)\n"

build
if [ $? -ne 0 ]; then
    echo -e "\e[1;91mFailed!\e[0m"
    exit 1
fi

package_appimage
package_deb
package_tarball
cleanup
