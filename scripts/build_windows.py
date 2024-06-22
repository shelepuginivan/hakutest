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
    os.environ["PKG_CONFIG_PATH"] = "/usr/x86_64-w64-mingw32/sys-root/mingw/lib/pkgconfig"
    os.environ["CC"] = "x86_64-w64-mingw32-gcc"
    os.environ["TARGET_DIR"] = "./target"
    os.environ["BUILD_DIR"] = "./build"
    os.environ["ZIP_DIR"] = "./target/hakutest"
    os.environ["GOARCH"] = "amd64"
    os.environ["GOOS"] = "windows"
    os.environ["CGO_ENABLED"] = "1"


def execute_commands(commands: list[str]) -> bool:
    with open(os.devnull, mode='w') as devnull:
        for command in commands:
            if subprocess.run(
                command,
                shell=True,
                stdout=devnull,
                stderr=devnull,
                universal_newlines=True,
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
        "go build -ldflags -H=windowsgui -trimpath -o $TARGET_DIR/windows/hakutest.exe ./cmd/hakutest-gtk",
    ])

    build_step(2, "Packaging zip archive      ", [
        "mkdir -p $ZIP_DIR/share/glib-2.0",
        "cp $TARGET_DIR/windows/hakutest.exe $ZIP_DIR",
        "cp $BUILD_DIR/vendor/dll/* $ZIP_DIR",
        "cp -r $BUILD_DIR/vendor/schemas $ZIP_DIR/share/glib-2.0",
        "glib-compile-schemas $ZIP_DIR/share/glib-2.0/schemas",
        """cd $TARGET_DIR
zip -q hakutest-windows-x86_64.zip -r hakutest
cd ..
"""
        "rm -r $ZIP_DIR",
    ])

    build_step(3, "Cleaning up                ", [
        "rm -r $TARGET_DIR/windows",
    ])

    print()


if __name__ == "__main__":
    main()
