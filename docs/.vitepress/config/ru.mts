import { DefaultTheme } from 'vitepress'

export const ruThemeConfig = {
    darkModeSwitchLabel: 'Оформление',
    lightModeSwitchTitle: 'Переключить на светлую тему',
    darkModeSwitchTitle: 'Переключить на тёмную тему',
    sidebarMenuLabel: 'Меню',
    returnToTopLabel: 'Вернуться к началу',
    langMenuLabel: 'Изменить язык',

    outline: { label: 'Содержание страницы' },

    lastUpdated: {
        text: 'Обновлено',
    },

    docFooter: {
        prev: 'Предыдущая страница',
        next: 'Следующая страница',
    },

    notFound: {
        title: 'Страница не найдена',
        quote: 'Хакутаку возвращаются, чтобы спасти образование',
        linkText: 'Вернуться на главную страницу',
    },

    footer: {
        message:
            'Hakutest распространяется под <a href="https://github.com/shelepuginivan/hakutest" target="_blank">лицензией MIT</a>.',
        copyright: '© 2024 – настоящее время, Иван Шелепугин',
    },

    nav: nav(),
    sidebar: sidebar(),
} satisfies DefaultTheme.Config

export const ruSearchConfig: Partial<
    Omit<DefaultTheme.LocalSearchOptions, 'locales'>
> = {
    translations: {
        button: {
            buttonText: 'Поиск',
            buttonAriaLabel: 'Поиск',
        },
        modal: {
            displayDetails: 'Отобразить подробный список',
            resetButtonTitle: 'Сбросить поиск',
            backButtonTitle: 'Закрыть поиск',
            noResultsText: 'Нет результатов по запросу',
            footer: {
                selectText: 'выбрать',
                selectKeyAriaLabel: 'выбрать',
                navigateText: 'перейти',
                navigateUpKeyAriaLabel: 'стрелка вверх',
                navigateDownKeyAriaLabel: 'стрелка вниз',
                closeText: 'закрыть',
                closeKeyAriaLabel: 'esc',
            },
        },
    },
}

function nav(): DefaultTheme.NavItem[] {
    return [
        { text: 'Главная', link: '/ru' },
        {
            text: 'Руководство',
            link: '/ru/handbook/getting-started',
        },
    ]
}

function sidebar(): DefaultTheme.SidebarItem[] {
    return [
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
                {
                    text: 'Тесты',
                    link: '/ru/handbook/guide/03-tests',
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
    ]
}
