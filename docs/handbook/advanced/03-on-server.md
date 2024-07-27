---
title: Running on a server
titleTemplate: Hakutest Handbook
description: Learn how to run Hakutest in a Wide Area Network (WAN)
---

# Running on a server

---

> [!TIP] You will learn:
>
> -   How to run Hakutest on a server

## Overview

Although Hakutest is primarily designed for use in a Local Area Network (LAN),
it can also run on a server in a Wide Area Network (WAN). It can be useful for

-   Conducting surveys;
-   Organizing municipal or city tests or competitions;
-   Sharing a test set between several educational institutions or organizations.

This page suggests typical use cases for running Hakutest on a server.

> [!IMPORTANT]
> To run Hakutest on a system without Xorg or Wayland, you must set the
> [`general.disable_tray`](/handbook/advanced/02-configuration#configuration-file-fields)
> configuration option to `true`.

> [!NOTE]
> We will use [NGINX](https://nginx.org/en/) in the examples, but Hakutest can
> also work with other web servers. Please check the documentation of your
> chosen web server.
>
> Please also note that these are just examples, you can change the web server
> configuration to suit your needs.

## Single instance

![Single instance](./diagrams/single-instance.svg)

You can run a single Hakutest instance by forwarding requests incoming on port
`80` to port on which Hakutest runs. For example, if Hakutest is
[configured](/handbook/advanced/02-configuration#configuration-file-fields) to
operate on port `8080`:

```nginx
server {
    listen 80;
    location / {
        proxy_pass http://localhost:8080;
        proxy_set_header Host $host;
        proxy_set_header X-Real-IP $remote_addr;
    }
}
```

You can also enable SSL by following the instruction below:

1. Generate SSL certificate and key:

    ```shell
    sudo openssl req -x509 -nodes -days 365 -newkey rsa:2048 -keyout /etc/nginx/ssl/nginx.key -out /etc/nginx/ssl/nginx.crt
    ```

2. Redirect requests incoming on port `80` (http) to port `443` (https):

    ```nginx
    server {
        listen 80;
        server_name example.com;

        return 301 https://$host$request_uri;
    }

    server {
        listen 443 ssl;
        server_name example.com;

        ssl_certificate /etc/nginx/ssl/nginx.crt;
        ssl_certificate_key /etc/nginx/ssl/nginx.key;

        location / {
            proxy_pass http://localhost:8000;
            proxy_set_header Host $host;
            proxy_set_header X-Real-IP $remote_addr;
        }
    }
    ```

## Multiple instances

![Multiple instances](./diagrams/multiple-instances.svg)

You can run multiple Hakutest instances behind a load balancer as follows:

:::code-group

```bash [launch.sh]
#!/usr/bin/env bash

hakutest --general.port 8000 & disown
hakutest --general.port 8001 & disown
hakutest --general.port 8002 & disown

# Other ports...
```

```nginx [nginx.conf]
http {
    upstream hakutest {
        server localhost:8000;
        server localhost:8001;
        server localhost:8002;

        # Other ports...
    }

    server {
        listen 80;
        server_name example.com;

        location / {
            proxy_pass http://hakutest;
            proxy_set_header Host $host;
            proxy_set_header X-Real-IP $remote_addr;
        }
    }
}
```

:::

You can also run multiple instances with different locales. In the following
example, requests without prefix will be redirected to the Hakutest instance
with an English locale, and requests with prefix `/ru` will be redirected to
another instance with a Russian locale:

:::code-group

```bash [launch.sh]
#!/usr/bin/env bash

hakutest --general.port 8000 --general.lang en & disown
hakutest --general.port 8001 --general.lang ru & disown

# Other locales...
```

```nginx [nginx.conf]
server {
    listen 80;

    location /ru/ {
        rewrite ^/ru/(.*)$ /\$1 break;
        proxy_pass http://localhost:8001;
        proxy_set_header Host $host;
        proxy_set_header X-Real-IP $remote_addr;
    }

    # Other locales...

    location / {
        proxy_pass http://localhost:8000;
        proxy_set_header Host $host;
        proxy_set_header X-Real-IP $remote_addr;
    }
}
```

:::

## Systemd service

You can create a `systemd` service to run Hakutest automatically at system startup.

> [!IMPORTANT]
> Note that this will only work on systems with `systemd`. If your distribution
> uses a different initialization system, please refer to the relevant
> documentation.

:::code-group

```systemd [hakutest.service]
[Unit]
Description=Hakutest

[Service]
ExecStart=/usr/bin/hakutest
Restart=always

[Install]
WantedBy=multi-user.target
```

:::

> [!NOTE]
> Adjust the path to Hakutest binary if it is incorrect.

To start the service:

```shell
sudo systemctl enable hakutest.service
sudo systemctl start hakutest.service
```
