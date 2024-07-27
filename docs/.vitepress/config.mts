import { defineConfig, TransformContext } from 'vitepress'
import { enThemeConfig } from './config/en.mts'
import { ruSearchConfig, ruThemeConfig } from './config/ru.mts'

const siteURL = 'https://hakutest.org'

function getOpenGraphImage(ctx: TransformContext): string {
    const page = ctx.page
    const lang = ctx.siteData.lang

    if (page.includes('handbook')) {
        return `${siteURL}/social/${lang}/handbook.png`
    }

    if (page.includes('reference')) {
        return `${siteURL}/social/${lang}/reference.png`
    }

    return `${siteURL}/social/${lang}/index.png`
}

export default defineConfig({
    lang: 'en-US',
    title: 'Hakutest',
    description: 'Modern and efficient educational testing',

    locales: {
        root: {
            label: 'English',
            lang: 'en',
            themeConfig: enThemeConfig,

            head: [['meta', { name: 'og:locale', content: 'en' }]],
        },
        ru: {
            label: 'Русский',
            lang: 'ru',
            themeConfig: ruThemeConfig,

            head: [['meta', { name: 'og:locale', content: 'ru' }]],
        },
    },

    head: [
        ['meta', { name: 'og:site_name', content: 'hakutest' }],
        ['meta', { name: 'twitter:card', content: 'summary_large_image' }],
    ],

    transformHead(ctx) {
        return [
            ['meta', { name: 'og:title', content: `${ctx.title}` }],
            ['meta', { name: 'og:description', content: `${ctx.description}` }],
            ['meta', { name: 'og:image', content: getOpenGraphImage(ctx) }],
            ...ctx.head,
        ]
    },

    themeConfig: {
        siteTitle: 'Hakutest',
        logo: '/hakutest.svg',

        search: {
            provider: 'local',
            options: {
                locales: {
                    ru: ruSearchConfig,
                },
            },
        },

        socialLinks: [
            {
                icon: 'github',
                link: 'https://github.com/shelepuginivan/hakutest',
            },
        ],
    },

    sitemap: {
        hostname: siteURL,
    },
})
