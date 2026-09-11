<script setup lang="ts">
import { useI18n } from 'vue-i18n'
import ThemeToggle from '@/components/common/ThemeToggle/index.vue'
import LanguageSwitcher from '@/components/common/LanguageSwitcher/index.vue'

const { t } = useI18n()
</script>

<template>
  <div class="flex h-screen bg-background text-foreground">
    <div class="flex flex-1 flex-col">
      <header class="flex items-center justify-end gap-2 border-b px-6 py-3">
        <LanguageSwitcher />
        <ThemeToggle />
      </header>
      <main class="flex-1 overflow-auto p-6">
        <h1 class="text-2xl font-bold">{{ t('dashboard.title') }}</h1>
      </main>
    </div>
  </div>
</template>
