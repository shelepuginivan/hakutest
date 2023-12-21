---
sidebar_position: 1
description: Overwiew of Hakutest command line interface
---

# About CLI

Hakutest provides an extensive command line interface (CLI) for a variety of tasks.

:::info

To use commands without specifying full path to the executable, add Hakutest installation directory to `$PATH`:

```shell
export PATH="$PATH:path/to/hakutest/installation"
```

:::

You can see all available commands by running `hakutest help`:

```shell
hakutest help
```

It will print the following help message:

```txt title='Help message'
Reliable and efficient educational testing platform

Usage:
  hakutest [command]

Available Commands:
  completion  Generate the autocompletion script for the specified shell
  config      Manage the configuration settings
  editor      Edit test files
  help        Help about any command
  import      Import test file
  server      Start hakutest server
  statistics  Test results statistics

Flags:
  -h, --help   help for hakutest

Use "hakutest [command] --help" for more information about a command.
```
