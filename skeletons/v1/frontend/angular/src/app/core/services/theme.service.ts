import { Injectable, signal, effect } from '@angular/core'

export type Theme = 'light' | 'dark' | 'system'

@Injectable({ providedIn: 'root' })
export class ThemeService {
  theme = signal<Theme>('system')

  constructor() {
    const saved = localStorage.getItem('theme') as Theme | null
    if (saved) this.theme.set(saved)

    effect(() => {
      this.applyTheme(this.theme())
    })

    window.matchMedia('(prefers-color-scheme: dark)').addEventListener('change', () => {
      if (this.theme() === 'system') {
        this.applyTheme('system')
      }
    })
  }

  setTheme(theme: Theme) {
    this.theme.set(theme)
    localStorage.setItem('theme', theme)
  }

  private applyTheme(theme: Theme) {
    const isDark =
      theme === 'dark' ||
      (theme === 'system' && window.matchMedia('(prefers-color-scheme: dark)').matches)
    document.documentElement.classList.toggle('dark', isDark)
  }
}
