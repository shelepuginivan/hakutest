---
sidebar_position: 1
title: Ручное редактирование
description: Узнайте, как редактировать JSON-файлы тестов вручную.
---

# Ручное редактирование JSON-файлов тестов

В Hakutest каждый тест представляет собой JSON-файл. Это означает, что файлы
можно создавать и редактировать вручную или программно.

:::tip

Чтобы узнать об основах структуры тестов в Hakutest, см. [Тесты](/docs/guide/tests).

:::

## Пример файла теста

Ниже представлен пример JSON-файла теста.

```json
{
    "title": "Example test",
    "description": "An example of Hakutest test JSON file",
    "target": "Hakutest users",
    "subject": "Documentation",
    "author": "John Doe",
    "institution": "-",
    "createdAt": "2024-06-08T09:44:14Z",
    "expiresIn": "0001-01-01T00:00:00Z",
    "tasks": [
        {
            "type": "open",
            "text": "Find the root of the equation: 2x - 5 = 0",
            "attachment": null,
            "options": null,
            "answer": "2.5"
        },
        {
            "type": "single",
            "text": "What is the output of the following script?",
            "attachment": {
                "name": "Python script",
                "type": "image",
                "src": "https://example.com/python-script.png"
            },
            "options": ["SyntaxError", "3", "KeyError", "foo"],
            "answer": "3"
        }
    ]
}
```

:::info

Поле `expiresIn` имеет значение `"0001-01-01T00:00:00Z"`, что означает, что у
теста нет времени истечения.

:::

## JSON-схема

Hakutest предоставляет [JSON-схему](https://json-schema.org) для валидации,
автодополнений и проч.

Схема доступна по адресу:
https://hakutest.shelepugin.ru/files/test.schema.json. Вы можете включить её в
JSON-файл следующим образом.

```json
{
    "$schema": "https://hakutest.shelepugin.ru/files/test.schema.json"
}
```

:::tip

См. [JSON schema (английский)](/docs/advanced/tests/json-schema) для более
подробной информации о схеме.

:::
