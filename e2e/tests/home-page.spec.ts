import {test, expect} from '@playwright/test';

test("check home page's title visiblity", async ({page}) => {
  await page.goto('http://localhost:5173/home');
  const title = await page.locator('id=fim-home-title');

  await expect(title).toBeVisible({timeout: 1000});
})