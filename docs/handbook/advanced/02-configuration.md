---
title: Configuration
titleTemplate: Hakutest Handbook
description: Learn how to configure Hakutest on advanced level
---

# Configuration

---

> [!TIP] You will learn:
>
> -   How to configure Hakutest via configuration file
> -   How to configure Hakutest via CLI flags

## Overview

While Hakutest can be configured [via the web interface
(Settings)](/handbook/guide/05-settings), it also provides several methods to
change the settings programmatically. This can be useful for automating tasks,
integrating with other systems, or managing multiple Hakutest instances.

## Configuration file

Hakutest settings are stored in the [YAML](https://yaml.org/) configuration
file `config.yaml`. Its location depends on the operating system you are using:

:::details Linux

The configuration file is in `$XDG_CONFIG_HOME/hakutest`.

This path defaults to `$HOME/.config/hakutest` if
`$XDG_CONFIG_HOME` is not set as defined in the [XDG Base Directory
Specification](https://specifications.freedesktop.org/basedir-spec/basedir-spec-latest.html#variables).

:::

:::details Windows

The configuration file is in one of the following paths:

-   `%USERPROFILE%/hakutest`
-   `%AppData%/hakutest`
-   `%LocalAppData%/hakutest`

:::

### Configuration file fields

| Field                  | Description                                                                | Value                                                                                                                                                                                                                                                                                                     |
| ---------------------- | -------------------------------------------------------------------------- | --------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| `general.debug`        | Run Hakutest in debug mode                                                 | `true` or `false`                                                                                                                                                                                                                                                                                         |
| `general.disable_tray` | Run Hakutest without a system tray icon                                    | `true` or `false`                                                                                                                                                                                                                                                                                         |
| `general.lang`         | Language of the Hakutest interface                                         | Any of supported languages                                                                                                                                                                                                                                                                                |
| `general.port`         | Port on which Hakutest server is started                                   | Integer between `1024` to `65535`                                                                                                                                                                                                                                                                         |
| `result.overwrite`     | Whether to overwrite the results when a student resubmits answers          | `true` or `false`                                                                                                                                                                                                                                                                                         |
| `result.show`          | Whether to show the student the result page after submitting their answers | `true` or `false`                                                                                                                                                                                                                                                                                         |
| `result.path`          | Directory in which Hakutest stores student results                         | Existing directory path                                                                                                                                                                                                                                                                                   |
| `test.path`            | Directory in which Hakutest stores tests                                   | Existing directory path                                                                                                                                                                                                                                                                                   |
| `security.student`     | Security policy applied to student interface                               | `no_verification` &mdash; [For everyone without verification](/handbook/advanced/01-security#no-verification).<br><br>`credentials` &mdash; [By username and password](/handbook/advanced/01-security#credentials).<br><br>`hostonly` &mdash; [Only this device](/handbook/advanced/01-security#hostonly) |
| `security.teacher`     | Security policy applied to teacher interface                               | `no_verification` &mdash; [For everyone without verification](/handbook/advanced/01-security#no-verification).<br><br>`credentials` &mdash; [By username and password](/handbook/advanced/01-security#credentials).<br><br>`hostonly` &mdash; [Only this device](/handbook/advanced/01-security#hostonly) |
| `security.dialect`     | Dialect of the database containing the user data                           | `sqlite`, `mysql`, or `postgres`                                                                                                                                                                                                                                                                          |
| `security.dsn`         | DSN (Data Source Name) of the database containing the user data            | For `sqlite` dialect &mdash; path to the database file.<br><br>For `mysql` and `postgres` dialects &mdash; a connection string.                                                                                                                                                                           |

### Example

Below is an example of a configuration file:

```yaml
general:
    debug: false
    disable_tray: false
    port: 8080
    lang: en
result:
    overwrite: true
    path: /home/user/.local/share/hakutest/results
    show: false
test:
    path: /home/user/.local/share/hakutest/tests
security:
    dsn: /home/user/.cache/hakutest/users.db
    dialect: sqlite
    teacher: hostonly
    student: no_verification
```

> [!TIP]
>
> On Linux, sending Hakutest a `SIGUSR1` signal forces it to reload configuration from file:
>
> ```shell
> pkill -USR1 hakutest
> ```

## CLI flags

You can also change the Hakutest settings via CLI flags. The flags are exactly the same as the fields in the configuration file, see [above](#configuration-file-fields).

Below are several examples of overriding configuration via CLI flags:

```shell
# Override language
hakutest --general.lang en

# Override port
hakutest --general.port 3000

# Override security policy for student
hakutest --security.student credentials

# Multiple overrides
hakutest --general.debug --results.overwrite --general.port 5000
```
