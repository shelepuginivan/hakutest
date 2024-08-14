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
    sidebar: {
        '/handbook/': { base: '/handbook/', items: sidebarHandbook() },
        '/reference/': { base: '/reference/', items: sidebarReference() },
    },
} satisfies DefaultTheme.Config

function nav(): DefaultTheme.NavItem[] {
    return [
        { text: 'Handbook', link: '/handbook/getting-started' },
        { text: 'Reference', link: '/reference/standards/test-schema' },
    ]
}

function sidebarHandbook(): DefaultTheme.SidebarItem[] {
    return [
        {
            text: 'Getting Started',
            link: 'getting-started',
        },
        {
            text: 'Install Hakutest',
            link: 'installation',
        },
        {
            text: 'Guide',
            base: '/handbook/guide/',
            items: [
                {
                    text: 'Server',
                    link: '01-server',
                },
                {
                    text: 'Dashboard',
                    link: '02-dashboard',
                },
                {
                    text: 'Tests',
                    link: '03-tests',
                },
                {
                    text: 'Results and Statistics',
                    link: '04-results-and-statistics',
                },
                {
                    text: 'Settings',
                    link: '05-settings',
                },
                {
                    text: 'Student Perspective',
                    link: '06-student-perspective',
                },
            ],
        },
        {
            text: 'Advanced',
            base: '/handbook/advanced/',
            items: [
                {
                    text: 'Security',
                    link: '01-security',
                },
                {
                    text: 'Configuration',
                    link: '02-configuration',
                },
                {
                    text: 'Running on a server',
                    link: '03-on-server',
                },
                {
                    text: 'hakuctl',
                    link: '04-hakuctl',
                },
                {
                    text: 'Log',
                    link: '05-log',
                },
            ],
        },
        {
            text: 'Troubleshooting',
            base: '/handbook/troubleshooting/',
            items: [
                {
                    text: 'Local IP detection',
                    link: '01-local-ip',
                },
            ],
        },
        {
            text: 'Reference',
            base: '/reference/',
            link: 'standards/test-schema',
        },
    ]
}

function sidebarReference(): DefaultTheme.SidebarItem[] {
    return [
        {
            text: 'Back to Handbook',
            base: '/handbook/',
            link: 'getting-started',
        },
        {
            text: 'Standards',
            base: '/reference/standards/',
            items: [
                {
                    text: 'Test JSON schema',
                    link: 'test-schema',
                },
                {
                    text: 'Result JSON schema',
                    link: 'result-schema',
                },
            ],
        },
        {
            text: 'Logotypes',
            base: '/',
            items: [
                {
                    text: 'Hakutest Icon (SVG)',
                    link: 'logo/icon.svg',
                    target: '_blank',
                },
                {
                    text: 'Hakutest Logo (SVG)',
                    link: 'logo/logo.svg',
                    target: '_blank',
                },
            ],
        },
    ]
}
