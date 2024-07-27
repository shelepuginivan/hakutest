---
title: hakuctl
titleTemplate: Руководство Hakutest
description: Узнайте, как использовать hakuctl — связки Hakutest для командной строки
---

# hakuctl

---

> [!TIP] Вы узнаете:
>
> -   Как использовать Hakutest через интерфейс командной строки

## Обзор

`hakuctl` (Hakutest Control) &mdash; это интерфейс командной строки для
Hakutest. Он предоставляет тот же функционал, что и Hakutest, но может быть
вызван программно, что позволяет автоматизировать задачи, интегрироваться с
другими системами и расширять возможности платформы.

## `test` &mdash; управление тестами

Вы можете управлять тестами с помощью команды `hakuctl test`.

### Справочное сообщение

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

### Примеры

1. `delete` &mdash; удалить тест(ы)

    ```shell
    # Удалить один тест
    hakuctl test delete "Мой тест"

    # Удалить несколько тестов
    hakuctl test delete "Мой тест" "Другой тест" "Третий тест"

    # Удалить тесты с автоматическим подтверждением, используя yes(1)
    yes | hakuctl test delete "Это будет удалено"
    ```

    > [!TIP] СОВЕТ
    > См. справочную страницу [`yes(1)`](https://man.archlinux.org/man/yes.1)
    > для более подробной информации.

2. `export` &mdash; экспортировать тест(ы)

    ```shell
    # Экспортировать один тест в JSON-файл.
    hakuctl test export "Мой тест.json" -o "Документы/Мой тест.json"

    # Экспортировать несколько тестов в ZIP-архив
    hakuctl test export тест.json другой.json третий.json -o тесты.zip

    # Экспортировать тест и вывести в стандартный вывод
    hakuctl test export тест.json -o -
    ```

3. `import` &mdash; импортировать тесты

    ```shell
    # Импортировать тест из файла
    hakuctl test import /путь/до/моего/теста.json
    ```

4. `list` &mdash; вывести доступные тесты

    ```shell
    # Вывести доступные тесты.
    hakuctl test list

    # Вывести тесты, которые подходят под регулярное выражение
    hakuctl test list | grep -E -i 'a.*?b'
    ```

    > [!TIP] СОВЕТ
    > См. справочную страницу [`grep(1)`](https://man.archlinux.org/man/grep.1)
    > для более подробной информации.

5. `search` &mdash; инкрементный поиск среди доступных тестов

    ```shell
    # Вывести тесты, начинающиеся на "Мой"
    hakuctl test search "Мой"
    ```

## `result` &mdash; управление результатами тестирований и статистикой

Вы можете управлять результатами и статистикой с помощью команды `hakuctl result`.

### Справочное сообщение

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

### Примеры

1. `delete` &mdash; удалить результаты

    ```shell
    # Удалить результаты одного теста
    hakuctl result delete "Мой тест"

    # Удалить результаты нескольких тестов
    hakuctl results delete 1 2 3
    ```

2. `export` &mdash; сгенерировать и экспортировать статистику результатов

    ```shell
    # Экспортировать статистику результатов в формате JSON и вывести ее в стандартный вывод
    hakuctl result export "Мой тест"

    # Экспортировать статистику результатов в формате XLSX в файл "Документы/hakutest.xlsx"
    hakuctl results export "Другой тест" -o Документы/hakutest.xlsx -f xlsx
    ```

3. `list` &mdash; вывести доступные результаты

    ```shell
    # Вывести доступные результаты
    hakuctl result list
    ```

4. `search` &mdash; инкрементный поиск среди доступных результатов

    ```shell
    # Вывести доступные результаты, начинающиеся с "Другой"
    hakuctl result search "Другой"
    ```

## `completion` &mdash; сгенерировать скрипты автодополнений

Вы можете сгенерировать скрипты автодополнений для оболочек `bash`, `zsh`,
`fish` и `powershell` для более простого и удобного использования `hakuctl`.

### Справочное сообщение

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

### Примеры

1. `bash` &mdash; сгенерировать скрипт автодополнения для bash

    ```bash
    # Применить автодополнения для текущей сессии
    source <(hakuctl completion bash)

    # Применять автодополнения для каждой новой сессии
    # (выполните команду единожды)
    hakuctl completion bash > /etc/bash_completion.d/hakuctl
    ```

2. `zsh` &mdash; сгенерировать скрипт автодополнения для zsh

    ```zsh
    # Применить автодополнения для текущей сессии
    source <(hakuctl completion zsh)

    # Применять автодополнения для каждой новой сессии
    # (выполните команду единожды)
    hakuctl completion zsh > "${fpath[1]}/_hakuctl"
    ```

3. `fish` &mdash; сгенерировать скрипт автодополнения для fish

    ```fish
    # Применить автодополнения для текущей сессии
    hakuctl completion fish | source

    # Применять автодополнения для каждой новой сессии
    # (выполните команду единожды)
    hakuctl completion fish > ~/.config/fish/completions/hakuctl.fish
    ```

4. `powershell` &mdash; сгенерировать скрипт автодополнения для PowerShell

    ```powershell
    # Применить автодополнения для текущей сессии
    hakuctl completion powershell | Out-String | Invoke-Expression
    ```

    > [!TIP] СОВЕТ
    > Чтобы применять автодополнения для каждой новой сессии, добавьте вывод
    > команды выше в ваш профиль PowerShell.
