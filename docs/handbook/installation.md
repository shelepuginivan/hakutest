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
2. Simply run it.

### Tarball

1. Download [`hakutest-linux-x86_64.tar.gz`](https://github.com/shelepuginivan/hakutest/releases/latest/download/hakutest-linux-x86_64.tar.gz);
2. Extract the archive, e.g. with `tar -xzf hakutest-linux-x86_64.tar.gz`;
3. Run `./hakutest`.

## SHA256

Below are the checksums for each file in the `sha256` format:

```
eff32d91be32f190afa2e89bdf76557e3601f4e19fdbc1e15de27b3f2620be11  hakutest.AppImage
03dcd7398b5daf90aa36f848a3c48e9ba805c31df533758424999e786f0d0c37  hakutest.deb
56fa61a3af9ada164dcf4230b25bfcf40b7ec0a8b6b020dad256a28b2c1007b9  hakutest-linux-x86_64.tar.gz
aad1ca3c66955ff9d7e5cd3342e6203c640eeedc23f8b5895040c426a4bc9f31  hakutest-win-x86_64.zip
```
