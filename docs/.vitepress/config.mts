import { defineConfig } from 'vitepress'

export default defineConfig({
    lang: 'en-US',
    title: 'Hakutest',
    description: 'Modern and efficient educational testing',

    themeConfig: {
        siteTitle: 'Hakutest',
        logo: '/hakutest.svg',

        nav: [
            { text: 'Home', link: '/' },
            { text: 'Handbook', link: '/handbook/getting-started' },
        ],

        search: {
            provider: 'local',
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
                        text: 'Lalala',
                        link: '/handbook/guide/lalala',
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

        socialLinks: [
            {
                icon: 'github',
                link: 'https://github.com/shelepuginivan/hakutest',
            },
        ],

        footer: {
            message:
                'Hakutest is released under the <a href="https://github.com/shelepuginivan/hakutest/blob/main/LICENSE.md" target="_blank">MIT License</a>.',
            copyright: 'Copyright © 2024-present Ivan Shelepugin',
        },
    },
})
