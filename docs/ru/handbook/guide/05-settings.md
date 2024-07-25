---
title: Настройки
description: Узнайте, как настраивать Hakutest
---

# Настройки

---

> [!TIP] Вы узнаете:
>
> -   Как настраивать Hakutest

![Настройки](./img/settings.png)

Несмотря на то, что Hakutest создан для работы "из коробки", вы можете изменить
некоторые параметры в соответствии с вашими потребностями. Вы можете настроить
Hakutest прямо из браузера в панели управления.

## Основная конфигурация

> [!NOTE] К СВЕДЕНИЮ
> Основные параметры можно изменить в любое время, они не требуют перезапуска.
> Именно эти параметры вы будете менять чаще всего.

В следующей таблице описаны основные параметры, которые можно настроить:

| Параметр                      | Описание                                                                       |
| ----------------------------- | ------------------------------------------------------------------------------ |
| **Язык**                      | Язык интерфейса Hakutest                                                       |
| **Перезаписывать результаты** | Перезаписывать ли результат проверки после повторной отправки ответов учеником |
| **Показывать результаты**     | Показывать ли ученику [его результат](#) сразу после отправки                  |

<!-- TODO: Add link to student interace page -->

## Промежуточная конфигурация

> [!WARNING] ПРЕДУПРЕЖДЕНИЕ
> Настройка промежуточных параметров требует понимания того, как работает
> Hakutest. Мы рекомендуем изменять эти параметры только в случае
> необходимости.

| Параметр                   | Описание                                                                                            | К сведению                                                                                                                    |
| -------------------------- | --------------------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------------------------------------------- |
| **Без иконки в трее**      | Запускать Hakutest без [значка в системном лотке](/ru/handbook/guide/01-server#systray-icon) (трее) | _Потребуется перезапуск_                                                                                                      |
| **Отладка**                | Запускать Hakutest в режиме отладки                                                                 | Позволяет видеть больше информации в [логах](#)                                                                               |
| **Порт**                   | Порт, на котором запускается сервер Hakutest.                                                       | _Потребуется перезапуск_<br><br>Последние 4 цифры адреса Hakutest. Например, в `http://192.168.1.34:8080` порт &mdash; `8080` |
| **Сохранять результаты в** | Директория (папка), в которой Hakutest хранит ответы учеников                                       | Папка должна существовать                                                                                                     |
| **Сохранять тесты в**      | Директория (папка), в которой Hakutest хранит ваши тесты                                            | Папка должна существовать                                                                                                     |

<!-- TODO: Add link to logs page -->

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
> Для дополнительных сведений см. [Безопасность](#)

<!-- TODO: Add link to security page -->

## Применение настроек

После завершения настройки Hakutest нажмите кнопку "Сохранить настройки", чтобы
применить обновленные параметры. Обратите внимание, что для вступления в силу
некоторых параметров требуется перезапуск Hakutest.

<button class="button button__primary">Сохранить настройки</button>