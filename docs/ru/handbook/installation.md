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
2. Предоставьте файлу права на исполнение: `chmod +x ./hakutest.AppImage`
3. Запустите `./hakutest.AppImage`

### Tarball

1. Скачайте [`hakutest-linux-x86_64.tar.gz`](https://github.com/shelepuginivan/hakutest/releases/latest/download/hakutest-linux-x86_64.tar.gz);
2. Распакуйте архив, например, командой `tar -xzf hakutest-linux-x86_64.tar.gz`;
3. Запустите `./bin/hakutest`.

## SHA256

Ниже предоставлены контрольные суммы в формате `sha256` для каждого файла:

```
7de25c96350f928d4bfe7ad13e7c8b769216a19ea27f9ec67ca10a370c241a55  hakutest.AppImage
87a37374b8ad2a56396b9c643942b1d1575e1466f3c3b895f6ea63a1f546c486  hakutest.deb
abde046865e9af1fb18eb3134afb8048187f8f4d7214c765c16069d57bf57cdc  hakutest-linux-x86_64.tar.gz
da126a087ef89d0f4ff245192cfc0a0287bceb44fa806776a406b86bb5682cc2  hakutest-win-x86_64.zip
```
