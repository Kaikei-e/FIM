import { test, expect } from '@playwright/test';

test('go to landing page', async ({ page }) => {
  await page.goto('http://localhost:5173');
  const element = page.locator('id=fim-title');

  await expect(element).toBeVisible({ timeout: 1000 });
})
