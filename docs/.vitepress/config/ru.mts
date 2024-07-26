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
    sidebar: {
        '/ru/handbook/': { base: '/ru/handbook/', items: sidebarHandbook() },
    },
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
        {
            text: 'Главная',
            link: '/ru',
        },
        {
            text: 'Руководство',
            link: '/ru/handbook/getting-started',
        },
    ]
}

function sidebarHandbook(): DefaultTheme.SidebarItem[] {
    return [
        {
            text: 'Введение',
            link: 'getting-started',
        },
        {
            text: 'Установка',
            link: 'installation',
        },
        {
            text: 'Гайд',
            base: '/ru/handbook/guide/',
            items: [
                {
                    text: 'Сервер',
                    link: '01-server',
                },
                {
                    text: 'Панель управления',
                    link: '02-dashboard',
                },
                {
                    text: 'Тесты',
                    link: '03-tests',
                },
                {
                    text: 'Результаты и статистика',
                    link: '04-results-and-statistics',
                },
                {
                    text: 'Настройки',
                    link: '05-settings',
                },
                {
                    text: 'С точки зрения ученика',
                    link: '06-student-perspective',
                },
            ],
        },
        {
            text: 'Продвинутый гайд',
            base: '/ru/handbook/advanced/',
            items: [
                {
                    text: 'Безопасность',
                    link: '01-security',
                },
                {
                    text: 'Конфигурация',
                    link: '02-configuration',
                },
                {
                    text: 'Запуск на сервере',
                    link: '03-on-server',
                },
                {
                    text: 'hakuctl',
                    link: '04-hakuctl',
                },
                {
                    text: 'Журнал',
                    link: '05-log',
                },
            ],
        },
        {
            text: 'Справочная информация',
            base: '/ru/reference/',
            link: 'index',
        },
    ]
}
