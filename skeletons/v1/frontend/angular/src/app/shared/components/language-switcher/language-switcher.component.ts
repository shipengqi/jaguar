import { Component } from '@angular/core'
import { CommonModule } from '@angular/common'
import { TranslocoService } from '@jsverse/transloco'

@Component({
  selector: 'app-language-switcher',
  standalone: true,
  imports: [CommonModule],
  template: `
    <div class="flex gap-1">
      <button
        *ngFor="let lang of languages"
        [class.font-bold]="currentLang === lang.code"
        class="px-2 py-1 text-sm rounded"
        (click)="switchLang(lang.code)"
      >
        {{ lang.label }}
      </button>
    </div>
  `,
})
export class LanguageSwitcherComponent {
  languages = [
    { code: 'zh', label: '中文' },
    { code: 'en', label: 'English' },
    { code: 'ja', label: '日本語' },
  ]

  get currentLang() {
    return this.translocoService.getActiveLang()
  }

  constructor(private translocoService: TranslocoService) {}

  switchLang(lang: string) {
    this.translocoService.setActiveLang(lang)
    localStorage.setItem('locale', lang)
  }
}
