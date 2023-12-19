---
sidebar_position: 3
description: Настройки сервера Hakutest
---

# Сервер

Настройки сервера Hakutest определены в секции `server` конфигурационного файла.

## Поля

### `port`

Определяет порт, на котором сервер Hakutest будет запущен.

-   **Значение**: строка, обозначающая целое число от 1024 до 65535.
-   **По умолчанию**: `'8080'`.

### `mode`

Определяет режим, в котором сервер Hakutest будет запущен.
Specifies the mode in which the Hakutest server will run.

-   **Значение**: `'release'`, `'debug'` или `'test'`. Любое другое значение возвращается к `'release'`.
-   **По умолчанию**: `'release'`.

## Пример

Пример секции `server` конфигурационного файла.

```yaml title='config.yaml'
server:
    port: '8080'
    mode: release
# Остальные поля...
```
