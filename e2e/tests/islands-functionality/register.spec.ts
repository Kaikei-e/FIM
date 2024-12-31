import { test, expect } from '@playwright/test';

test('register a new feed', {tag:'@wip'}, async ({ page }) => {
  await page.goto('http://localhost:5173/home');

  const [dialog] = await Promise.all([
    page.waitForEvent('dialog'),
    page.click('button#registerButton'),
  ]);

  expect(dialog.message()).toContain('The feed has been registered successfully');
  await dialog.accept();
  await expect(page.locator('button#registerButton')).toBeVisible();
});