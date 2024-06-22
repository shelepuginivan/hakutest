#!/usr/bin/env python3
import os
import subprocess

PLATFORM = "Linux (x86_64)"
BOLD_LIGHT_GREEN = "\033[1;92m"
GREEN = "\033[0;32m"
RED = "\033[0;31m"
FAINT = "\033[2m"
RESET = "\033[0m"


def set_environment() -> None:
    os.environ["TARGET_DIR"] = "./target"
    os.environ["BUILD_DIR"] = "./build"
    os.environ["DPKG_DIR"] = "./target/hakutest"
    os.environ["GOARCH"] = "amd64"
    os.environ["GOOS"] = "linux"
    os.environ["NO_STRIP"] = "true"


def execute_commands(commands: list[str]) -> bool:
    with open(os.devnull, mode='w') as devnull:
        for command in commands:
            if subprocess.run(
                command,
                shell=True,
                stdout=devnull,
                stderr=devnull,
            ).returncode != 0:
                return False
        return True


def build_step(
        index: int,
        name: str,
        commands: list[str],
        end = " : ",
    ) -> None:
    print(f"{FAINT}[{index}]{RESET} {name}{end}", end="", flush=True)

    if execute_commands(commands):
        print(f"{GREEN}OK{RESET}")
    else:
        print(f"{RED}FAILED{RESET}")


def main():
    print(f"\nBuilding {BOLD_LIGHT_GREEN}Hakutest{RESET} for {PLATFORM}", end="\n\n")

    set_environment()

    build_step(1, "Compiling Hakutest binaries", [
        "go build -trimpath -o $TARGET_DIR/linux/hakutest ./cmd/hakutest",
        "CGO_ENABLED=1 go build -trimpath -o $TARGET_DIR/linux/hakutest-gtk ./cmd/hakutest-gtk",
    ])

    build_step(2, "Packaging AppImage         ", [
        "mkdir -p $TARGET_DIR/hakutest.AppDir/usr/bin",
        "cp $BUILD_DIR/appimage/AppRun $TARGET_DIR/hakutest.AppDir",
        "cp $BUILD_DIR/resources/hakutest-gtk.desktop $TARGET_DIR/hakutest.AppDir",
        "cp $BUILD_DIR/resources/hakutest.svg $TARGET_DIR/hakutest.AppDir",
        "cp $TARGET_DIR/linux/hakutest-gtk $TARGET_DIR/hakutest.AppDir/usr/bin",
        "$BUILD_DIR/vendor/linuxdeploy-x86_64.AppImage --appdir $TARGET_DIR/hakutest.AppDir --plugin gtk --output appimage",
        "mv ./Hakutest-x86_64.AppImage $TARGET_DIR/hakutest.AppImage",
        "rm -r $TARGET_DIR/hakutest.AppDir",
    ])

    build_step(3, "Packaging .deb             ", [
        "mkdir -p $DPKG_DIR/DEBIAN",
        "cp $BUILD_DIR/deb/control $DPKG_DIR/DEBIAN",
        "cp $BUILD_DIR/deb/preinst $DPKG_DIR/DEBIAN",
        "mkdir -p $DPKG_DIR/usr/bin",
        "cp $TARGET_DIR/linux/hakutest $DPKG_DIR/usr/bin",
        "cp $TARGET_DIR/linux/hakutest-gtk $DPKG_DIR/usr/bin",
        "mkdir -p $DPKG_DIR/usr/share/applications",
        "cp $BUILD_DIR/resources/hakutest-gtk.desktop $DPKG_DIR/usr/share/applications",
        "mkdir -p $DPKG_DIR/usr/share/icons",
        "cp $BUILD_DIR/resources/hakutest.svg $DPKG_DIR/usr/share/icons",
        "dpkg --build $DPKG_DIR",
        "rm -r $DPKG_DIR",
    ])

    build_step(4, "Packaging tarball          ", [
        "tar -czf $TARGET_DIR/hakutest-linux-x86_64.tar.gz --transform 's/^./hakutest/' -C $TARGET_DIR/linux .",
    ])

    build_step(5, "Cleaning up                ", [
        "rm -r $TARGET_DIR/linux",
    ])

    print()


if __name__ == "__main__":
    main()
