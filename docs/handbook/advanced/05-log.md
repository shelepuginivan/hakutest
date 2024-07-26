---
title: Log
description: Learn how to view Hakutest logs
---

# Log

---

> [!TIP] You will learn:
>
> -   Structure of Hakutest log
> -   How to view Hakutest log

## Overview

Hakutest logs events such as HTTP requests, errors that occur, etc. in
JSON-like format. You can view the Hakutest log for troubleshooting and
debugging purposes.

## Log file

Log is stored in `hakutest.log` file. Its location depends on the operating
system you are using:

:::details Linux

The log file is in `$XDG_CACHE_HOME/hakutest`.

This path defaults to `$HOME/.cache/hakutest` if `$XDG_CACHE_HOME` is not set
as defined in the [XDG Base Directory
Specification](https://specifications.freedesktop.org/basedir-spec/basedir-spec-latest.html#variables).

:::

:::details Windows

The log file is in one of the following paths:

-   `%AppData%/cache/hakutest`
-   `%LocalAppData%/cache/hakutest`

:::

## Log format

Each line (entry) of the log file is a JSON object. It contains the following information:

| Field     | Description                                     | Note                                                         |
| --------- | ----------------------------------------------- | ------------------------------------------------------------ |
| `level`   | Level of the log entry                          | One of `debug`, `info`, `warn`, `error`, `fatal`, or `panic` |
| `time`    | UNIX timestamp of the log entry                 | &ndash;                                                      |
| `message` | Any additional information logged by Hakutest   | &ndash;                                                      |
| `error`   | Additional information about the occurred error | &ndash;                                                      |
| `method`  | HTTP method of the incoming request             | Only present in HTTP log entries                             |
| `path`    | HTTP path of the incoming request               | Only present in HTTP log entries                             |
| `status`  | HTTP status of Hakutest response                | Only present in HTTP log entries                             |
| `latency` | Latency of HTTP request (ms)                    | Only present in HTTP log entries                             |

> [!IMPORTANT]
> While log entries are valid JSON object, the entire log file itself is not.
> That means that Hakutest log should be parsed line by line.

## Debug

By default Hakutest won't log debug information. You can
[configure](/handbook/advanced/02-configuration#configuration-file-fields)
Hakutest to include these messages in log.
