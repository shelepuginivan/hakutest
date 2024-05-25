---
sidebar_position: 2
description: Hakutest common settings
---

# General

Hakutest common settings, specified under the `general` field in the config file.

## Fields

### `lang`

Sets the application language.

-   **Value**: supported language.
-   **Default**: `en` (English).

Currently, the following languages are supported.

| Language | Value |
|:--------:|:-----:|
| English  | `en`  |
| Russian  | `ru`  |

### `tests_directory`

Specifies the folder where tests files are stored.

-   **Value**: path (absolute or relative) to the folder where you want to store the tests.
-   **Default**: _Depends on your operating system._

### `results_directory`

Specifies the folder where students' results are stored.

-   **Value**: path (absolute or relative) to the folder where you want to store the results.
-   **Default**: _Depends on your operating system._

### `show_results`

Specifies whether the results will be displayed immediately after the response is sent.

-   **Value**: bool (`true` or `false`), show results or not respectively.
-   **Default**: `true`.

### `overwrite_results`

Specifies whether the results are allowed to be overwritten if the same student resubmits the solution again.

-   **Value**: bool (`true` or `false`), allow overwrite or not respectively.
-   **Default**: `false`.

## Example

Example of general configuration:

```yaml title='config.yaml'
general:
    lang: en
    tests_directory: ./data/tests
    results_directory: ./data/results
    show_results: true
    overwrite_results: false
# Other fields...
```
