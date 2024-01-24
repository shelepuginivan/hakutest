---
sidebar_position: 6
description: Learn how to manage Hakutest configuration
---

# Managing configuration

Hakutest offers two ways to manage the configuration:

-   Manually editing the configuration file.
-   Using the `config` command.

## Using `config` command

Syntax: `hakutest config [field] [value]`.

The `hakutest config` command is used to manage the configuration for Hakutest via the command line interface. It provides functionality for printing and updating configuration fields. This command has three different uses depending on the number of arguments provided.

:::info

To use commands without specifying full path to the executable, add Hakutest installation directory to `$PATH`:

```shell
export PATH="$PATH:path/to/hakutest/installation"
```

:::

### Print the entire config

If no arguments are provided, the command will print the entire config to the console. The printed config includes tables and other formatting for better readability.

**Example**:

```shell title='Command'
hakutest config
```

```txt title='Output'
server
Key   Value
mode  release
port  8080

general
Key                Value
results_directory  ./data/results
tests_directory    ./data/tests
```

### Print a specific field or category

If a single argument `<field>` is provided, the command will print only the specified field of the config. It can also log an entire category by specifying the category name.

**Example (specific field)**:

```shell title='Command'
hakutest config server.port
```

```txt title='Output'
8080
```

**Example (category)**:

```shell title='Command'
hakutest config server
```

```txt title='Output'
Key   Value
port  8080
mode  release
```

### Update configuration

If two arguments are provided, the command will update the configuration by setting the specified `<field>` to the specified `<value>`.

**Example**:

```shell title='Command'
hakutest config server.port 8000
```

This command will update config and set field `server.port` to `8000`. You can validate that the value was set properly by running:

```shell title='Command'
hakutest config server.port
```

It should output the updated value:

```txt title='Output'
8000
```

:::warning

`hakutest config` can only set values for primitive fields and cannot change values within sub-branches. For instance, the following command will result in error:

```shell title='Command'
hakutest config server some_value
```

```txt title='Output'
2023/11/18 20:55:27 can only set primitive values
exit status 1
```

:::
