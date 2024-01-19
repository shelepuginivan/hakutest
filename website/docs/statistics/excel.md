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

Syntax: `hakutest statistics <test-name> excel`

_Where `test-name` is the name of the test results folder (i.e. its file name) which statistics you want to export_.

This command creates a file named `<test-name>.xlsx` in the current working directory - Excel spreadsheet with exported statistics.

### Example

Assume there is a "My test" folder with test results in the results directory.

```shell title='Command'
hakutest statistics "My test" excel
```

This command will create a file `My test.xlsx` in the current working directory.

:::tip

See [`statistics` command](/docs/cli/statistics) for more information.

:::

## Using `stats_excel.bat` script

Assume there is a "My test" folder with test results in the results directory.

1. Double-click file `stats_excel.bat` in the Hakutest installation directory:

    ```txt {4} title='Directory structure'
    hakutest/
        hakutest.exe
        ...
        stats_excel.bat
    ```

2. Script will prompt you for a test filename:

    > ```
    > Enter the name of the test:
    > ```

3. Enter the filename of the test which statistics you want to export:

    > ```
    > Enter the name of the test: My test
    > ```

4. It will create an Excel spreadsheet in this directory:

    ```txt {5} title='Directory structure'
    hakutest/
        hakutest.exe
        ...
        stats_excel.bat
        My test.xlsx
    ```
