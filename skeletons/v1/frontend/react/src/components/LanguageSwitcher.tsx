import { useTranslation } from 'react-i18next'
import i18n from '@/i18n'
import { Languages } from 'lucide-react'

type Locale = 'zh' | 'en' | 'ja'

const languages: { code: Locale; label: string }[] = [
  { code: 'zh', label: '中文' },
  { code: 'en', label: 'English' },
  { code: 'ja', label: '日本語' },
]

export function LanguageSwitcher() {
  const { i18n: i18nInstance } = useTranslation()

  const switchLanguage = (lang: Locale) => {
    i18n.changeLanguage(lang)
    localStorage.setItem('locale', lang)
  }

  return (
    <div className="relative group">
      <button
        className="inline-flex h-9 w-9 items-center justify-center rounded-md border border-input bg-background hover:bg-accent hover:text-accent-foreground"
      >
        <Languages className="h-4 w-4" />
      </button>
      <div className="absolute right-0 top-full mt-1 hidden w-28 rounded-md border bg-popover p-1 shadow-md group-hover:block">
        {languages.map((lang) => (
          <button
            key={lang.code}
            onClick={() => switchLanguage(lang.code)}
            className={`flex w-full items-center rounded-sm px-2 py-1.5 text-sm hover:bg-accent ${i18nInstance.language === lang.code ? 'font-semibold' : ''}`}
          >
            {lang.label}
          </button>
        ))}
      </div>
    </div>
  )
}
