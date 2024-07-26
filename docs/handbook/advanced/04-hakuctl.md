---
title: hakuctl
description: Learn how to use hakuctl &mdash; command line bindings for Hakutest
---

# hakuctl

---

> [!TIP] You will learn:
>
> -   How to use Hakutest via command-line interface

## Overview

`hakuctl` (Hakutest Control) is a command-line interface for Hakutest. It
provides the same functionality as Hakutest that can be invoked
programmatically, allowing for task automation, integrations with another
systems, and extending platform functionality.

## `test` &mdash; manage test(s)

You can manage tests with `hakuctl test` command.

### Help message

```
Manage local test files

Usage:
  hakuctl test [command]

Available Commands:
  delete      Delete test files
  export      Export test files
  import      Import test from a file
  list        List available test files
  search      Incremental search among tests

Flags:
  -h, --help   help for test

Global Flags:
      --no-color   Disable color output

Use "hakuctl test [command] --help" for more information about a command.
```

### Examples

1. `delete` &mdash; delete test(s)

    ```shell
    # Delete single test
    hakuctl test delete "My test"

    # Delete multiple tests
    hakuctl test delete "My test" "Another test" "Third test"

    # Delete tests with automatic confirmation using yes(1)
    yes | hakuctl test delete "This will be deleted"
    ```

    > [!TIP]
    > See [`yes(1)`](https://man.archlinux.org/man/yes.1) manual page for more
    > information.

2. `export` &mdash; export test(s)

    ```shell
    # Export single test to a JSON file.
    hakuctl test export "My Test.json" -o "Documents/My Test.json"

    # Export multiple tests to a ZIP archive
    hakuctl test export test.json another.json third.json -o tests.zip

    # Export test and write it to the standard out
    hakuctl test export test.json -o -
    ```

3. `import` &mdash; import test

    ```shell
    # Import test from a file
    hakuctl test import /path/to/my/test.json
    ```

4. `list` &mdash; list available tests

    ```shell
    # List available tests.
    hakuctl test list

    # List tests that match a regular expression
    hakuctl test list | grep -E -i 'a.*?b'
    ```

    > [!TIP]
    > See [`grep(1)`](https://man.archlinux.org/man/grep.1) manual page for
    > more information.

5. `search` &mdash; incremental search among available tests

    ```shell
    # List tests starting with "My"
    hakuctl test search "My"
    ```

## `result` &mdash; manage test results and statistics

You can manage results and statistics with `hakuctl result` command.

### Help message

```
Manage test results and statistics

Usage:
  hakuctl result [command]

Available Commands:
  delete      Delete results
  export      Generate and export result statistics
  list        List available results
  search      Incremental search among results

Flags:
  -h, --help   help for result

Global Flags:
      --no-color   Disable color output

Use "hakuctl result [command] --help" for more information about a command.
```

### Examples

1. `delete` &mdash; delete results

    ```shell
    # Delete results for a single test
    hakuctl result delete "My test"

    # Delete results for multiple tests
    hakuctl results delete 1 2 3
    ```

2. `export` &mdash; generate and export result statistics

    ```shell
    # Export result statistics as JSON and print it to the standard out
    hakuctl result export "My test"

    # Export result statistics as XLSX into file "Documents/hakutest.xlsx"
    hakuctl results export "Another test" -o Documents/hakutest.xlsx -f xlsx
    ```

3. `list` &mdash; list available results

    ```shell
    # List available results
    hakuctl result list
    ```

4. `search` &mdash; incremental search among available results

    ```shell
    # List available results starting with "Another"
    hakuctl result search "Another"
    ```

## `completion` &mdash; generate shell completions

You can generate shell completion script for `bash`, `zsh`, `fish`, and
`powershell` for easier and more convenient use of `hakuctl`.

### Help message

```
Generate the autocompletion script for hakuctl for the specified shell.
See each sub-command's help for details on how to use the generated script.

Usage:
  hakuctl completion [command]

Available Commands:
  bash        Generate the autocompletion script for bash
  fish        Generate the autocompletion script for fish
  powershell  Generate the autocompletion script for powershell
  zsh         Generate the autocompletion script for zsh

Flags:
  -h, --help   help for completion

Global Flags:
      --no-color   Disable color output

Use "hakuctl completion [command] --help" for more information about a command.
```

### Examples

1. `bash` &mdash; generate the autocompletion script for bash

    ```bash
    # Load completions in the current shell session
    source <(hakuctl completion bash)

    # Load completions for every new session (execute once)
    hakuctl completion bash > /etc/bash_completion.d/hakuctl
    ```

2. `zsh` &mdash; generate the autocompletion script for zsh

    ```zsh
    # Load completions in the current shell session
    source <(hakuctl completion zsh)

    # Load completions for every new session (execute once)
    hakuctl completion zsh > "${fpath[1]}/_hakuctl"
    ```

3. `fish` &mdash; generate the autocompletion script for fish

    ```fish
    # Load completions in the current shell session
    hakuctl completion fish | source

    # Load completions for every new session (execute once)
    hakuctl completion fish > ~/.config/fish/completions/hakuctl.fish
    ```

4. `powershell` &mdash; generate the autocompletion script for PowerShell

    ```powershell
    # Load completions in the current shell session
    hakuctl completion powershell | Out-String | Invoke-Expression
    ```

    > [!TIP]
    > To load completions for every new session, add the output of the above command
    > to your PowerShell profile.
