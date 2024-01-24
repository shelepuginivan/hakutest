---
sidebar_position: 3
description: Hakutest internationalization for exporting student results and statistics
---

# Statistics

Hakutest internationalization for exporting student results and statistics. Specified under the `stats` field in the internationalization file.

## Fields

### `excel`

Specifies the internationalization for exporting results to an Excel document.

**Options**:

-   `test_results_sheet` - Name of sheet with students' results.
-   `statistics_sheet` - Name of sheet with test statistics.
-   `header_student` - Header for the column with student names.
-   `header_points` - Header for the column with points scored by students.
-   `header_percentage` - Header for the column with percentage of correct answers.

**Visual example**:

![Excel internationalization example](./img/excel-example.webp)

### `image`

Specifies the internationalization for exporting results to a PNG histogram.

**Options**:

-   `title` - Title of the histogram.
-   `label_x` - Label for histogram x-axis (Points).
-   `label_y` - Label for histogram y-axis (Students).

**Visual example**:

![Histogram internationalization example](./img/histogram-example.webp)

## Example

Example of statistics configuration:

```yaml title='i18n.yaml'
stats:
    excel:
        test_results_sheet: Test Results
        statistics_sheet: Test Statistics
        header_student: Student
        header_points: Points
        header_percentage: '%'
    image:
        title: Student Performance
        label_x: Points
        label_y: Students
# Other fields...
```
