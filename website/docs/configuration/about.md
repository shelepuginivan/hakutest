---
sidebar_position: 1
description: A brief introduction into Hakutest configuration
---

import Tabs from '@theme/Tabs'
import TabItem from '@theme/TabItem'

# About configuration

Hakutest provides various customization options. All settings are organized into several categories:

-   **General**: Common Hakutest settings, such as data storage directories.
-   **Server**: Settings of the Hakutest server.

### Config file

The configuration of Hakutest is stored in a `config.yaml` file. The program searches for this file in the following directories:

-   **OS config directory** - Configuration directory defined by your operating system. It can vary depending on the operating system and its version used. The common paths are:

    <Tabs>
        <TabItem value="windows" label="Windows" default>
            -   `%USERPROFILE%/hakutest`
            -   `%AppData%/hakutest`
            -   `%LocalAppData%/hakutest`
        </TabItem>
        <TabItem value="unix" label="Linux/macOS">
            -   `~/.config/hakutest`
        </TabItem>
    </Tabs>

-   **Working directory** - The config file can be placed in the same directory with the executable:

    <Tabs>
        <TabItem value="windows" label="Windows" default>
            ```txt title='Directory structure'
            hakutest/
                ...
                hakutest.exe
                config.yaml
            ```
        </TabItem>
        <TabItem value="unix" label="Linux/macOS">
            ```txt title='Directory structure'
            hakutest/
                ...
                hakutest
                config.yaml
            ```
        </TabItem>
    </Tabs>

:::note

The configuration file located in the Hakutest installation directory has higher priority.

:::

### Default configuration

The default Hakutest configuration is as follows:

```yaml title='config.yaml'
general:
    tests_directory: ~/.cache/hakutest/tests # May vary depending on OS
    results_directory: ~/.cache/hakutest/results # May vary depending on OS
    show_results: true
    overwrite_results: false
server:
    port: 8080
    mode: release
```
