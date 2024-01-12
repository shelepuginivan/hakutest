---
sidebar_position: 2
title: Server
description: Hakutest server start command
---

# `server` command

Syntax: `hakutest server`

The `hakutest server` command is used to start the Hakutest server.

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

3.  Invalid port (`80` \< `1024`), results in error:

    ```shell
    hakutest server --port 80
    ```

    ```txt title='Output'
    listen tcp :80: bind: permission denied
    exit status 1
    ```

4.  Invalid port (not a number), results in error:

    ```shell
    hakutest server --port some_string
    ```

    ```txt title='Output'
    listen tcp: lookup tcp/some_string: Servname not supported for ai_socktype
    exit status 1
    ```
