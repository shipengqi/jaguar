import { defineStore } from 'pinia'
import { ref } from 'vue'

export type Theme = 'light' | 'dark' | 'system'
export type Language = 'zh' | 'en' | 'ja'

export const useUIStore = defineStore(
  'ui',
  () => {
    const theme = ref<Theme>('system')
    const language = ref<Language>('zh')
    const sidebarOpen = ref(true)

    function setTheme(t: Theme) { theme.value = t }
    function setLanguage(l: Language) { language.value = l }
    function toggleSidebar() { sidebarOpen.value = !sidebarOpen.value }

    return { theme, language, sidebarOpen, setTheme, setLanguage, toggleSidebar }
  },
  { persist: true }
)
