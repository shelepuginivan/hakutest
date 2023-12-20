---
sidebar_position: 1
description: Обзор интерфейса командной строки Hakutest
---

# О CLI

Hakutest предоставляет широкий интерфейс командной строки (CLI) для различных задач.

:::info

Для использования команд без указания полного пути до исполняемого файла, добавьте директорию установки Hakutest в `$PATH`:

```shell
export PATH="$PATH:path/to/hakutest/installation"
```

:::

Вы можете увидеть список всех команд, запустив `hakutest help`:

```shell
hakutest help
```

Программа выведет следующее сообщение:

```txt title='Help message'
Start hakutest server

Usage:
  hakutest [port] [flags]
  hakutest [command]

Available Commands:
  completion  Generate the autocompletion script for the specified shell
  config      Manage the configuration settings
  editor      Edit test files
  help        Help about any command
  import      Import test file
  statistics  Test results statistics

Flags:
  -h, --help   help for hakutest

Use "hakutest [command] --help" for more information about a command.
```
