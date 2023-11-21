---
sidebar_position: 3
description: Export test results to a PNG histogram
---

# Histogram

Student test results can be exported to a histogram via two methods:

-   By using the `hakutest statistics` command.
-   If you are using the Windows version of Hakutest with scripts, you can run the file `stats_image.bat`.

## Data format

Test results statistics is exported as a PNG image with histogram as follows:

![Histogram example](./img/histogram.png)

-   The horizontal axis indicates the number of points scored by students.
-   The vertical axis indicates the quantity of students who scored this number of points.

:::tip

You can change the static text of the histogram by changing the configuration parameters. See [Statistics configuration](/docs/configuration/stats#image) for more information.

:::

## Using `hakutest statistics` command

Syntax: `hakutest statistics <name-of-the-test-file> image`

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

This command creates a file named `<name-of-the-test-file>.png` in the current working directory - PNG histogram with exported statistics.

### Example

Assume there is a test file named "My test.json" in the tests directory.

```shell title='Command'
hakutest statistics "My test" image
```

This command will create a file `My test.png`:

```txt {4} title='Directory structure'
hakutest/
    hakutest(.exe)
    ...
    My test.png
```

## Using `stats_image.bat` script

Assume there is a test file named "My test.json" in the tests directory.

1. Double-click file `stats_image.bat` in the Hakutest installation directory:

    ```txt {4} title='Directory structure'
    hakutest/
        hakutest.exe
        ...
        stats_image.bat
    ```

2. Script will prompt you for a test filename:

    ![Script prompt](./img/script-stats-prompt-empty.png)

3. Enter the filename of the test which statistics you want to export:

    ![Script prompt with value entered](./img/script-stats-prompt-value.png)

4. It will create a PNG image in this directory:

    ```txt {5} title='Directory structure'
    hakutest/
        hakutest.exe
        ...
        stats_image.bat
        My test.png
    ```
