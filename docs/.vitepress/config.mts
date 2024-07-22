import { defineConfig } from 'vitepress'

export default defineConfig({
    lang: 'en-US',
    title: 'Hakutest',
    description: 'Modern and efficient educational testing',

    locales: {
        root: {
            label: 'English',
            lang: 'en',
            themeConfig: {
                nav: [
                    { text: 'Home', link: '/' },
                    { text: 'Handbook', link: '/handbook/getting-started' },
                ],
                notFound: {
                    quote: 'Hakutaku return to save education',
                },
                sidebar: [
                    {
                        text: 'Getting Started',
                        link: '/handbook/getting-started',
                    },
                    {
                        text: 'Install Hakutest',
                        link: '/handbook/installation',
                    },
                    {
                        text: 'Guide',
                        items: [
                            {
                                text: 'Server',
                                link: '/handbook/guide/01-server',
                            },
                            {
                                text: 'Dashboard',
                                link: '/handbook/guide/02-dashboard',
                            },
                        ],
                    },
                    {
                        text: 'Advanced guide',
                        items: [
                            {
                                text: 'Hakuctl',
                                link: '/handbook/advanced/hakuctl',
                            },
                        ],
                    },
                ],
                footer: {
                    message:
                        'Hakutest is released under the <a href="https://github.com/shelepuginivan/hakutest/blob/main/LICENSE.md" target="_blank">MIT License</a>.',
                    copyright: 'Copyright © 2024-present Ivan Shelepugin',
                },
            },
        },
        ru: {
            label: 'Русский',
            lang: 'ru',
            themeConfig: {
                notFound: {
                    title: 'Страница не найдена',
                    quote: 'Хакутаку возвращаются, чтобы спасти образование',
                    linkText: 'Вернуться на главную страницу',
                },
                nav: [
                    { text: 'Главная', link: '/ru' },
                    {
                        text: 'Справочник',
                        link: '/ru/handbook/getting-started',
                    },
                ],
                sidebar: [
                    {
                        text: 'Введение',
                        link: '/ru/handbook/getting-started',
                    },
                    {
                        text: 'Установка',
                        link: '/ru/handbook/installation',
                    },
                    {
                        text: 'Гайд',
                        items: [
                            {
                                text: 'Сервер',
                                link: '/ru/handbook/guide/01-server',
                            },
                            {
                                text: 'Панель управления',
                                link: '/ru/handbook/guide/02-dashboard',
                            },
                        ],
                    },
                    {
                        text: 'Продвинутый гайд',
                        items: [
                            {
                                text: 'Hakuctl',
                                link: '/handbook/advanced/hakuctl',
                            },
                        ],
                    },
                ],
                footer: {
                    message:
                        'Hakutest лицензирован под <a href="https://github.com/shelepuginivan/hakutest/blob/main/LICENSE.md" target="_blank">Лицензией MIT</a>.',
                    copyright:
                        '© Все права защищены с 2024 по настоящее время, Иван Шелепугин',
                },
            },
        },
    },

    themeConfig: {
        siteTitle: 'Hakutest',
        logo: '/hakutest.svg',

        search: {
            provider: 'local',
        },

        socialLinks: [
            {
                icon: 'github',
                link: 'https://github.com/shelepuginivan/hakutest',
            },
        ],
    },
})
