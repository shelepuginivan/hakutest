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
4322eecd64de69c8ba0e606e6269946439d7aacaf4087bab82f23a54827102ab  hakutest.AppImage
e47217779ea83124d66fdd276a905650222309c22922b4d70880426fc4438fbf  hakutest.deb
9fce101ebcfdd356e9e8a29a004a4641ec7ee9107b633dc618dcd623c4a483e4  hakutest-linux-x86_64.tar.gz
9baf9b9db0ebf12d4212acce4dd566b0f2628636bd45b74b714086adf5cd4e2b  hakutest-win-x86_64.zip
```
