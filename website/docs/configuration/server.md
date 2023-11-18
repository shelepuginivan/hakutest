---
sidebar_position: 3
description: Hakutest server configuration
---

# Server

Hakutest server configuration, specified under the `server` field in the config file.

## Fields

### `port`

Specifies the port on which the Hakutest server will run.

-   **Value**: a string representing an integer between 1024 and 65535 (e.g., `'8000'`).
-   **Default**: `'8080'`.

### `mode`

Specifies the mode in which the Hakutest server will run.

-   **Value**: `'release'`, `'debug'` or `'test'`. Any other string will fallback to `'release'`.
-   **Default**: `'release'`.

## Example

Example of server configuration:

```yaml title='config.yaml'
server:
    port: '8080'
    mode: release
# Other fields...
```
