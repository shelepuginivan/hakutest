---
sidebar_position: 1
description: Learn how to start Hakutest server
---

# Server

Hakutest web pages, including the test page and test editor, require a running server.

You can start the Hakutest server using two methods:

-   By running the `hakutest` command in a terminal.
-   If you are using the Windows version of Hakutest with scripts, you can run the file `server.bat`.

## Using `hakutest` command

Syntax: `hakutest [port]`

The `hakutest` command is used to start the Hakutest server.

:::info

Commands must be run from the installation directory.

:::

:::info

You may need to add ./ before the command depending on your system settings as follows:

```shell
./hakutest
```

:::

If no arguments are provided, server will listen port specified in the configuration (see [Server configuration](/docs/configuration/server#port)).

You can override `port` by specifying it as the second argument. The port must be an integer in the range 1024 to 65535.

### Examples

1.  Run on the port specified in the configuration:

    ```shell
    hakutest
    ```

2.  Run on port `8000`:

    ```shell
    hakutest 8000
    ```
