---
title: Запуск на сервере
titleTemplate: Руководство Hakutest
description: Узнайте, как запускать Hakutest в широкополосной сети на сервере
---

# Запуск на сервере

---

> [!TIP] Вы узнаете:
>
> -   Как запускать Hakutest на сервере

## Обзор

Хотя Hakutest в первую очередь предназначен для использования в локальной сети
(LAN), он также может работать на сервере в глобальной (широкополосной) сети
(WAN). Это может быть полезно для:

-   Проведения опросов;
-   Организации муниципальных или городских тестов или соревнований;
-   Совместного использования набора тестов несколькими учебными заведениями
    или организациями.

На этой странице описаны типичные случаи использования Hakutest на сервере.

> [!IMPORTANT] ВАЖНО
> Чтобы запустить Hakutest на системе без Xorg или Wayland, вам необходимо
> установить значение `true` для поля конфигурации
> [`general.disable_tray`](/ru/handbook/advanced/02-configuration#поля-конфигурационного-фаила).

> [!NOTE] К СВЕДЕНИЮ
> Мы будем использовать [NGINX](https://nginx.org/ru/) в примерах, однако
> Hakutest может работать и с другими веб-серверами. Пожалуйста, обратитесь к
> документации выбранного вами веб-сервера.
>
> Также обратите внимание, что здесь приведены лишь примеры, вы можете изменить
> настройки веб-сервера в соответствии с вашими потребностями.

## Один экземпляр

![Один экземпляр](./diagrams/single-instance.svg)

Вы можете запустить один экземпляр Hakutest, направив запросы, поступающие на
порт `80`, на порт, на котором работает Hakutest. Например, если Hakutest
[настроен](/ru/handbook/advanced/02-configuration#поля-конфигурационного-фаила)
для работы на порту `8080`:

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

Вы также можете включить поддержку SSL, следуя инструкции ниже:

1. Сгенерируйте SSL-сертификат и ключ:

    ```shell
    sudo openssl req -x509 -nodes -days 365 -newkey rsa:2048 -keyout /etc/nginx/ssl/nginx.key -out /etc/nginx/ssl/nginx.crt
    ```

2. Перенаправляйте запросы, поступающие на порт `80` (http), на порт `443` (https):

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

## Несколько экземпляров

![Несколько экземпляров](./diagrams/multiple-instances.svg)

Вы можете запустить несколько экземпляров Hakutest за балансировщиком нагрузки
следующим образом:

:::code-group

```bash [launch.sh]
#!/usr/bin/env bash

hakutest --general.port 8000 & disown
hakutest --general.port 8001 & disown
hakutest --general.port 8002 & disown

# Другие порты...
```

```nginx [nginx.conf]
http {
    upstream hakutest {
        server localhost:8000;
        server localhost:8001;
        server localhost:8002;

        # Другие порты...
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

Вы также можете запустить несколько экземпляров с разными локалями. В следующем
примере запросы без префикса будут перенаправлены на экземпляр Hakutest с
русской локалью, а запросы с префиксом `/en` будут перенаправлены на другой
экземпляр с английской локалью:

:::code-group

```bash [launch.sh]
#!/usr/bin/env bash

hakutest --general.port 8000 --general.lang ru & disown
hakutest --general.port 8001 --general.lang en & disown

# Другие локали...
```

```nginx [nginx.conf]
server {
    listen 80;

    location /en/ {
        rewrite ^/en/(.*)$ /\$1 break;
        proxy_pass http://localhost:8001;
        proxy_set_header Host $host;
        proxy_set_header X-Real-IP $remote_addr;
    }

    # Другие локали...

    location / {
        proxy_pass http://localhost:8000;
        proxy_set_header Host $host;
        proxy_set_header X-Real-IP $remote_addr;
    }
}
```

:::

## Сервис systemd

Вы можете создать службу (сервис) `systemd` для автоматического запуска
Hakutest при старте системы.

> [!IMPORTANT] ВАЖНО
> Обратите внимание, что это будет работать только на системах с `systemd`.
> Если в вашем дистрибутиве используется другая система инициализации,
> обратитесь к соответствующей документации.

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

> [!NOTE] К СВЕДЕНИЮ
> Измените путь до исполняемого файла Hakutest в случае необходимости.

Чтобы запустить сервис:

```shell
sudo systemctl enable hakutest.service
sudo systemctl start hakutest.service
```
