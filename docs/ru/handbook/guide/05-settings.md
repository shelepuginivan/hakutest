---
title: Настройки
titleTemplate: Руководство Hakutest
description: Узнайте, как настраивать Hakutest через встроенный веб-интерфейс
---

# Настройки

---

> [!TIP] Вы узнаете:
>
> -   Как настраивать Hakutest

![Настройки](./img/settings.png)

> [!TIP] СОВЕТ
> Вы можете использовать клавиши <kbd>1</kbd> &ndash; <kbd>4</kbd> для быстрой навигации по интерфейсу учителя:
>
> -   <kbd>1</kbd> &mdash; перейти в [Панель управления](/ru/handbook/guide/02-dashboard)
> -   <kbd>2</kbd> &mdash; перейти в [Меню тестов](/ru/handbook/guide/03-tests#меню-тестов)
> -   <kbd>3</kbd> &mdash; перейти в [Меню результатов](/ru/handbook/guide/04-results-and-statistics#меню-результатов)
> -   <kbd>4</kbd> &mdash; перейти в **Настройки**

Несмотря на то, что Hakutest создан для работы "из коробки", вы можете изменить
некоторые параметры в соответствии с вашими потребностями. Вы можете настроить
Hakutest прямо из браузера в панели управления.

## Основная конфигурация

> [!NOTE] К СВЕДЕНИЮ
> Основные параметры можно изменить в любое время, они не требуют перезапуска.
> Именно эти параметры вы будете менять чаще всего.

В следующей таблице описаны основные параметры, которые можно настроить:

| Параметр                      | Описание                                                                                                                |
| ----------------------------- | ----------------------------------------------------------------------------------------------------------------------- |
| **Язык**                      | Язык интерфейса Hakutest                                                                                                |
| **Перезаписывать результаты** | Перезаписывать ли результат проверки после повторной отправки ответов учеником                                          |
| **Показывать результаты**     | Показывать ли ученику [его результат](/ru/handbook/guide/06-student-perspective#показ-результатов) сразу после отправки |
| **Тип задания по умолчанию**  | Тип задания по умолчанию при добавлении в [редакторе](/handbook/guide/03-tests#test-editor)                             |

## Промежуточная конфигурация

> [!WARNING] ПРЕДУПРЕЖДЕНИЕ
> Настройка промежуточных параметров требует понимания того, как работает
> Hakutest. Мы рекомендуем изменять эти параметры только в случае
> необходимости.

| Параметр                   | Описание                                                                                            | К сведению                                                                                                                    |
| -------------------------- | --------------------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------------------------------------------- |
| **Без иконки в трее**      | Запускать Hakutest без [значка в системном лотке](/ru/handbook/guide/01-server#systray-icon) (трее) | _Потребуется перезапуск_                                                                                                      |
| **Открывать при запуске**  | Открывать Hakutest в браузере при запуске                                                           | _Потребуется перезапуск_                                                                                                      |
| **Отладка**                | Запускать Hakutest в режиме отладки                                                                 | Позволяет видеть больше информации в [журнале](/ru/handbook/advanced/05-log)                                                  |
| **Порт**                   | Порт, на котором запускается сервер Hakutest                                                        | _Потребуется перезапуск_<br><br>Последние 4 цифры адреса Hakutest. Например, в `http://192.168.1.34:8080` порт &mdash; `8080` |
| **Сохранять результаты в** | Директория (папка), в которой Hakutest хранит ответы учеников                                       | Папка должна существовать                                                                                                     |
| **Сохранять тесты в**      | Директория (папка), в которой Hakutest хранит ваши тесты                                            | Папка должна существовать                                                                                                     |

> [!TIP] Путь до папки
> В Windows вы можете скопировать полный путь к каталогу, в котором будут
> храниться тесты или результаты, используя приведенную ниже инструкцию:
>
> 1. Откройте Проводник;
> 2. Перейдите в нужную папку или создайте ее (например, `Документы/Hakutest/Результаты`);
> 3. Удерживая нажатой клавишу Shift, нажмите правой кнопкой мыши;
> 4. Выберите "Копировать как путь".

## Настройки безопасности

> [!CAUTION] ОСТОРОЖНО
> Настройки безопасности предназначены для опытных пользователей, имеющих
> фундаментальное представление о Hakutest. Изменяйте эти настройки с
> осторожностью.
>
> ---
>
> Для дополнительных сведений см. [Безопасность](/ru/handbook/advanced/01-security).

## Применение настроек

После завершения настройки Hakutest нажмите кнопку "Сохранить настройки", чтобы
применить обновленные параметры. Обратите внимание, что для вступления в силу
некоторых параметров требуется перезапуск Hakutest.

<button class="button button__primary">Сохранить настройки</button>
