---
sidebar_position: 2
title: Server
description: Hakutest server start command
---

# `hakutest` command

Syntax: `hakutest [port]`

The `hakutest` command is used to start the Hakutest server.

:::info

This command must be run from the installation directory since it uses relative path to load web files:

```txt {2-4} title='Directory structure'
hakutest/
    web/              # This is loaded
        static/
        templates/
    hakutest(.exe)
    ...
```

:::

:::info

You may need to add ./ before the command depending on your system settings as follows:

```shell
./hakutest
```

:::

If no arguments are provided, server will listen port specified in the configuration (see [Server configuration](/docs/configuration/server#port)).

You can override `port` by specifying it as the first argument. The port must be an integer in the range 1024 to 65535.

### Examples

1.  Run on the port specified in the configuration:

    ```shell
    hakutest
    ```

2.  Run on port `8000`:

    ```shell
    hakutest 8000
    ```

3.  Invalid port (`80` \< `1024`), results in error:

    ```shell
    hakutest 80
    ```

    ```txt title='Output'
    listen tcp :80: bind: permission denied
    exit status 1
    ```

4.  Invalid port (not a number), results in error:

    ```shell
    hakutest some_string
    ```

    ```txt title='Output'
    listen tcp: lookup tcp/some_string: Servname not supported for ai_socktype
    exit status 1
    ```
