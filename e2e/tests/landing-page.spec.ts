import { test, expect } from '@playwright/test';

test('go to landing page', async ({ page }) => {
  await page.goto('http://localhost:5173');
  const element = page.locator('id=fim-title');

  await expect(element).toBeVisible({ timeout: 1000 });
})

test('go to landing page and click get started', async ({ page }) => {
  await page.goto('http://localhost:5173');
  await page.click('id=get-started');
  const url = page.url();

  expect(url).toBe('http://localhost:5173/home');
})