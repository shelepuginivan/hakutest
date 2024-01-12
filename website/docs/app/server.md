---
sidebar_position: 1
description: Learn how to start Hakutest server
---

# Server

Hakutest web pages, including the test page and test editor, require a running server.

You can start the Hakutest server using two methods:

-   By running the `hakutest server` command in a terminal.
-   If you are using the Windows version of Hakutest with scripts, you can run the file `server.bat`.

## Using `hakutest server` command

Syntax: `hakutest server`

The `hakutest server` command is used to start the Hakutest server.

:::info

To use commands without specifying full path to the executable, add Hakutest installation directory to `$PATH`:

```shell
export PATH="$PATH:path/to/hakutest/installation"
```

:::

If no arguments are provided, server will listen port specified in the configuration (see [Server configuration](/docs/configuration/server#port)).

You can override the default port with a flag `-p|--port`. The port must be an integer in the range 1024 to 65535.

### Examples

1.  Run on the port specified in the configuration:

    ```shell
    hakutest server
    ```

2.  Run on port `8000`:

    ```shell
    hakutest server -p 8000
    ```

    ```shell
    hakutest server --port 8000
    ```
