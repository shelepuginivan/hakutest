---
sidebar_position: 4
description: Test editor that allows to create and edit tests
---

# Test editor

Test editor allows to create and edit tests.

When the server is running, it can be accessed at the following URL:

`http://localhost:<port>/editor/upload`

Where `port` is the port on which server is listening.

Default test editor URL is http://localhost:8080/editor/upload.

![Test editor upload](./img/test-editor-upload.png)

## Using the editor

:::tip

For more information about test structure, it is recommended to read [Tests guide](/docs/guide/tests).

:::

### Create new test

To create new test, click "Create new test" button.

### Edit existing test

To edit existing test file, follow the instruction:

1.  Upload test file that you want to edit by clicking on "Upload test file".
2.  Click "Edit test" button.

### Editor

Once you have selected one of the options, an editor page will open where you can edit the test by filling in the required fields:

![Test editor empty](./img/test-editor-edit-empty.png)
![Test editor filled](./img/test-editor-edit-filled.png)

:::tip

To create a test with no expiry time, leave the field "Expires in" unchanged or empty.

:::

### Adding tasks

To add a task, click the "+ Add task" button:

![Test editor add task](./img/test-editor-add-task-empty.png)
![Test editor add filled](./img/test-editor-add-task-filled.png)

Here is the example for other types of tasks:

![Test editor many tasks](./img/test-editor-many-tasks.png)

### Adding attachment

You can also add an attachment to the task by ticking the checkbox "Add attachment":

![Test editor add attachment](./img/test-editor-add-attachment.png)

### Download the test

Once you have finished editing the test, download it by clicking "Download test" button.

To have access to the created test, don't forget to move it to the tests folder.
