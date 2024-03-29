---
sidebar_position: 9
title: Statistics
description: Hakutest statistics export
---

# `statistics` command

Syntax: `hakutest statistics <test-name> [format]`

Aliases:

-   `hakutest stats`

The `hakutest statistics` command is used to export test results statistics to different formats.

-   The first argument `test-name` is the name of the test results folder (i.e. its file name) which statistics you want to export.
-   The second argument `format` is the format of statistics.

### Print results table

If no second argument is provided, the command will print a table showing the results of the students. This format is also used as a fallback for unsupported formats.

```shell
hakutest statistics "My test"
```

```txt title='Output'
#   Student   Points  %
1   Alex      2       40
2   Amelia    4       80
3   Andrew    4       80
4   Benjamin  5       100
5   George    3       60
6   John      5       100
7   Levy      3       60
8   Lisa      4       80
9   Olivia    3       60
10  Peter     5       100
11  Sam       4       80
12  Victor    3       60
13  William   2       40
```

:::note

The actual formatting of the output statistics may vary depending on the terminal.

:::

### Export statistics to Excel

If second argument `excel` is provided, command generates [Excel spreadsheet](/docs/statistics/excel) with statistics.

```shell
hakutest statistics "My test" excel
```

This command will create a file `My test.xlsx` in the current working directory.

### Export statistics to PNG histogram

If second argument `image` is provided, command generates [PNG histogram](/docs/statistics/histogram).

```shell
hakutest statistics "My test" image
```

This command will create a file `My test.png` in the current working directory.
