import type { Config, I18nConfig, PluginConfig } from '@docusaurus/types'
import type { Options, ThemeConfig } from '@docusaurus/preset-classic'
import { themes as prismThemes } from 'prism-react-renderer'
import { EnumChangefreq } from 'sitemap'

const config: Config = {
    title: 'Hakutest',
    tagline: 'Reliable and efficient educational testing platform',
    favicon: 'img/favicon.ico',

    url: 'https://hakutest.shelepugin.ru',
    baseUrl: '/',

    organizationName: 'shelepuginivan',
    projectName: 'hakutest',

    onBrokenLinks: 'throw',
    onBrokenMarkdownLinks: 'warn',

    i18n: {
        defaultLocale: 'en',
        path: 'i18n',
        locales: ['en', 'ru'],
        localeConfigs: {
            ru: {
                label: 'Русский',
                htmlLang: 'ru',
            },
        },
    } satisfies I18nConfig,

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
                    changefreq: EnumChangefreq.WEEKLY,
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
                    type: 'localeDropdown',
                    position: 'right',
                },
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
                    title: 'Development',
                    items: [
                        {
                            label: 'GitHub',
                            href: 'https://github.com/shelepuginivan/hakutest',
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
                            html: `
              <a href="https://www.netlify.com" target="_blank" rel="noreferrer noopener" aria-label="Deploys by Netlify">
                  <img src="https://www.netlify.com/img/global/badges/netlify-color-accent.svg" alt="Deploys by Netlify" width="114" height="51" />
              </a>`,
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
