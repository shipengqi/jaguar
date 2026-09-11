import { create } from 'zustand'
import { persist } from 'zustand/middleware'
import { immer } from 'zustand/middleware/immer'

type Theme = 'light' | 'dark' | 'system'
type Language = 'zh' | 'en' | 'ja'

interface UIState {
  theme: Theme
  language: Language
  sidebarOpen: boolean
  setTheme: (theme: Theme) => void
  setLanguage: (language: Language) => void
  toggleSidebar: () => void
}

export const useUIStore = create<UIState>()(
  persist(
    immer((set) => ({
      theme: 'system',
      language: 'zh',
      sidebarOpen: true,

      setTheme: (theme) =>
        set((state) => {
          state.theme = theme
        }),

      setLanguage: (language) =>
        set((state) => {
          state.language = language
        }),

      toggleSidebar: () =>
        set((state) => {
          state.sidebarOpen = !state.sidebarOpen
        }),
    })),
    { name: 'ui-storage' }
  )
)
