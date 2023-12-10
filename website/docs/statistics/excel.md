---
sidebar_position: 2
description: Export test results to an Excel spreadsheet
---

# Excel

Student test results can be exported to an Excel spreadsheet via two methods:

-   By using the `hakutest statistics` command.
-   If you are using the Windows version of Hakutest with scripts, you can run the file `stats_excel.bat`.

## Data format

On the test results page, there is a table displaying the student's scores and the percentage of tasks they completed correctly:

![Excel test results example](./img/excel-results.png)

On the statistics page, there is a table displaying the correctness of tasks performed by students:

![Excel test statistics example](./img/excel-stats.png)

-   Each cell contains an answer submitted by this student.
-   Green cells represent a correct answer.
-   Red cells represent an incorrect answer.
-   If the cell is empty, student didn't submit answer for this task.

:::tip

You can change the static text of the spreadsheet by changing the configuration parameters. See [Statistics configuration](/docs/configuration/stats#excel) for more information.

:::

## Using `hakutest statistics` command

Syntax: `hakutest statistics <name-of-the-test-file> excel`

_Where `name-of-the-test-file` is the filename of test you want to export_.

:::info

Commands must be run from the installation directory.

:::

:::info

You may need to add ./ before the command depending on your system settings as follows:

```shell
./hakutest statistics ...
```

:::

This command creates a file named `<name-of-the-test-file>.xlsx` in the current working directory - Excel spreadsheet with exported statistics.

### Example

Assume there is a test file named "My test.json" in the tests directory.

```shell title='Command'
hakutest statistics "My test" excel
```

This command will create a file `My test.xlsx`:

```txt {4} title='Directory structure'
hakutest/
    hakutest(.exe)
    ...
    My test.xlsx
```

:::tip

See [`statistics` command](/docs/cli/statistics) for more information.

:::

## Using `stats_excel.bat` script

Assume there is a test file named "My test.json" in the tests directory.

1. Double-click file `stats_excel.bat` in the Hakutest installation directory:

    ```txt {4} title='Directory structure'
    hakutest/
        hakutest.exe
        ...
        stats_excel.bat
    ```

2. Script will prompt you for a test filename:

    ![Script prompt](./img/script-stats-prompt-empty.png)

3. Enter the filename of the test which statistics you want to export:

    ![Script prompt with value entered](./img/script-stats-prompt-value.png)

4. It will create an Excel spreadsheet in this directory:

    ```txt {5} title='Directory structure'
    hakutest/
        hakutest.exe
        ...
        stats_excel.bat
        My test.xlsx
    ```
