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
7de25c96350f928d4bfe7ad13e7c8b769216a19ea27f9ec67ca10a370c241a55  hakutest.AppImage
87a37374b8ad2a56396b9c643942b1d1575e1466f3c3b895f6ea63a1f546c486  hakutest.deb
abde046865e9af1fb18eb3134afb8048187f8f4d7214c765c16069d57bf57cdc  hakutest-linux-x86_64.tar.gz
da126a087ef89d0f4ff245192cfc0a0287bceb44fa806776a406b86bb5682cc2  hakutest-win-x86_64.zip
```
