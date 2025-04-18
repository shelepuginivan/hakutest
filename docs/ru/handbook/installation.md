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
66162c7884ea816eeb3e75dd2c6a2a56e460046e2cdd94292cefce4403e8a1e2  hakutest.AppImage
68cd545c4883d108b4297ee6423b12ce23a9086e5b00910e5a57d19bfaa1f002  hakutest.deb
a1a435b867c3e7c98615c1df2f7513be25c4eb8981e3d6e6024fe9a152ba223e  hakutest-linux-x86_64.tar.gz
9a7da46b2f2dec13121b38b38024e9a16d59839bddd52ad820177d1a8adf17e9  hakutest-win-x86_64.zip
```
