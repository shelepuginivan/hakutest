---
sidebar_position: 3
description: Learn about students results
---

# Results

As with tests, Hakutest stores test results as text files on your device.
Results of the same test are grouped into folders, all folders are stored in `results_directory`, which is usually:

-   `~/AppData/Local/hakutest/results` on Windows.
-   `~/.cache/hakutest/results` on Linux and macOS.

:::tip

For more information about results directory, see [General configuration](/docs/configuration/general#results_directory).

:::

### Structure of the test result

Each test result contains the following information:

-   Student name
-   Submission time
-   Student score:
    -   Scored points
    -   Maximum points
    -   Scored percentage
    -   Correctness of answer for each task
-   Information about the test:
    -   Test title
    -   Test author
    -   Test checksum (SHA256)

### Example

Let's look at the following example of test results:

```yaml
student: Alex
submittedAt: 2023-12-10T20:46:06.566462172+03:00
results:
    points: 3
    total: 5
    percentage: 60
    tasks:
        '1':
            answer: '1'
            correct: false
        '2':
            answer: some text
            correct: true
        '3':
            answer: '2'
            correct: true
        '4':
            answer: '1'
            correct: true
        '5':
            answer: some text
            correct: false
test:
    title: Information Security Grade 9
    author: Jane Doe
    sha256: b0f8bf6a584f3002ff9bcf1653a62d8d9b8100468e443bafffab5838354ae17c
```

In this example:

-   The student's name is Alex.
-   The submission time is 2023-12-10T20:46:06.566462172+03:00.
-   The student scored 3 points out of a maximum of 5 points, resulting in a percentage of 60%.
-   The correctness of the answers for each task is as follows:
    -   Task 1: false (answer "1" is incorrect)
    -   Task 2: true (answer "some text" is correct)
    -   Task 3: true (answer "2" is correct)
    -   Task 4: true (answer "1" is correct)
    -   Task 5: false (answer "some text" is incorrect)
-   The test is titled Information Security Grade 9 and was authored by Jane Doe.
-   The test checksum (SHA256) is b0f8bf6a584f3002ff9bcf1653a62d8d9b8100468e443bafffab5838354ae17c.
