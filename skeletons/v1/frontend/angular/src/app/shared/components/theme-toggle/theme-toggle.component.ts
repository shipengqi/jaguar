import { Component } from '@angular/core'
import { CommonModule } from '@angular/common'
import { TranslocoModule, TranslocoService } from '@jsverse/transloco'
import { ThemeService, type Theme } from '@/app/core/services/theme.service'

@Component({
  selector: 'app-theme-toggle',
  standalone: true,
  imports: [CommonModule, TranslocoModule],
  template: `
    <div class="relative">
      <button
        class="inline-flex items-center justify-center rounded-md border px-3 py-2"
        [attr.aria-label]="'theme.toggle' | transloco"
        (click)="toggleMenu()"
      >
        ☀️
      </button>
      <div *ngIf="menuOpen" class="absolute right-0 mt-1 rounded border bg-white shadow dark:bg-gray-800">
        <button class="flex w-full items-center gap-2 px-3 py-2 text-sm" (click)="setTheme('light')">
          {{ 'theme.light' | transloco }}
        </button>
        <button class="flex w-full items-center gap-2 px-3 py-2 text-sm" (click)="setTheme('dark')">
          {{ 'theme.dark' | transloco }}
        </button>
        <button class="flex w-full items-center gap-2 px-3 py-2 text-sm" (click)="setTheme('system')">
          {{ 'theme.system' | transloco }}
        </button>
      </div>
    </div>
  `,
})
export class ThemeToggleComponent {
  menuOpen = false

  constructor(private themeService: ThemeService) {}

  toggleMenu() {
    this.menuOpen = !this.menuOpen
  }

  setTheme(theme: Theme) {
    this.themeService.setTheme(theme)
    this.menuOpen = false
  }
}
