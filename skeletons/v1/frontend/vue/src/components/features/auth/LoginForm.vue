<script setup lang="ts">
import { useI18n } from 'vue-i18n'
import { useForm, useField } from 'vee-validate'
import { toTypedSchema } from '@vee-validate/zod'
import { z } from 'zod'

const { t } = useI18n()

const loginSchema = toTypedSchema(
  z.object({
    email: z.string().email(t('auth.login.emailError')),
    password: z.string().min(8, t('auth.login.passwordError')),
  })
)

const { handleSubmit, isSubmitting } = useForm({ validationSchema: loginSchema })
const { value: email, errorMessage: emailError } = useField<string>('email')
const { value: password, errorMessage: passwordError } = useField<string>('password')

const onSubmit = handleSubmit(async (_values) => {
  // TODO: call auth service
})
</script>

<template>
  <form class="space-y-4" @submit="onSubmit">
    <div>
      <label class="block text-sm font-medium">{{ t('auth.login.email') }}</label>
      <input
        v-model="email"
        type="email"
        :placeholder="t('auth.login.emailPlaceholder')"
        class="mt-1 w-full rounded border px-3 py-2"
      />
      <p v-if="emailError" class="mt-1 text-sm text-red-500">{{ emailError }}</p>
    </div>
    <div>
      <label class="block text-sm font-medium">{{ t('auth.login.password') }}</label>
      <input
        v-model="password"
        type="password"
        :placeholder="t('auth.login.passwordPlaceholder')"
        class="mt-1 w-full rounded border px-3 py-2"
      />
      <p v-if="passwordError" class="mt-1 text-sm text-red-500">{{ passwordError }}</p>
    </div>
    <button
      type="submit"
      class="w-full rounded bg-primary px-4 py-2 text-primary-foreground"
      :disabled="isSubmitting"
    >
      {{ isSubmitting ? t('auth.login.submitting') : t('auth.login.submit') }}
    </button>
  </form>
</template>
