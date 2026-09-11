import { Component } from '@angular/core'
import { CommonModule } from '@angular/common'
import { ReactiveFormsModule, FormBuilder, Validators } from '@angular/forms'
import { TranslocoModule } from '@jsverse/transloco'
import { z } from 'zod'

@Component({
  selector: 'app-login',
  standalone: true,
  imports: [CommonModule, ReactiveFormsModule, TranslocoModule],
  template: `
    <div class="flex min-h-screen items-center justify-center">
      <div class="w-full max-w-md space-y-8 p-8">
        <form [formGroup]="form" (ngSubmit)="onSubmit()" class="space-y-4">
          <div>
            <label class="block text-sm font-medium">{{ 'auth.login.email' | transloco }}</label>
            <input
              type="email"
              formControlName="email"
              class="mt-1 w-full rounded border px-3 py-2"
            />
            <p *ngIf="form.get('email')?.invalid && form.get('email')?.touched" class="mt-1 text-sm text-red-500">
              {{ 'auth.login.emailError' | transloco }}
            </p>
          </div>
          <div>
            <label class="block text-sm font-medium">{{ 'auth.login.password' | transloco }}</label>
            <input
              type="password"
              formControlName="password"
              class="mt-1 w-full rounded border px-3 py-2"
            />
            <p *ngIf="form.get('password')?.invalid && form.get('password')?.touched" class="mt-1 text-sm text-red-500">
              {{ 'auth.login.passwordError' | transloco }}
            </p>
          </div>
          <button
            type="submit"
            [disabled]="form.invalid || submitting"
            class="w-full rounded bg-primary px-4 py-2 text-primary-foreground"
          >
            {{ submitting ? ('auth.login.submitting' | transloco) : ('auth.login.submit' | transloco) }}
          </button>
        </form>
      </div>
    </div>
  `,
})
export class LoginComponent {
  submitting = false

  form = this.fb.group({
    email: ['', [Validators.required, Validators.email]],
    password: ['', [Validators.required, Validators.minLength(8)]],
  })

  constructor(private fb: FormBuilder) {}

  async onSubmit() {
    if (this.form.invalid) return
    this.submitting = true
    // TODO: call auth service
    this.submitting = false
  }
}
