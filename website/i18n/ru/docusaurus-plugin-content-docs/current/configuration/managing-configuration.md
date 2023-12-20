---
sidebar_position: 6
description: Узнайте об управлении конфигурацией Hakutest
---

# Управление конфигурацией

Hakutest предоставляет 2 варианта управления конфигурацией:

-   Ручное редактирование конфигурационного файла.
-   Используя команду `config`.

## Использование команды `config`

Синтаксис: `hakutest config [поле] [значение]`.

Команда `hakutest config` используется для управления конфигурацией Hakutest через интерфейс командной строки. Она предоставляет функционал отображения и изменения настроек. Использование отличается в зависимости от количества аргументов.

:::info

Для использования команд без указания полного пути до исполняемого файла, добавьте директорию установки Hakutest в `$PATH`:

```shell
export PATH="$PATH:path/to/hakutest/installation"
```

:::

### Показать всю конфигурацию

Если не передать команде ни одного аргумента, она печатает всю конфигурацию в стандартный вывод, сгрупированных в таблицы.

**Пример**:

```shell title='Команда'
hakutest config
```

```txt title='Вывод'
ui
editor
Key                          Value
label_task_type              Type:
label_institution            Institution:
label_task_type_open         Open question
label_task_header            Task
label_add_task               + Add task
...

test
Key                  Value
open_answer_label    Answer:
submit_button_label  Submit
student_name_label   Your name:

error
Key                  Value
error_details_label  Details
error_header_label   An error occurred!


server
Key   Value
mode  release
port  8080

stats
image
Key      Value
label_x  Points
label_y  Students
title    Student Performance

excel
Key                 Value
header_student      Student
statistics_sheet    Test Statistics
header_percentage   %
test_results_sheet  Test Results
header_points       Points


general
Key                Value
results_directory  ./data/results
tests_directory    ./data/tests
```

### Вывод конкретного поля или секции

Если передать команде один аргумент, она печатает в стандартный вывод только указанное поле (или секцию полей).

**Пример (одно поле)**:

```shell title='Команда'
hakutest config server.port
```

```txt title='Вывод'
8080
```

**Пример (секция)**:

```shell title='Команда'
hakutest config server
```

```txt title='Вывод'
Key   Value
port  8080
mode  release
```

### Изменение конфигурации

Если передать команде два аргумента, она изменяет конфигурацию, устанавливая в указанное `поле` данное `значение`.

**Пример**:

```shell title='Команда'
hakutest config server.port 8000
```

Эта команда изменит конфигурацию, установит значение `8000` полю `server.port`. Вы можете убедиться, что команда верно установила значение, запустив:

```shell title='Команда'
hakutest config server.port
```

Должно быть выведено новое значение:

```txt title='Output'
8000
```

:::warning

Команда `hakutest config` может устанавливать только значения примитивных типов, но не секций конфигурации. К примеру, следующая команда приведёт к ошибке:

```shell title='Команда'
hakutest config server some_value
```

```txt title='Вывод'
2023/11/18 20:55:27 can only set primitive values
exit status 1
```

:::
