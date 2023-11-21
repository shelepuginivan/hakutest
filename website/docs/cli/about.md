---
sidebar_position: 1
description: Overwiew of Hakutest command line interface
---

# About CLI

Hakutest provides an extensive command line interface (CLI) for a variety of tasks.

You can see all available commands by running `hakutest help`

```shell
hakutest help
```

:::info

You may need to add ./ before the command depending on your system settings as follows:

```shell
./hakutest help
```

:::

It will print the following help message:

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
