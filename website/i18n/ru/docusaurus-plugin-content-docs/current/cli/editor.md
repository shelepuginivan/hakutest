---
sidebar_position: 5
title: Редактор
description: CLI-редактор тестов Hakutest
---

# Команда `editor`

Синтаксис: `hakutest editor [имя теста]`

Псевдонимы:

-   `hakutest e`

Команда `hakutest editor` - CLI-редактор тестов. Она позволяет создавать новые тесты и редактировать существующие.

## Создание новых тестов

Чтобы создать новый тест, запустите команды без агрументов:

```shell
hakutest editor
```

## Редактирование существующих тестов

Чтобы изменить существующий тест, запустите команду с одним аргументом - именем файла теста, который вы хотите отредактировать. Файл должен находиться в папке тестов:

```shell
hakutest editor my-test
```

## Использование

### Поля теста

После выбора одного из вариантов редактор начнет запрашивать значения для каждого из полей теста:

-   Название теста
-   Описание теста
-   Тема теста
-   Автор теста
-   Целевая аудитория теста
-   Учебное заведение
-   Срок действия теста

:::tip

См. [Гайд - Тесты](/docs/guide/tests) для дополнительной информации.

:::

:::info

Значения по умолчанию указаны в квадратных скобках:

```txt title='Ввод'
Add new task (y/n) [n]:
```

Если вы редактируете существующий тест, его поля также будут показаны в скобках:

```txt title='Ввод'
Title of the test [Trigonometric equations]:
```

:::

### Добавление заданий

После заполнения полей выше, вы можете добавить задания теста:

```txt title='Ввод'
Add new task (y/n) [n]:
```

Если выбран вариант `y` (да), редактор попросит ввести поля задания:

-   Тип задания
-   Текст задания
-   Варианты ответа
-   Правильный ответ

#### Вложения

К каждому заданию вы можете добавить вложение:

```txt title='Ввод'
- Add attachment (y/n) [n]:
```

Если выбран вариант `y`, вы редактор попросит ввести поля вложения:

-   Название вложения
-   Тип вложения
-   Источник вложения (URL или путь до файла)

:::tip

В качестве источника вложений можно указать путь к локальному файлу. В этом случае редактор преобразует его в строку base64 и добавит в качестве источника вложений.

:::

### Сохранение файлов тестов

По завершении редактирования теста, он будет автоматически сохранён в нужную директорию.
