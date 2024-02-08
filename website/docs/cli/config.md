---
sidebar_position: 4
title: Config
description: Hakutest configuration management command
---

# `config` command

Syntax: `hakutest config [field] [value]`.

Aliases:

-   `hakutest cfg`

The `hakutest config` command is used to manage the configuration for Hakutest via the command line interface. It provides functionality for printing and updating configuration fields. This command has three different uses depending on the number of arguments provided.

### Print the entire config

If no arguments are provided, the command will print the entire config to the console. The printed config includes tables and other formatting for better readability.

**Example**:

```shell title='Command'
hakutest config
```

```txt title='Output'
ui
editor
Key                          Value
label_task_type              Type:
label_institution            Institution:
label_task_type_open         Open question
label_task_header            Task
label_add_task               + Add task
...

test
Key                  Value
open_answer_label    Answer:
submit_button_label  Submit
student_name_label   Your name:

error
Key                  Value
error_details_label  Details
error_header_label   An error occurred!


server
Key   Value
mode  release
port  8080

stats
image
Key      Value
label_x  Points
label_y  Students
title    Student Performance

excel
Key                 Value
header_student      Student
statistics_sheet    Test Statistics
header_percentage   %
test_results_sheet  Test Results
header_points       Points


general
Key                Value
results_directory  ./data/results
tests_directory    ./data/tests
```

:::note

The actual formatting of the output configuration may vary depending on the terminal.

:::

### Print a specific field or category

If a single argument `<field>` is provided, the command will print only the specified field of the config. It can also log an entire category by specifying the category name.

**Example (specific field)**:

```shell title='Command'
hakutest config server.port
```

```txt title='Output'
8080
```

**Example (category)**:

```shell title='Command'
hakutest config server
```

```txt title='Output'
Key   Value
port  8080
mode  release
```

### Update configuration

If two arguments are provided, the command will update the configuration by setting the specified `<field>` to the specified `<value>`.

**Example**:

```shell title='Command'
hakutest config server.port 8000
```

This command will update config and set field `server.port` to `8000`. You can validate that the value was set properly by running:

```shell title='Command'
hakutest config server.port
```

It should output the updated value:

```txt title='Output'
8000
```

:::warning

`hakutest config` can only set values for primitive fields and cannot change values within sub-branches. For instance, the following command will result in error:

```shell title='Command'
hakutest config server some_value
```

```txt title='Output'
2023/11/18 20:55:27 can only set primitive values
exit status 1
```

:::
