---
sidebar_position: 1
title: Manual edit
description: Learn how to manually edit test JSON files.
---

# Manual editing of JSON test files

In Hakutest, each test is represented by a JSON file. This means that the files
may be created and edited manually or programmatically.

:::tip

To learn about the basics of test structure in Hakutest, see
[Tests](/docs/guide/tests).

:::

## Test file example

Below is an example of a test JSON file.

```json
{
    "title": "Example test",
    "description": "An example of Hakutest test JSON file",
    "target": "Hakutest users",
    "subject": "Documentation",
    "author": "John Doe",
    "institution": "-",
    "createdAt": "2024-06-08T09:44:14Z",
    "expiresIn": "0001-01-01T00:00:00Z",
    "tasks": [
        {
            "type": "open",
            "text": "Find the root of the equation: 2x - 5 = 0",
            "attachment": null,
            "options": null,
            "answer": "2.5"
        },
        {
            "type": "single",
            "text": "What is the output of the following script?",
            "attachment": {
                "name": "Python script",
                "type": "image",
                "src": "https://example.com/python-script.png"
            },
            "options": ["SyntaxError", "3", "KeyError", "foo"],
            "answer": "3"
        }
    ]
}
```

:::info

Field `expiresIn` is set to `"0001-01-01T00:00:00Z"` meaning that the test has
no expiration time.

:::

## JSON schema

Hakutest provides a [JSON schema](https://json-schema.org) for test validation,
completion, etc.

Schema is available at https://hakutest.shelepugin.ru/files/test.schema.json.
You can include it in the JSON file as follows:

```json
{
    "$schema": "https://hakutest.shelepugin.ru/files/test.schema.json"
}
```

:::tip

See [JSON schema](/docs/advanced/tests/json-schema) for details about the
schema.

:::
