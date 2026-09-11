import { useTheme } from '@/hooks/useTheme'
import { useTranslation } from 'react-i18next'
import { Sun, Moon, Monitor } from 'lucide-react'

type Theme = 'light' | 'dark' | 'system'

export function ThemeToggle() {
  const { theme, setTheme } = useTheme()
  const { t } = useTranslation()

  const options: { value: Theme; label: string; icon: React.ReactNode }[] = [
    { value: 'light', label: t('theme.light'), icon: <Sun className="mr-2 h-4 w-4" /> },
    { value: 'dark', label: t('theme.dark'), icon: <Moon className="mr-2 h-4 w-4" /> },
    { value: 'system', label: t('theme.system'), icon: <Monitor className="mr-2 h-4 w-4" /> },
  ]

  return (
    <div className="relative group">
      <button
        aria-label={t('theme.toggle')}
        className="inline-flex h-9 w-9 items-center justify-center rounded-md border border-input bg-background hover:bg-accent hover:text-accent-foreground"
      >
        {theme === 'dark' ? (
          <Moon className="h-4 w-4" />
        ) : (
          <Sun className="h-4 w-4" />
        )}
      </button>
      <div className="absolute right-0 top-full mt-1 hidden w-36 rounded-md border bg-popover p-1 shadow-md group-hover:block">
        {options.map((opt) => (
          <button
            key={opt.value}
            onClick={() => setTheme(opt.value)}
            className={`flex w-full items-center rounded-sm px-2 py-1.5 text-sm hover:bg-accent ${theme === opt.value ? 'font-semibold' : ''}`}
          >
            {opt.icon}
            {opt.label}
          </button>
        ))}
      </div>
    </div>
  )
}
