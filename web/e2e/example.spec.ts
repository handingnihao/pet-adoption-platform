import { test, expect } from '@playwright/test'

test.describe('基础功能测试', () => {
  test('访问首页', async ({ page }) => {
    await page.goto('/')
    
    // 验证页面标题或关键元素
    await expect(page).toHaveTitle(/宠物领养/)
  })

  test('导航到宠物列表', async ({ page }) => {
    await page.goto('/')
    
    // 点击宠物列表链接
    await page.click('a[href*="/pets"]')
    
    // 验证URL变化
    await expect(page).toHaveURL(/\/pets/)
  })
})
