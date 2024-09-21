---
title: Политика несовместимых изменений
titleTemplate: Руководство Hakutest
description: 'Узнайте о политике несовместимых изменений Hakutest.'
---

# Политика несовместимых изменений

Hakutest следует стандарту [Go Module Version
Numbering](https://go.dev/doc/modules/version-numbers)

Ключевые слова "MUST", "MUST NOT", "REQUIRED", "SHALL", "SHALL NOT", "SHOULD",
"SHOULD NOT", "RECOMMENDED", "MAY" и "OPTIONAL" в этом документе должны
интерпретироваться так, как описано в [RFC
2119](https://datatracker.ietf.org/doc/html/rfc2119).

> [!INFO] Информация
> Здесь и далее под публичным API подразумеваются:
>
> -   [Конфигурация](/ru/handbook/advanced/02-configuration)
> -   [Формат файла теста](/ru/reference/standards/test-schema)
> -   [Формат файла результата](/ru/reference/standards/result-schema)
> -   [Форматы экспорта статистики](/ru/handbook/guide/04-results-and-statistics)
> -   [Библиотека Hakutest (Go)](https://pkg.go.dev/github.com/shelepuginivan/hakutest)

Выпуск _патча_ (например, Hakutest с `1.0.0` до `1.0.1`) **НЕ ДОЛЖЕН** (_MUST
NOT_) затрагивать публичный API и его зависимости. Этот тип выпуска гарантирует
обратную совместимость и стабильность.

_Минорный_ выпуск (например, Hakutest с `1.0.0` до `1.1.0`) **МОЖЕТ** (_MAY_)
повлиять на публичный API или его зависимости, но **ДОЛЖЕН** (_MUST_) должен
быть обратно совместимым. Этот тип выпуска гарантирует обратную совместимость и
стабильность.

_Мажорный_ выпуск (например, Hakutest с `1.0.0` до `2.0.0`) **МОЖЕТ** (_MAY_)
затрагивать публичный API или его зависимости и **МОЖЕТ** (_MAY_) быть
несовместим с предыдущими мажорными версиями.

Коммиты между выпусками Hakutest **МОГУТ** (_MAY_) содержать несовместимые
изменения.
