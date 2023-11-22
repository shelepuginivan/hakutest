import type { Config, PluginConfig } from '@docusaurus/types'
import type { Options, ThemeConfig } from '@docusaurus/preset-classic'
import { themes as prismThemes } from 'prism-react-renderer'

const config: Config = {
    title: 'Hakutest',
    tagline: 'Reliable and efficient educational testing platform',
    favicon: 'img/favicon.ico',

    // Set the production url of your site here
    url: 'https://your-docusaurus-site.example.com',
    // Set the /<baseUrl>/ pathname under which your site is served
    // For GitHub pages deployment, it is often '/<projectName>/'
    baseUrl: '/',

    organizationName: 'shelepuginivan',
    projectName: 'hakutest',

    onBrokenLinks: 'throw',
    onBrokenMarkdownLinks: 'warn',

    // Even if you don't use internationalization, you can use this field to set
    // useful metadata like html lang. For example, if your site is Chinese, you
    // may want to replace "en" with "zh-Hans".
    i18n: {
        defaultLocale: 'en',
        locales: ['en'],
    },

    plugins: [
        [
            '@docusaurus/plugin-pwa',
            {
                pwaHead: [
                    {
                        tagName: 'link',
                        rel: 'manifest',
                        href: '/manifest.json',
                    },
                ],
            },
        ] satisfies PluginConfig,
    ],

    presets: [
        [
            'classic',
            {
                docs: {
                    sidebarPath: './sidebars.js',
                    editUrl:
                        'https://github.com/shelepuginivan/hakutest/tree/main/website',
                },
                blog: {
                    showReadingTime: true,
                    editUrl:
                        'https://github.com/shelepuginivan/hakutest/tree/main/website',
                },
                theme: {
                    customCss: './src/css/custom.css',
                },
                sitemap: {
                    changefreq: 'weekly',
                    priority: 0.5,
                    ignorePatterns: ['/tags/**'],
                    filename: 'sitemap.xml',
                },
            } satisfies Options,
        ],
    ],

    themeConfig: {
        image: 'img/hakutest-social-card.jpg',
        navbar: {
            title: 'Hakutest',
            logo: {
                alt: 'Hakutest Logo',
                src: 'img/logo.svg',
            },
            items: [
                {
                    type: 'docSidebar',
                    sidebarId: 'tutorialSidebar',
                    position: 'left',
                    label: 'Documentation',
                },
                { to: '/blog', label: 'Blog', position: 'left' },
                {
                    href: 'https://github.com/shelepuginivan/hakutest',
                    label: 'GitHub',
                    position: 'right',
                },
            ],
        },
        footer: {
            style: 'dark',
            links: [
                {
                    title: 'Docs',
                    items: [
                        {
                            label: 'Getting started',
                            to: '/docs/intro',
                        },
                    ],
                },
                {
                    title: 'More',
                    items: [
                        {
                            label: 'Blog',
                            to: '/blog',
                        },
                        {
                            label: 'GitHub',
                            href: 'https://github.com/shelepuginivan/hakutest',
                        },
                    ],
                },
            ],
            copyright: `Copyright © ${new Date().getFullYear()} Ivan Shelepugin. Built with Docusaurus.`,
        },
        prism: {
            theme: prismThemes.github,
            darkTheme: prismThemes.dracula,
        },
    } satisfies ThemeConfig,
}

export default config
