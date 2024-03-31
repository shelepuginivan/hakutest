---
sidebar_position: 2
---

# Установка

## Windows

1. Скачайте [hakutest-win64.zip](https://github.com/shelepuginivan/hakutest/releases/download/v0.1.1/hakutest-win64.zip).
2. Распакуйте архив.
3. Запустите `hakutest-server.exe`.

## Linux

### Через скрипт установки

```shell
curl -fsSL https://raw.githubusercontent.com/shelepuginivan/hakutest/main/scripts/install.sh | bash
```

### Tarball

1. Скачайте [hakutest-linux64.tar.gz](https://github.com/shelepuginivan/hakutest/releases/download/v0.1.1/hakutest-linux64.tar.gz).
2. Распакуйте архив.
3. Запустите `./hakutest`.

## macOS

1. Скачайте [hakutest-macos.tar.gz](https://github.com/shelepuginivan/hakutest/releases/download/v0.1.1/hakutest-macos.tar.gz).
2. Распакуйте архив.
3. Запустите `./hakutest`.

## Интернационализация

Чтобы изменить язык интерфейса Hakutest на русский, выполните следующие действия:

1. Скачайте файл интернационализации [i18n.yaml](pathname:///files/i18n/ru/i18n.yaml).
2. Поместите скачанный файл в директорию установки Hakutest.

## Контрольные суммы SHA256

```
ac3c03abb7995c28de4240ac1cb617d64e2ea48dfeb0a62bf1c81445b1c30b38  hakutest-linux64.tar.gz
995559196be3005aeca71f5b358dc37f120d7f73030506e636a42abf2dcfe32f  hakutest-macos.tar.gz
417eb430a839fc567b1ed79d0f4461b99afb9bd0b1cc1048f2688b3863c98136  hakutest-manual.tar.gz
0e01f36a19517d8a929a9e7f86149648cd175cd2d9164dca639af94e37427def  hakutest-win64.zip
```
