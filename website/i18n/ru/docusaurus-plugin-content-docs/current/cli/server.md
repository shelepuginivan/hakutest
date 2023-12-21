---
sidebar_position: 3
title: Сервер
description: Команда запуска сервера Hakutest
---

# Команда `server`

Синтаксис: `hakutest [порт]`

Команда `hakutest server` используется для запуска сервера Hakutest.

Если не было передано ни одного аргумента, сервер будет прослушивать порт, указанный в конфигурационном файле (см. [Конфигурация сервера](/docs/configuration/server#port))

Вы можете указать другой порт, передав его как первый аргумент. Значение порта должно быть целым числом от 1024 до 65535.

### Примеры

1.  Запустить сервер на порту, указанном в конфигурационном файле:

    ```shell
    hakutest server
    ```

2.  Запустить на порту `8000`:

    ```shell
    hakutest server 8000
    ```

3.  Невалидный порт (`80` \< `1024`), приводит к ошибке:

    ```shell
    hakutest server 80
    ```

    ```txt title='Output'
    listen tcp :80: bind: permission denied
    exit status 1
    ```

4.  Невалидный порт (нечисловое значение), приводит к ошибке:

    ```shell
    hakutest server some_string
    ```

    ```txt title='Output'
    listen tcp: lookup tcp/some_string: Servname not supported for ai_socktype
    exit status 1
    ```
