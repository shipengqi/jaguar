<script setup lang="ts">
import { useTheme } from '@/composables/useTheme'
import { useI18n } from 'vue-i18n'
import { Sun, Moon, Monitor } from 'lucide-vue-next'

const { theme, setTheme } = useTheme()
const { t } = useI18n()
</script>

<template>
  <div class="relative">
    <button
      class="inline-flex items-center justify-center rounded-md border px-3 py-2"
      :aria-label="t('theme.toggle')"
    >
      <Sun v-if="theme !== 'dark'" class="h-4 w-4" />
      <Moon v-else class="h-4 w-4" />
    </button>
    <div class="absolute right-0 mt-1 rounded border bg-background shadow">
      <button class="flex w-full items-center gap-2 px-3 py-2 text-sm" @click="setTheme('light')">
        <Sun class="h-4 w-4" />{{ t('theme.light') }}
      </button>
      <button class="flex w-full items-center gap-2 px-3 py-2 text-sm" @click="setTheme('dark')">
        <Moon class="h-4 w-4" />{{ t('theme.dark') }}
      </button>
      <button class="flex w-full items-center gap-2 px-3 py-2 text-sm" @click="setTheme('system')">
        <Monitor class="h-4 w-4" />{{ t('theme.system') }}
      </button>
    </div>
  </div>
</template>
