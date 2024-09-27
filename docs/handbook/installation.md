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
c5422130ed5e261e6893ec2bd5d26936ad8894975d847e37ba8d4f04d8cc2a8c  hakutest.AppImage
b32e66bb9fcb8015cd29fcae0969f30bbe1bd7e39e8eac19d8454eacd716c9d7  hakutest.deb
edfad15f776ad52bb1700e506bda569dc51bd1a9acf04c0f16eab281c9bb9d9b  hakutest-linux-x86_64.tar.gz
86ffdc054c9fe6cea77606da452d3ce28a2dddfedfe0414c9d97ae08623f7ed7  hakutest-win-x86_64.zip
```
