---
title: Установка
titleTemplate: Руководство Hakutest
description: 'Установить Hakutest. Образовательная платформа, предназначенная для проведения тестирования, викторин и экзаменов с автоматической проверкой ответов'
---

# Установка

Hakutest доступен для [Windows](#windows) и [Linux](#linux).
Следуйте инструкциям для вашей операционной системы.

Обратите внимание на значок ниже. Он показывает последнюю версию (выпуск)
Hakutest. Если ваша версия ниже (её можно посмотреть в [Панели
управления](/ru/handbook/guide/02-dashboard)) &mdash; вы можете обновить
программу до последней версии, следуя инструкциям по установке для вашей
операционной системы.

![Последняя Версия Hakutest](https://img.shields.io/github/v/release/shelepuginivan/hakutest?style=for-the-badge&color=1b9e14&label=Версия)

## Windows

1. Скачайте [`hakutest-win-x86_64.zip`](https://github.com/shelepuginivan/hakutest/releases/latest/download/hakutest-win-x86_64.zip);
2. Распакуйте архив;
3. Запустите `hakutest.exe`.

## Linux

### Debian-based дистрибутивы

На Debian-based дистрибутивах (таких как Ubuntu), вы можете установить Hakutest как deb пакет:

1. Скачайте [`hakutest.deb`](https://github.com/shelepuginivan/hakutest/releases/latest/download/hakutest.deb);
2. Установите пакет командой `sudo dpkg -i hakutest.deb`;
3. Запустите `hakutest`.

### AppImage

1. Скачайте [`hakutest.AppImage`](https://github.com/shelepuginivan/hakutest/releases/latest/download/hakutest.AppImage);
2. Просто запустите скачанный файл.

### Tarball

1. Скачайте [`hakutest-linux-x86_64.tar.gz`](https://github.com/shelepuginivan/hakutest/releases/latest/download/hakutest-linux-x86_64.tar.gz);
2. Распакуйте архив, например, командой `tar -xzf hakutest-linux-x86_64.tar.gz`;
3. Запустите `./hakutest`.

## SHA256

Ниже предоставлены контрольные суммы в формате `sha256` для каждого файла:

```
4e4560020ada595b7523fc7a8371d7f56cb522a2ff54e806610f669acb9f9886  hakutest.AppImage
f0c5a4bb61c4b8ddfca40d028bf646d398f38cf11650ff2c67ed78b4ff3d0245  hakutest.deb
d76bde6aa6ed76e4c8978c9b31d8e6ee6c693bf48447f0d649f28e46dbb7db2d  hakutest-linux-x86_64.tar.gz
d1e36df82b801dcb3e4c08767a2d1c40c34563fea12af1a92e4fc3eed7bcccc1  hakutest-win-x86_64.zip
```
