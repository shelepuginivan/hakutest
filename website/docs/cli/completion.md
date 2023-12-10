---
sidebar_position: 2
title: Completion
description: Generate Hakutest autocompletion scripts
---

# `completion` command

Syntax: `hakutest completion [bash|zsh|fish|powershell]`

The `hakutest completion` command is used for generating completion scripts for your shell.

## Usage

It is recommended to save the autocompletion script to a file and then execute it in the shell profile:

1.  Generate autocompletion scripts for your shell as follows:

    ```shell
    hakutest completion [bash|zsh|fish|powershell] > hakutest_comp
    ```

2.  Move the created file to the completions folder (if any):

    ```shell
    mv hakutest_comp ~/path/to/completion/folder
    ```

3.  Enable completion script in your shell profile. Below are instructions for each supported shell:

    -   **Bash**: Add the following line to your `~/.bashrc` or `~/.bash_profile`:

        ```shell
        source ~/path/to/completion/folder/hakutest_comp
        ```

    -   **Zsh**: Add the following line to your `~/.zshrc`:

        ```zsh
        source ~/path/to/completion/folder/hakutest_comp
        ```

    -   **Fish**: Run the following command:

        ```fish
        source ~/path/to/completion/folder/hakutest_comp
        ```

    -   **PowerShell**: Add the following line to your `$PROFILE`:

        ```powershell
        . ~/path/to/completion/folder/hakutest_comp
        ```

Alternatively, you may generate the completion script dynamically, though it will be slower than the prior method as the scripts will be generated each time you log in to a shell. It is also important to add the installation folder of hakutest to your `$PATH`:

-   **Bash**: Add the following line to your `~/.bashrc` or `~/.bash_profile`:

    ```shell
    source <(hakutest completion bash)
    ```

-   **Zsh**: Add the following line to your `~/.zshrc`:

    ```zsh
    source <(hakutest completion zsh)
    ```

-   **Fish**: Run the following command:

    ```fish
    hakutest completion fish | source
    ```

-   **PowerShell**: Add the following line to your `$PROFILE`:

    ```powershell
    hakutest completion powershell | Out-String | Invoke-Expression
    ```
