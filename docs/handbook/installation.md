---
title: Installation
titleTemplate: Hakutest Handbook
description: 'Install Hakutest. Educational platform designed for testing, quizzes, and exams with automatic answer checking'
---

# Installation

Hakutest is available for [Windows](#windows) and [Linux](#linux).
Follow the instructions for your operating system.

Notice the badge below. It shows the latest version of Hakutest. If your
version (you can check this in [Dashboard](/handbook/guide/02-dashboard)) is
lower, you can update it to the latest version by following the installation
instructions for your operating system.

![Latest Hakutest Release](https://img.shields.io/github/v/release/shelepuginivan/hakutest?style=for-the-badge&color=1b9e14)

## Windows

1. Download [`hakutest-win-x86_64.zip`](https://github.com/shelepuginivan/hakutest/releases/latest/download/hakutest-win-x86_64.zip);
2. Extract the archive;
3. Run `hakutest.exe`.

## Linux

### Debian-based distributions

On Debian-based distributions (like Ubuntu), you can install Hakutest as a deb package:

1. Download [`hakutest.deb`](https://github.com/shelepuginivan/hakutest/releases/latest/download/hakutest.deb);
2. Install the package by running `sudo dpkg -i hakutest.deb`;
3. Run `hakutest`.

### AppImage

1. Download [`hakutest.AppImage`](https://github.com/shelepuginivan/hakutest/releases/latest/download/hakutest.AppImage);
2. Give the file execution permissions: `chmod +x ./hakutest.AppImage`
3. Run `./hakutest.AppImage`

### Tarball

1. Download [`hakutest-linux-x86_64.tar.gz`](https://github.com/shelepuginivan/hakutest/releases/latest/download/hakutest-linux-x86_64.tar.gz);
2. Extract the archive, e.g. with `tar -xzf hakutest-linux-x86_64.tar.gz`;
3. Run `./bin/hakutest`.

## SHA256

Below are the checksums for each file in the `sha256` format:

```
fae44b53c2d8056aff5dcdbb5a8976eab9184dd930a0af96f77e3ca55118283a  hakutest.AppImage
06b19324730b7ff437fd540705f8fe387eb2f3f5a16325eba70af359e9d17227  hakutest.deb
f90cbabb60b74af418fec918233e6c8d6b7a7a71075a2aca21d42e8c32d89599  hakutest-linux-x86_64.tar.gz
bf3d274dadb812bb3c2e50f05a81fde0a8e861c863c663bde68fbbab785a4977  hakutest-win-x86_64.zip
```
