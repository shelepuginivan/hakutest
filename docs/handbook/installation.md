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
66162c7884ea816eeb3e75dd2c6a2a56e460046e2cdd94292cefce4403e8a1e2  hakutest.AppImage
68cd545c4883d108b4297ee6423b12ce23a9086e5b00910e5a57d19bfaa1f002  hakutest.deb
a1a435b867c3e7c98615c1df2f7513be25c4eb8981e3d6e6024fe9a152ba223e  hakutest-linux-x86_64.tar.gz
9a7da46b2f2dec13121b38b38024e9a16d59839bddd52ad820177d1a8adf17e9  hakutest-win-x86_64.zip
```
