import '@testing-library/jest-dom'
import { config } from '@vue/test-utils'
import { createI18n } from 'vue-i18n'

const i18n = createI18n({ legacy: false, locale: 'zh', messages: { zh: {} } })
config.global.plugins = [i18n]
