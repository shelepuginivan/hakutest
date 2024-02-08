---
sidebar_position: 5
title: Editor
description: Hakutest CLI test editor
---

# `editor` command

Syntax: `hakutest editor [test-name]`

Aliases:

-   `hakutest e`

The `hakutest editor` command is a CLI test editor. It can create new tests and edit existing tests.

## Creating new test

To create new test, run command without any argument:

```shell
hakutest editor
```

## Edit existing test

To edit existing test, run command with one argument - the filename of the test to be edited. The test should be placed in tests directory.

```shell
hakutest editor my-test
```

## Usage

### Test fields

Once you have selected one of the options, the editor will start prompting for values for each of the test fields:

-   Test title
-   Description of the test
-   Subject of the test
-   Author of the test
-   Target audience of the test
-   Educational institution
-   Test expiry time

:::tip

For more information about test structure, it is recommended to read [Tests guide](/docs/guide/tests).

:::

:::info

The default values are shown in brackets as follows:

```txt title='Prompt'
Add new task (y/n) [n]:
```

If you edit an existing test, its fields will be shown in brackets as well:

```txt title='Prompt'
Title of the test [Trigonometric equations]:
```

:::

### Adding tasks

Once you have filled in the fields above, you will be able to add the test tasks:

```txt title='Prompt'
Add new task (y/n) [n]:
```

If `y` is chosen, editor will prompt you to enter the task fields:

-   Task type
-   Task text
-   Answer options
-   Correct answer to the task

#### Task attachment

For each task, you can add an attachment:

```txt title='Prompt'
- Add attachment (y/n) [n]:
```

If `y` is chosen, you will be able to enter the following fields of the attachment:

-   Name of attachment
-   Attachment type
-   Attachment source (URL or path to a file)

:::tip

You can specify the path to a local file as the attachment source. In this case, the editor will convert it to a base64 string and add it as an attachment source.

:::

### Saving test files

When you finish editing a test, it will be automatically saved in the tests directory.
