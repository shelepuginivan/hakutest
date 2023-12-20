---
sidebar_position: 2
title: Автозаполнение
description: Генерация скрипта автозаполнения Hakutest
---

# Команда `completion`

Синтаксис: `hakutest completion [bash|zsh|fish|powershell]`

Команда `hakutest completion` используется для генерации скриптов автозаполнения для различных оболочек.

## Использование

Рекомендуется сохранить скрипт в файл, а затем выполнять его в профиле оболочки:

1.  Сгенерируйте скрипт для вашей оболочки как показано ниже:

    ```shell
    hakutest completion [bash|zsh|fish|powershell] > _hakutest
    ```

2.  Переместите сгенерированный файл в папку скриптов автозаполнений (если таковая имеется):

    ```shell
    mv _hakutest ~/path/to/completion/folder
    ```

3.  Разрешите автозаполнение в профиле вашей оболочки. Ниже представлены инструкции для каждой поддерживаемой оболочки:

    -   **Bash**: Добавьте строку ниже в файл `~/.bashrc` или `~/.bash_profile`:

        ```shell
        source ~/path/to/completion/folder/_hakutest
        ```

    -   **Zsh**: Добавьте строку ниже в файл `~/.zshrc`:

        ```zsh
        source ~/path/to/completion/folder/_hakutest
        ```

    -   **Fish**: Запустите следующую команду:

        ```fish
        source ~/path/to/completion/folder/_hakutest
        ```

    -   **PowerShell**: Добавьте строку ниже в файл `$PROFILE`:

        ```powershell
        . ~/path/to/completion/folder/_hakutest
        ```

В качестве альтернативы, вы можете генерировать скрипт динамически, однако этот метод менее эффективен по сравнению с предыдущим, так как скрипт будет генерироваться каждый раз, когда вы запускаете оболочку. Кроме того, важно добавить директорию установки Hakutest в `$PATH`:

-   **Bash**: Добавьте строку ниже в файл `~/.bashrc` or `~/.bash_profile`:

    ```shell
    source <(hakutest completion bash)
    ```

-   **Zsh**: Добавьте строку ниже в файл `~/.zshrc`:

    ```zsh
    source <(hakutest completion zsh)
    ```

-   **Fish**: Запустите следующую команду:

    ```fish
    hakutest completion fish | source
    ```

-   **PowerShell**: Добавьте строку ниже в файл `$PROFILE`:

    ```powershell
    hakutest completion powershell | Out-String | Invoke-Expression
    ```
