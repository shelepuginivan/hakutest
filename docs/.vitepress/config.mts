import { defineConfig } from 'vitepress'

export default defineConfig({
  lang: 'en-US',
  title: 'Hakutest',
  description: 'Modern and efficient educational testing',

  themeConfig: {
    siteTitle: 'Hakutest',
    
    nav: [
      { text: 'Home', link: '/' },
    ],

    sidebar: [
      {
        text: 'Examples',
        items: [
          { text: 'Markdown Examples', link: '/markdown-examples' },
          { text: 'Runtime API Examples', link: '/api-examples' }
        ]
      }
    ],

    socialLinks: [
      { icon: 'github', link: 'https://github.com/shelepuginivan/hakutest' }
    ]
  }
})
