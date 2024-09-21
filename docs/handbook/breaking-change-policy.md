---
title: Breaking Change Policy
titleTemplate: Hakutest Handbook
description: 'Learn about Hakutest Breaking Change Policy.'
---

# Breaking Change Policy

Hakutest follows the [Go Module Version
Numbering](https://go.dev/doc/modules/version-numbers) stardard.

The key words "MUST", "MUST NOT", "REQUIRED", "SHALL", "SHALL NOT", "SHOULD",
"SHOULD NOT", "RECOMMENDED", "MAY", and "OPTIONAL" in this document are to be
interpreted as described in [RFC
2119](https://datatracker.ietf.org/doc/html/rfc2119).

> [!INFO]
> By public API hereinafter is meant:
>
> -   [Configuration](/handbook/advanced/02-configuration)
> -   [Test file format](/reference/standards/test-schema)
> -   [Result file format](/reference/standards/result-schema)
> -   [Statistics export formats](/handbook/guide/04-results-and-statistics)
> -   [Go Hakutest Library](https://pkg.go.dev/github.com/shelepuginivan/hakutest)

A _patch release_ (e.g. Hakutest `1.0.0` to `1.0.1`) **MUST NOT** affect public
API and its dependencies. This type of release guarantees backward
compatibility and stability.

A _minor release_ (e.g. Hakutest `1.0.0` to `1.1.0`) **MAY** affect public API
or its dependencies, but **MUST** be backward-compatible. This type of release
guarantees backward compatibility and stability.

A _major release_ (e.g. Hakutest `1.0.0` to `2.0.0`) **MAY** affect public API
or its dependencies and **MAY** be incompatible with preceding major versions.

Commits between Hakutest releases **MAY** provide incompatible changes.
