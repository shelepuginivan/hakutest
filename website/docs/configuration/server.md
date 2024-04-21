---
sidebar_position: 3
description: Hakutest server configuration
---

# Server

Hakutest server configuration, specified under the `server` field in the config file.

## Fields

### `port`

Specifies the port on which the Hakutest server will run.

-   **Value**: an integer in range 1024 to 65535 (e.g., `8000`).
-   **Default**: `8080`.

### `mode`

Specifies the mode in which the Hakutest server will run.

-   **Value**: `'release'`, `'debug'` or `'test'`. Any other string will fallback to `'release'`.
-   **Default**: `'release'`.

### `max_upload_size`

Sets the limit of the upload file for the file task type.

-   **Value**: integer, max file size in bytes.
-   **Default**: `1048576` (1 MB).

## Example

Example of server configuration:

```yaml title='config.yaml'
server:
    port: 8080
    mode: release
    max_upload_size: 1048576
# Other fields...
```
