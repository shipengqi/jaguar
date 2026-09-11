import { test, expect } from '@playwright/test'

test.describe('Language Switching', () => {
  test('should switch to English', async ({ page }) => {
    await page.goto('/')
    await page.getByText('English').click()
    await expect(page).toHaveURL(/\/en\//)
  })

  test('should switch to Japanese', async ({ page }) => {
    await page.goto('/')
    await page.getByText('日本語').click()
    await expect(page).toHaveURL(/\/ja\//)
  })
})
