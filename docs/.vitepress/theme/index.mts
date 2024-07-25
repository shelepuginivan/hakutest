import DefaultTheme from 'vitepress/theme'
import './colors.css'
import './guide.css'
import './iconfont.css'

import { h } from 'vue'
import HomeFeaturesBefore from './components/HomeFeaturesBefore.vue'

export default {
    extends: DefaultTheme,
    Layout() {
        return h(DefaultTheme.Layout, null, {
            'home-features-before': () => h(HomeFeaturesBefore),
        })
    },
}
