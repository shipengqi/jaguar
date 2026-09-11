import { z } from 'zod'

export const createLoginSchema = (t: (key: string) => string) =>
  z.object({
    email: z.string().email(t('auth.login.emailError')),
    password: z.string().min(8, t('auth.login.passwordError')),
  })

export const createRegisterSchema = (t: (key: string) => string) =>
  z
    .object({
      name: z.string().min(2),
      email: z.string().email(t('auth.login.emailError')),
      password: z.string().min(8, t('auth.login.passwordError')),
      confirmPassword: z.string(),
    })
    .refine((data) => data.password === data.confirmPassword, {
      message: t('auth.register.confirmPasswordError'),
      path: ['confirmPassword'],
    })
