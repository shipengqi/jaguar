import { test, expect } from '@playwright/test'

test.describe('Theme Switching', () => {
  test('should switch to dark mode', async ({ page }) => {
    await page.goto('/')
    await page.getByRole('button', { name: /切换主题|toggle theme/i }).click()
    await expect(page.locator('html')).toHaveClass(/dark/)
  })
})
