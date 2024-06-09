---
sidebar_position: 1
title: GUI
description: Learn how to customize appearance of the Hakutest graphical app.
---

# Customizing appearance of the graphical app

Hakutest graphical application is built with GTK3, hence you can customize its
appearance by changing the default GTK theme.

## GTK3 configuration file

GTK3 reads a configuration file `$XDG_CONFIG_HOME/gtk-3.0/settings.ini`
(Usually `~/.config/gtk-3.0/settings.ini` on most systems).

In this file you can change parameters such as theme, icons, and cursors. Below
is an example file that applies `Adwaita-dark` theme and `Papirus-Dark` icon
theme:

```ini
[Settings]
gtk-theme-name=Adwaita-dark
gtk-icon-theme-name=Papirus-Dark
```

## Resources for finding GTK themes

If you're looking for more theme options, below are some resources where you
can find themes suitable for your needs:

-   [GNOME-look.org](https://www.gnome-look.org/browse/)
-   [pling.com](https://www.pling.com/)

Additionally, on Linux, repositories of most distributions provide packages
with GTK themes. You cay search for packages with suffix `-gtk-theme` or `-gtk`
using the package manager of your distribution.
