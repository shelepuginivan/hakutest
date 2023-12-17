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
student: Sam
submittedAt: 2023-11-05T15:45:03.885641702+03:00
results:
    points: 4
    total: 5
    percentage: 80
    tasks:
        '1': true
        '2': false
        '3': true
        '4': true
        '5': true
test:
    title: Information Security Grade 9
    author: Jane Doe
    sha256: b0f8bf6a584f3002ff9bcf1653a62d8d9b8100468e443bafffab5838354ae17c
```

In this example:

-   The student's name is Sam.
-   The submission time is 2023-11-05T15:45:03.885641702+03:00.
-   The student scored 4 points out of a maximum of 5 points, resulting in a percentage of 80%.
-   The correctness of the answers for each task is as follows:
    -   Task 1: true (correct)
    -   Task 2: false (incorrect)
    -   Task 3: true (correct)
    -   Task 4: true (correct)
    -   Task 5: true (correct)
-   The test is titled Information Security Grade 9 and was authored by Jane Doe.
-   The test checksum (SHA256) is b0f8bf6a584f3002ff9bcf1653a62d8d9b8100468e443bafffab5838354ae17c.
