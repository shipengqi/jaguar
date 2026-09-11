import { useDark, usePreferredDark } from '@vueuse/core'
import { ref, watch } from 'vue'

export type Theme = 'light' | 'dark' | 'system'

export function useTheme() {
  const isDark = useDark()
  const prefersDark = usePreferredDark()
  const theme = ref<Theme>('system')

  const setTheme = (newTheme: Theme) => {
    theme.value = newTheme
    localStorage.setItem('theme', newTheme)

    if (newTheme === 'system') {
      isDark.value = prefersDark.value
    } else {
      isDark.value = newTheme === 'dark'
    }
  }

  const saved = localStorage.getItem('theme') as Theme | null
  if (saved) setTheme(saved)

  watch(prefersDark, (val) => {
    if (theme.value === 'system') {
      isDark.value = val
    }
  })

  return { theme, isDark, setTheme }
}
