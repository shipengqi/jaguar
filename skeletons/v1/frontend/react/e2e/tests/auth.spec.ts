import { test, expect } from '@playwright/test'

test.describe('Authentication Flow', () => {
  test.beforeEach(async ({ page }) => {
    await page.goto('/login')
  })

  test('should redirect to login when unauthenticated', async ({ page }) => {
    await page.goto('/dashboard')
    await expect(page).toHaveURL('/login')
  })

  test('should show validation errors on empty submit', async ({ page }) => {
    await page.getByRole('button', { name: /登录/i }).click()
    await expect(page.getByText(/有效邮箱|valid email/i)).toBeVisible()
  })
})
