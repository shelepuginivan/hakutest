---
sidebar_position: 3
title: Remove
description: Remove a test file by its name
---

# `remove` command

Syntax: `hakutest tests remove <test>`

Aliases:

-   `hakutest tests rm`

The `hakutest tests remove` command is used to remove a test file by its name. It assumes that the test file with this name is located in the tests directory.

The extension of the test file (`.json`) can be ommited.

### Example

Remove the test file `my-test.json` located in the tests directory:

```shell
hakutest tests remove my-test
```
