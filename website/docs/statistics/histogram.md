---
sidebar_position: 3
description: Export test results to a PNG histogram
---

# Histogram

Student test results can be exported to a histogram via two methods:

-   By using the `hakutest-statistics` executable (recommended for most users).
-   By using the `hakutest statistics` command.

## Data format

Test results statistics is exported as a PNG image with histogram as follows:

![Histogram example](./img/histogram.webp)

-   The horizontal axis indicates the number of points scored by students.
-   The vertical axis indicates the quantity of students who scored this number of points.

:::tip

You can change the static text of the histogram by changing the configuration parameters. See [Statistics internationalization](/docs/i18n/stats#image) for more information.

:::

## Using the `hakutest-statistics` executable

To export results statistics using the `hakutest-statistics` executable, follow the instruction bellow:

1. Run the `hakutest-statistics` (`hakutest-statistics.exe` on Windows).

2. In the opened window, select the test, export format, and the directory where the statistics will be exported:

![Hakutest statistics Histogram](./img/example-image.webp)

3. Press the "Export" button to export statistics.

This will produce a file named `<test-name>.png` (in this example, `Example test.png`) in the selected export directory.

## Using `hakutest statistics` command

Syntax: `hakutest statistics <test-name> excel`

_Where `test-name` is the name of the test results folder (i.e. its file name) which statistics you want to export_.

This command creates a file named `<test-name>.png` in the current working directory - PNG histogram with exported statistics.

### Example

Assume there is a "My test" folder with test results in the results directory.

```shell title='Command'
hakutest statistics "My test" image
```

This command will create a file `My test.png` in the current working directory.

:::tip

See [`statistics` command](/docs/cli/statistics) for more information.

:::
