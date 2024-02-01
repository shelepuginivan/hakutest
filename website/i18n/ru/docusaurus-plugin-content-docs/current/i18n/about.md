---
sidebar_position: 1
description: Введение в интернационализацию Hakutest
---

import Tabs from '@theme/Tabs'
import TabItem from '@theme/TabItem'

# Обзор

Hakutest предоставляет возможности для интернационализации (i18n). Все настройки i18n разделены на несколько категорий:

-   **Статистика**: Интернационализация экспорта статистики учеников.
-   **Web**: Интернационализация интерфейса веб-страниц Hakutest.

### Файл интернационализации

Настройки интернационализации Hakutest хранятся в файле `i18n.yaml`. Программа проверяет наличие этого файла в следующих директориях:

-   **Папка конфигурации ОС** - Директория с конфигурациями приложений, определённая операционной системой. Она может варьироваться в зависимости от ОС. Наиболее распространённые пути:

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

-   **Рабочая директория** - Файл интернационализации может находиться в той же папке, что и исполняемый файл Hakutest:

    <Tabs>
        <TabItem value="windows" label="Windows" default>
            ```txt title='Структура папок'
            hakutest/
                ...
                hakutest.exe
                i18n.yaml
            ```
        </TabItem>
        <TabItem value="unix" label="Linux/macOS">
            ```txt title='Структура папок'
            hakutest/
                ...
                hakutest
                i18n.yaml
            ```
        </TabItem>
    </Tabs>

:::note

Файл интернационализации, расположенный в каталоге установки Hakutest, имеет больший приоритет.

:::

### Настройки интернационализации по умолчанию

Ниже представлены настройки интернационализации Hakutest по умолчанию:

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
