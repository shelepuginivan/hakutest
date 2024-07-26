---
title: Конфигурация
description: Узнайте, как настраивать Hakutest на продвинутом уровне
---

# Конфигурация

---

> [!TIP] Вы узнаете:
>
> -   Как настраивать Hakutest через конфигурационный файл
> -   Как настраивать Hakutest через флаги командной строки

## Обзор

Несмотря на то, что Hakutest можно настроить [через веб-интерфейс
(Настройки)](/ru/handbook/guide/05-settings), он также предоставляет несколько
методов для изменения настроек программным путем. Это может быть полезно для
автоматизации задач, интеграции с другими системами или управления несколькими
экземплярами Hakutest.

## Конфигурационный файл

Настройки Hakutest хранятся в конфигурационном [YAML](https://yaml.org/)-файле
`config.yaml`. Его расположение зависит от используемой вами операционной
системы:

:::details Linux

Конфигурационный файл находится в `$XDG_CONFIG_HOME/hakutest`.

По умолчанию (если переменная окружения `$XDG_CONFIG_HOME` не задана), этот
путь имеет значение `$HOME/.config/hakutest`, как описано в [Спецификации XDG Base
Directory](https://specifications.freedesktop.org/basedir-spec/basedir-spec-latest.html#variables).

:::

:::details Windows

Файл конфигурации находится по одному из следующих путей:

-   `%USERPROFILE%/hakutest`
-   `%AppData%/hakutest`
-   `%LocalAppData%/hakutest`

:::

### Поля конфигурационного файла

| Поле                   | Описание                                                                          | Значение                                                                                                                                                                                                                                                                                                             |
| ---------------------- | --------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| `general.debug`        | Запускать Hakutest в режиме отладки                                               | `true` или `false`                                                                                                                                                                                                                                                                                                   |
| `general.disable_tray` | Запускать Hakutest без значка в системном лотке (трее)                            | `true` или `false`                                                                                                                                                                                                                                                                                                   |
| `general.lang`         | Язык интерфейса Hakutest                                                          | Любой из поддерживаемых языков                                                                                                                                                                                                                                                                                       |
| `general.port`         | Порт, на котором запускается сервер Hakutest                                      | Целое число между `1024` и `65535`                                                                                                                                                                                                                                                                                   |
| `result.overwrite`     | Перезаписывать ли результат проверки после повторной отправки ответов учеником    | `true` или `false`                                                                                                                                                                                                                                                                                                   |
| `result.show`          | Показывать ли ученику его результат сразу после отправки                          | `true` или `false`                                                                                                                                                                                                                                                                                                   |
| `result.path`          | Директория, в которой Hakutest хранит ответы учеников                             | Путь к существующей директории                                                                                                                                                                                                                                                                                       |
| `test.path`            | Директория, в которой Hakutest хранит ваши тесты                                  | Путь к существующей директории                                                                                                                                                                                                                                                                                       |
| `security.student`     | Политика безопасности, применяемая к интерфейсу ученика                           | `no_verification` &mdash; [Для всех без верификации](/ru/handbook/advanced/01-security#no-verification).<br><br>`credentials` &mdash; [По имени пользователя и паролю](/ru/handbook/advanced/01-security#credentials).<br><br>`hostonly` &mdash; [Только это устройство](/ru/handbook/advanced/01-security#hostonly) |
| `security.teacher`     | Политика безопасности, применяемая к интерфейсу учителя                           | `no_verification` &mdash; [Для всех без верификации](/ru/handbook/advanced/01-security#no-verification).<br><br>`credentials` &mdash; [По имени пользователя и паролю](/ru/handbook/advanced/01-security#credentials).<br><br>`hostonly` &mdash; [Только это устройство](/ru/handbook/advanced/01-security#hostonly) |
| `security.dialect`     | Диалект базы данных, в которой хранится информация о пользователях                | `sqlite`, `mysql` или `postgres`                                                                                                                                                                                                                                                                                     |
| `security.dsn`         | DSN (Data Source Name) базы данных, в которой хранится информация о пользователях | Для диалекта `sqlite` &mdash; путь до файла базы данных.<br><br>Для диалектов `mysql` и `postgres` &mdash; строка подключения.                                                                                                                                                                                       |

### Пример

Ниже приведён пример конфигурационного файла:

```yaml
general:
    debug: false
    disable_tray: false
    port: 8080
    lang: ru
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

> [!TIP] СОВЕТ
> На Linux, отправка Hakutest'у сигнала `SIGUSR1` заставляет его перезагрузить
> конфигурацию из файла:
>
> ```shell
> pkill -USR1 hakutest
> ```

## Флаги командной строки

Вы также можете изменить настройки Hakutest с помощью флагов командной строки.
Флаги полностью совпадают с полями конфигурационного файла, см.
[выше](#поля-конфигурационного-фаила).

Ниже приведено несколько примеров переопределения конфигурации с помощью флагов
командной строки:

```shell
# Переопределить язык
hakutest --general.lang ru

# Переопределить порт
hakutest --general.port 3000

# Переопределить политику безопасности для ученика
hakutest --security.student credentials

# Несколько переопределений
hakutest --general.debug --results.overwrite --general.port 5000
```
