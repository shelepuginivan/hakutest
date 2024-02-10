---
sidebar_position: 1
description: Введение в конфигурацию Hakutest
---

import Tabs from '@theme/Tabs'
import TabItem from '@theme/TabItem'

# О конфигурации

Hakutest предоставляет различные возможности кастомизации. Все настройки организованы в несколько категорий:

-   **Основные (`general`)**: Общие настройки Hakutest, такие как директории хранения данных.
-   **Сервер (`server`)**: Настройки сервера Hakutest.

### Конфигурационный файл

Конфигурация Hakutest хранится в файле `config.yaml`. Программа проверяет наличие этого файла в следующих директориях:

-   **Папка конфигурации ОС** - Директория с конфигурациями приложений, определённая операционной системой. Она может варьироваться в зависимости от ОС. Наиболее распространённые пути:

    <Tabs>
        <TabItem value="windows" label="Windows" default>
            -   `%USERPROFILE%/hakutest`
            -   `%AppData%/hakutest`
            -   `%LocalAppData%/hakutest`
        </TabItem>
        <TabItem value="unix" label="Linux/macOS">
            -   `~/.config/hakutest`
        </TabItem>
    </Tabs>

-   **Рабочая директория** - Конфигурационный файл может находиться в той же папке, что и исполняемый файл Hakutest:

    <Tabs>
        <TabItem value="windows" label="Windows" default>
            ```txt title='Структура папок'
            hakutest/
                ...
                hakutest.exe
                config.yaml
            ```
        </TabItem>
        <TabItem value="unix" label="Linux/macOS">
            ```txt title='Структура папок'
            hakutest/
                ...
                hakutest
                config.yaml
            ```
        </TabItem>
    </Tabs>

:::note

Файл конфигурации, расположенный в каталоге установки Hakutest, имеет больший приоритет.

:::

### Конфигурация по умолчанию

Ниже представлена конфигурация Hakutest по умолчанию:

```yaml title='config.yaml'
general:
    # может отличаться в зависимости от ОС
    tests_directory: ~/.cache/hakutest/tests
    results_directory: ~/.cache/hakutest/results
    show_results: true
    overwrite_results: false
server:
    port: 8080
    mode: release
```
