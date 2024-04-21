---
sidebar_position: 2
description: Learn more about test files
---

import Tabs from '@theme/Tabs'
import TabItem from '@theme/TabItem'

# Tests

In Hakutest, each test is represented by a JSON file that is stored locally on your computer. When the Hakutest server is running, students can access the test through the browser. The platform dynamically generates a web page for each test, displaying the questions and any additional content specified in the test file.

### Structure of the test file

Test files contain the following information:

-   Test title
-   Description of the test
-   Subject of the test
-   Author of the test
-   Target audience of the test
-   Educational institution
-   Test creation time (automatically generated when the test is created)
-   Test expiry time (optional)
-   Tasks of the test

### Task

The test tasks contain the following information:

-   Task type
    -   Single answer - the student can choose one answer from the given options.
    -   Multiple answers - the student can choose several answers from the given options.
    -   Open question - the student must write the answer in the input field.
    -   File - the student attaches file(s) as an answer.
-   Task text (usually terms of the task)
-   Answer options from which the student can choose an answer
-   Correct answer to the task

:::note

The form of the correct answer differs depending on the type of task.

<Tabs>
    <TabItem value="single" label="Single answer" default>
        Correct answer number.

        For example:

        >   Task text: "2 + 2 = ?"
        >
        >   Answer options:
        >       -   1
        >       -   4
        >       -   5
        >       -   9

        In this case, the correct answer is "2", as this is the correct answer number.
    </TabItem>

    <TabItem value="multiple" label="Multiple answers">
        Numbers of correct answers, separated with commas.

        For example:

        >   Task text: "Choose prime numbers".
        >
        >   Answer options:
        >       -   2
        >       -   3
        >       -   9
        >       -   13
        >       -   15

        In this case, the correct answer is "1,2,4", as these are numbers of correct answers.
    </TabItem>

    <TabItem value="open" label="Open question">
        A string representing correct answer.

        For example:

        >   Task text: "Find root of the equation: x + 7 = 2".
        >
        >   *(There are no answer options since it is an open question)*

        In this case, the correct answer is a string "-5".
    </TabItem>

    <TabItem value="file" label="File">
        Correct answer is not checked. Instead, Hakutest checks whether file was uploaded.
    </TabItem>
</Tabs>

:::

:::info

To prevent accidental or deliberate uploading of too large files, Hakutest allows you to configure the maximum size of files to be uploaded. In case uploaded file is larger than the limit, Hakutest resets the connection.

See [`server.max_upload_size`](/docs/configuration/server#max_upload_size) for more details.

:::

You can also attach an attachment to the task. An attachment can be a video, audio, image, or link. It includes the following fields:

-   Name of attachment
-   Attachment type
    -   Image
    -   Video
    -   Audio
    -   Link
-   Attachment Source (URL)

### Creating and editing test files

Hakutest provides two ways of creating and editing test files:

-   Web editor (see [Test editor](/docs/app/test-editor)) - recommended for most users
-   CLI editor (see [Editor command](/docs/cli/editor))
