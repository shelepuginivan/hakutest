---
sidebar_position: 1
title: Импорт
description: Импортировать тест по абсолютному пути
---

# Команда `import`

Синтаксис: `hakutest tests import <путь-до-файла>`

Команда `hakutest tests import` используется для импорта тестовых файлов. Она копирует указанный файл теста в нужную директорию. В сущности, это альтернатива ручному копированию файла.

### Пример

Скопировать файл `my-test.json`, расположенный в `~/Downloads`, в директорию тестов:

```shell
hakutest tests import ~/Downloads/my-test.json
```
