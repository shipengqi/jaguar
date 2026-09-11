import { createApp } from 'vue'
import { createPinia } from 'pinia'
import piniaPluginPersistedstate from '@pinia/plugin-persistedstate'
import { VueQueryPlugin } from '@tanstack/vue-query'
import router from './router'
import { i18n } from './i18n'
import App from './App.vue'
import './assets/globals.css'

const pinia = createPinia()
pinia.use(piniaPluginPersistedstate)

const app = createApp(App)
app.use(pinia)
app.use(router)
app.use(VueQueryPlugin)
app.use(i18n)
app.mount('#app')
