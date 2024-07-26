import { DefaultTheme } from 'vitepress'

export const enThemeConfig = {
    notFound: {
        quote: 'Hakutaku return to save education',
    },

    footer: {
        message:
            'Hakutest is released under the <a href="https://github.com/shelepuginivan/hakutest/blob/main/LICENSE.md" target="_blank">MIT License</a>.',
        copyright: 'Copyright © 2024-present Ivan Shelepugin',
    },

    nav: nav(),
    sidebar: sidebar(),
} satisfies DefaultTheme.Config

function nav(): DefaultTheme.NavItem[] {
    return [
        { text: 'Home', link: '/' },
        { text: 'Handbook', link: '/handbook/getting-started' },
    ]
}

function sidebar(): DefaultTheme.SidebarItem[] {
    return [
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
                {
                    text: 'Tests',
                    link: '/handbook/guide/03-tests',
                },
                {
                    text: 'Results and Statistics',
                    link: '/handbook/guide/04-results-and-statistics',
                },
                {
                    text: 'Settings',
                    link: '/handbook/guide/05-settings',
                },
                {
                    text: 'Student Perspective',
                    link: '/handbook/guide/06-student-perspective',
                },
            ],
        },
        {
            text: 'Advanced',
            items: [
                {
                    text: 'Security',
                    link: '/handbook/advanced/01-security',
                },
                {
                    text: 'Configuration',
                    link: '/handbook/advanced/02-configuration',
                },
                {
                    text: 'Running on a server',
                    link: '/handbook/advanced/03-on-server',
                },
                {
                    text: 'hakuctl',
                    link: '/handbook/advanced/04-hakuctl',
                },
            ],
        },
    ]
}
