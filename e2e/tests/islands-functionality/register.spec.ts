import { test, expect } from '@playwright/test';

test('register a new feed', {tag:'@wip'}, async ({ page }) => {
  await page.goto('http://localhost:5173/home');
  page.on('dialog', async (dialog) => {
    expect(dialog.message()).toContain('The feed has been registered successfully');
    await dialog.accept();
  });

  await expect(page.locator('button#registerButton')).toBeVisible();

  await page.locator('button#registerButton').click();

  await page.isVisible('button#registerButton');
});