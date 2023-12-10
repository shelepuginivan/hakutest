---
sidebar_position: 6
title: Import
description: Import test files command
---

# `import` command

Syntax: `hakutest import <path-to-test-file>`

The `hakutest import` command is used to import test files. It copies specified test file in tests directory. This command is essentially an alternative to manually copying a file to the tests directory.

### Example

```shell
hakutest import ~/Downloads/my-test.json
```

This command will copy the file `my-test.json` located in `~/Downloads` folder to the tests directory.
