---
sidebar_position: 1
description: A brief introduction into Hakutest configuration
---

import Tabs from '@theme/Tabs'
import TabItem from '@theme/TabItem'

# About configuration

Hakutest provides various customization options. All settings are organized into several categories:

-   **General**: Common Hakutest settings, such as data storage directories.
-   **Server**: Settings of the Hakutest server.
-   **UI**: User interface customization, including configuration of different pages of the application.
-   **Statistics**: Settings for exporting student results and statistics.

### Config file

The configuration of Hakutest is stored in a `config.yaml` file. The program searches for this file in the following directories:

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

-   **Working directory** - The config file can be placed in the same directory with the executable:

    <Tabs>
        <TabItem value="windows" label="Windows" default>
            ```txt title='Directory structure'
            hakutest/
                ...
                hakutest.exe
                config.yaml
            ```
        </TabItem>
        <TabItem value="unix" label="Linux/macOS">
            ```txt title='Directory structure'
            hakutest/
                ...
                hakutest
                config.yaml
            ```
        </TabItem>
    </Tabs>

### Default configuration

The default Hakutest configuration is as follows:

```yaml title='config.yaml'
general:
    tests_directory: ~/.cache/hakutest/tests # May vary depending on OS
    results_directory: ~/.cache/hakutest/results # May vary depending on OS
    show_results: true
server:
    port: 8080
    mode: release
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
ui:
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
