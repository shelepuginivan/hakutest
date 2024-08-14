---
title: System tray icon is not displayed
titleTemplate: Hakutest Handbook
description: Learn what to do if the system tray icon is not displayed
---

# System tray icon is not displayed

This problem may occur in some desktop environments on Linux that use older
system tray implementations.

Hakutest system tray implementation uses DBus to communicate through the
`SystemNotifier/AppIndicator` specification, older tray implementations may not
load the icon.

If you are using an older desktop environment or system tray provider, you may
need a proxy application to convert the new DBus calls to the old format.
The recommended tool for GNOME-based trays is [`snixembed`](https://git.sr.ht/~steef/snixembed),
but others are available. Search for "StatusNotifierItems XEmbedded" in the
package manager of your distribution.
