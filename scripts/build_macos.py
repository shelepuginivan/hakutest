#!/usr/bin/env python3
import os
import subprocess

PLATFORM = "macOS (arm64)"
BOLD_LIGHT_GREEN = "\033[1;92m"
GREEN = "\033[0;32m"
RED = "\033[0;31m"
FAINT = "\033[2m"
RESET = "\033[0m"


def set_environment() -> None:
    os.environ["TARGET_DIR"] = "./target"
    os.environ["BUILD_DIR"] = "./build"
    os.environ["GOARCH"] = "arm64"
    os.environ["GOOS"] = "darwin"


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
        "go build -trimpath -o $TARGET_DIR/macos/hakutest ./cmd/hakutest",
    ])

    build_step(2, "Packaging tarball          ", [
        "tar -czf $TARGET_DIR/hakutest-macos-arm64.tar.gz --transform 's/^./hakutest/' -C $TARGET_DIR/macos .",
    ])

    build_step(3, "Cleaning up                ", [
        "rm -r $TARGET_DIR/macos",
    ])

    print()


if __name__ == "__main__":
    main()
