---
sidebar_position: 1
description: Overview of the Hakutest internationalization
---

import Tabs from '@theme/Tabs'
import TabItem from '@theme/TabItem'

# About

Hakutest provides internationalization (i18n) support. All i18n settings are organized into several categories:

-   **Statistics**: Internationalization of student statistics export.
-   **Web**: Internationalization of the Hakutest web interface.

### I18n file

The i18n of Hakutest is stored in a `i18n.yaml` file. The program searches for this file in the following directories:

-   **OS config directory** - Configuration directory defined by your operating system. It can vary depending on the operating system and its version used. The common paths are:

    <Tabs>
        <TabItem value="windows" label="Windows" default>
            -   `%USERPROFILE%/hakutest`
            -   `%AppData%/hakutest`
            -   `%LocalAppData%/hakutest`
        </TabItem>
        <TabItem value="unix" label="Linux/macOS">
            -   `~/.config/hakutest`
        </TabItem>
    </Tabs>

-   **Working directory** - The i18n file can be placed in the same directory with the executable:

    <Tabs>
        <TabItem value="windows" label="Windows" default>
            ```txt title='Directory structure'
            hakutest/
                ...
                hakutest.exe
                i18n.yaml
            ```
        </TabItem>
        <TabItem value="unix" label="Linux/macOS">
            ```txt title='Directory structure'
            hakutest/
                ...
                hakutest
                i18n.yaml
            ```
        </TabItem>
    </Tabs>

:::note

The internationalization file located in the Hakutest installation directory has higher priority.

:::

### Default i18n settings

Default settings of the Hakutest internationalization are as follows:

```yaml title="i18n.yaml"
server:
    stop_title: Stop Hakutest
    stop_tooltip: Stop Hakutest server and quit
stats:
    app:
        label_test: Test
        label_format: Format
        label_directory: Export to
        submit_text: Export
        cancel_text: Cancel
        select_text: (Select one)
        success_text: Statistics exported successfully!
        error_prefix: 'An error occurred! Detail:'
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
web:
    editor:
        header: Test Editor
        label_title: 'Title:'
        label_description: 'Description:'
        label_subject: 'Subject:'
        label_author: 'Author:'
        label_target: 'Target audience:'
        label_institution: 'Institution:'
        label_expires_in: 'Expires in:'
        label_add_task: + Add task
        label_task_header: Task
        label_task_type: 'Type:'
        label_task_type_single: Single answer
        label_task_type_multiple: Multiple answers
        label_task_type_open: Open question
        label_task_text: 'Text:'
        label_task_answer: 'Answer:'
        label_task_options: Answer options
        label_task_add_option: + Add option
        label_add_attachment: Add attachment
        label_attachment_name: 'Name:'
        label_attachment_type: 'Type:'
        label_attachment_type_file: File
        label_attachment_type_image: Image
        label_attachment_type_video: Video
        label_attachment_type_audio: Audio
        label_attachment_src: 'Source (URL):'
        label_upload_test_input: Upload test file
        label_upload_test_button: Upload and edit
        label_new_test: Create new test
        label_download_test: Download test
    error:
        header: An error occurred!
        details: Details
    expired:
        header: Test expired!
        message: This test is no longer available
    search:
        input_placeholder: Search for a test
        search_button_label: Search
    submitted:
        header: Submitted!
        message: The test results are not displayed according to the system settings
    test:
        student_name_label: 'Your name:'
        open_answer_label: 'Answer:'
        submit_button_label: Submit
```
