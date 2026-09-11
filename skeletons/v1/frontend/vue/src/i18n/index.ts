import { createI18n } from 'vue-i18n'
import zh from './locales/zh.json'
import en from './locales/en.json'
import ja from './locales/ja.json'

export type Locale = 'zh' | 'en' | 'ja'

const savedLocale = localStorage.getItem('locale') as Locale | null

export const i18n = createI18n({
  legacy: false,
  locale: savedLocale || 'zh',
  fallbackLocale: 'en',
  messages: { zh, en, ja },
})
