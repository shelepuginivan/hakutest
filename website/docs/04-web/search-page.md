---
sidebar_position: 2
description: Main page of the app with test search
---

# Search page

The main page of the application is the search page that allows to search for and access existing tests.

![Search page](./img/search.webp)

As you enter text into the search input, a list of suggested existing tests will appear. You can choose one of the suggestions or type the name manually. Once you have entered the name of the test, press "Search" button (or the Enter key). The [page](/docs/app/test-page) with the selected test will open, if it exists.

When the server is running, the search page can be accessed from any device in the local network at the following URL:

`http://<local-ip-of-your-device>:<port>/`

Where:

-   `local-ip-of-your-device` &mdash; The local IP address of the device running the Hakutest server.
-   `port` &mdash; The port on which server is listening (see [Server configuration](/docs/configuration/server#port)).

:::tip

You can create a shortcut to the search page on your students' workstations so
they can easily open Hakutest.

:::

:::tip

To find out your local IP address, see [Local IP address guide](/docs/guide/local-ip).

:::

### Example

Let's assume that:

-   The local IP of your device is `192.168.1.34`.
-   The port on which the server is running is `8080`.

In this example, when the server is running, you can access the search page on http://192.168.1.34:8080/.
