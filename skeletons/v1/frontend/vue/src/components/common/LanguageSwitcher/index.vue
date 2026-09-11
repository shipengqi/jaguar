<script setup lang="ts">
import { useI18n } from 'vue-i18n'
import type { Locale } from '@/i18n'

const { locale } = useI18n()

const languages = [
  { code: 'zh' as Locale, label: '中文' },
  { code: 'en' as Locale, label: 'English' },
  { code: 'ja' as Locale, label: '日本語' },
]

const switchLanguage = (lang: Locale) => {
  locale.value = lang
  localStorage.setItem('locale', lang)
}
</script>

<template>
  <div class="flex gap-1">
    <button
      v-for="lang in languages"
      :key="lang.code"
      :class="['px-2 py-1 text-sm rounded', locale === lang.code ? 'font-bold' : '']"
      @click="switchLanguage(lang.code)"
    >
      {{ lang.label }}
    </button>
  </div>
</template>
