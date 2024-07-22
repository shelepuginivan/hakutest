import { defineConfig } from 'vitepress'
import { enThemeConfig } from './config/en.mts'
import { ruSearchConfig, ruThemeConfig } from './config/ru.mts'

export default defineConfig({
    lang: 'en-US',
    title: 'Hakutest',
    description: 'Modern and efficient educational testing',

    locales: {
        root: {
            label: 'English',
            lang: 'en',
            themeConfig: enThemeConfig,
        },
        ru: {
            label: 'Русский',
            lang: 'ru',
            themeConfig: ruThemeConfig,
        },
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
})
