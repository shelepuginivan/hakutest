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
f6bbb0b95274390f6e859bf6759f12e643929ec1d8fe2512fbe2d620b4524d24  hakutest.AppImage
28323efe32a520aa75d54301876b3d7cdbe9f9ac33be09b02b27b270a0aa16eb  hakutest.deb
d6013b9e961beb12afcad3734507371987b69a5ea68611fe5dc87f3f0f25a7d6  hakutest-linux-x86_64.tar.gz
3ec57e96228a8d5f959b0d78fe25fa7c2a524cd46cf90f09eb14d482f390e800  hakutest-win-x86_64.zip
```
