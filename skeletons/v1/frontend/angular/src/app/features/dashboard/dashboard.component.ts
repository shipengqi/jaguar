import { Component } from '@angular/core'
import { CommonModule } from '@angular/common'
import { TranslocoModule } from '@jsverse/transloco'
import { ThemeToggleComponent } from '@/app/shared/components/theme-toggle/theme-toggle.component'
import { LanguageSwitcherComponent } from '@/app/shared/components/language-switcher/language-switcher.component'

@Component({
  selector: 'app-dashboard',
  standalone: true,
  imports: [CommonModule, TranslocoModule, ThemeToggleComponent, LanguageSwitcherComponent],
  template: `
    <div class="flex h-screen">
      <div class="flex flex-1 flex-col">
        <header class="flex items-center justify-end gap-2 border-b px-6 py-3">
          <app-language-switcher />
          <app-theme-toggle />
        </header>
        <main class="flex-1 overflow-auto p-6">
          <h1 class="text-2xl font-bold">{{ 'dashboard.title' | transloco }}</h1>
        </main>
      </div>
    </div>
  `,
})
export class DashboardComponent {}
