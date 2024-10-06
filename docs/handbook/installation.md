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
4322eecd64de69c8ba0e606e6269946439d7aacaf4087bab82f23a54827102ab  hakutest.AppImage
e47217779ea83124d66fdd276a905650222309c22922b4d70880426fc4438fbf  hakutest.deb
9fce101ebcfdd356e9e8a29a004a4641ec7ee9107b633dc618dcd623c4a483e4  hakutest-linux-x86_64.tar.gz
9baf9b9db0ebf12d4212acce4dd566b0f2628636bd45b74b714086adf5cd4e2b  hakutest-win-x86_64.zip
```
