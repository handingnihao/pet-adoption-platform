# 宠物模块API测试脚本

Write-Host "`n========================================" -ForegroundColor Cyan
Write-Host "   宠物模块API完整测试" -ForegroundColor Cyan
Write-Host "========================================`n" -ForegroundColor Cyan

# 1. 用户登录获取Token
Write-Host "【1】用户登录..." -ForegroundColor Yellow
$loginBody = '{"username":"testuser1","password":"Test@123456"}'
$loginResponse = Invoke-WebRequest -Uri http://localhost:8080/api/v1/users/login -Method POST -Body $loginBody -ContentType 'application/json'
$loginData = $loginResponse.Content | ConvertFrom-Json
$token = $loginData.data.token
Write-Host "✅ 登录成功，Token已获取" -ForegroundColor Green

# 2. 测试创建宠物
Write-Host "`n【2】创建宠物..." -ForegroundColor Yellow
$headers = @{
    "Authorization" = "Bearer $token"
    "Content-Type" = "application/json"
}
$petBody = @"
{
  "name": "小黑",
  "type": "dog",
  "breed": "拉布拉多",
  "gender": "male",
  "age": 24,
  "size": "large",
  "color": "黑色",
  "weight": 30.5,
  "is_vaccinated": true,
  "is_sterilized": false,
  "health_status": "健康良好",
  "description": "小黑是一只非常活泼可爱的拉布拉多犬，性格温顺，喜欢和人玩耍。",
  "character": "温顺、活泼、聪明",
  "photos": ["https://example.com/dog1.jpg", "https://example.com/dog2.jpg"],
  "cover_photo": "https://example.com/dog-cover.jpg",
  "province": "广东省",
  "city": "深圳市",
  "district": "南山区",
  "address": "科技园南区"
}
"@

try {
    $createResponse = Invoke-WebRequest -Uri http://localhost:8080/api/v1/pets -Method POST -Headers $headers -Body $petBody
    $createData = $createResponse.Content | ConvertFrom-Json
    $petId = $createData.data.id
    Write-Host "✅ 创建成功！宠物ID: $petId" -ForegroundColor Green
    Write-Host "   提示: $($createData.data.message)" -ForegroundColor Gray
} catch {
    Write-Host "❌ 创建失败: $($_.Exception.Message)" -ForegroundColor Red
    $petId = $null
}

# 3. 测试获取宠物列表
Write-Host "`n【3】获取宠物列表..." -ForegroundColor Yellow
try {
    $listResponse = Invoke-WebRequest -Uri "http://localhost:8080/api/v1/pets?page=1&page_size=10" -Method GET
    $listData = $listResponse.Content | ConvertFrom-Json
    Write-Host "✅ 获取成功！总数: $($listData.data.pagination.total)" -ForegroundColor Green
    if ($listData.data.list.Count -gt 0) {
        Write-Host "`n宠物列表（前5个）:" -ForegroundColor Cyan
        $listData.data.list | Select-Object -First 5 | Format-Table -Property id,name,type,breed,city,status -AutoSize
    }
} catch {
    Write-Host "❌ 获取失败: $($_.Exception.Message)" -ForegroundColor Red
}

# 4. 测试获取宠物详情
if ($petId) {
    Write-Host "`n【4】获取宠物详情..." -ForegroundColor Yellow
    try {
        $detailResponse = Invoke-WebRequest -Uri "http://localhost:8080/api/v1/pets/$petId" -Method GET
        $detailData = $detailResponse.Content | ConvertFrom-Json
        Write-Host "✅ 获取成功！" -ForegroundColor Green
        Write-Host "   名称: $($detailData.data.name)" -ForegroundColor Cyan
        Write-Host "   品种: $($detailData.data.breed)" -ForegroundColor Cyan
        Write-Host "   年龄: $($detailData.data.age)个月" -ForegroundColor Cyan
        Write-Host "   状态: $($detailData.data.status)" -ForegroundColor Cyan
        Write-Host "   浏览次数: $($detailData.data.view_count)" -ForegroundColor Cyan
    } catch {
        Write-Host "❌ 获取失败: $($_.Exception.Message)" -ForegroundColor Red
    }
}

# 5. 测试获取我的宠物
Write-Host "`n【5】获取我的宠物..." -ForegroundColor Yellow
try {
    $myPetsResponse = Invoke-WebRequest -Uri "http://localhost:8080/api/v1/pets/my?page=1&page_size=10" -Method GET -Headers $headers
    $myPetsData = $myPetsResponse.Content | ConvertFrom-Json
    Write-Host "✅ 获取成功！我发布的宠物数: $($myPetsData.data.pagination.total)" -ForegroundColor Green
} catch {
    Write-Host "❌ 获取失败: $($_.Exception.Message)" -ForegroundColor Red
}

# 6. 测试条件查询
Write-Host "`n【6】测试条件查询（狗类型）..." -ForegroundColor Yellow
try {
    $queryResponse = Invoke-WebRequest -Uri "http://localhost:8080/api/v1/pets/query?type=dog&page=1&page_size=10" -Method GET
    $queryData = $queryResponse.Content | ConvertFrom-Json
    Write-Host "✅ 查询成功！狗的数量: $($queryData.data.pagination.total)" -ForegroundColor Green
} catch {
    Write-Host "❌ 查询失败: $($_.Exception.Message)" -ForegroundColor Red
}

# 7. 测试搜索
Write-Host "`n【7】测试搜索功能..." -ForegroundColor Yellow
try {
    $searchResponse = Invoke-WebRequest -Uri "http://localhost:8080/api/v1/pets/search?keyword=拉布拉多&page=1&page_size=10" -Method GET
    $searchData = $searchResponse.Content | ConvertFrom-Json
    Write-Host "✅ 搜索成功！找到 $($searchData.data.pagination.total) 个结果" -ForegroundColor Green
} catch {
    Write-Host "❌ 搜索失败: $($_.Exception.Message)" -ForegroundColor Red
}

# 8. 测试推荐宠物
Write-Host "`n【8】获取推荐宠物..." -ForegroundColor Yellow
try {
    $recommendResponse = Invoke-WebRequest -Uri "http://localhost:8080/api/v1/pets/recommended?limit=5" -Method GET
    $recommendData = $recommendResponse.Content | ConvertFrom-Json
    Write-Host "✅ 获取成功！推荐宠物数: $($recommendData.data.Count)" -ForegroundColor Green
} catch {
    Write-Host "❌ 获取失败: $($_.Exception.Message)" -ForegroundColor Red
}

# 9. 管理员功能测试
Write-Host "`n【9】管理员功能测试..." -ForegroundColor Yellow
Write-Host "  登录管理员账号..." -ForegroundColor Gray
$adminLoginBody = '{"username":"admin","password":"Test@123456"}'
try {
    $adminLoginResponse = Invoke-WebRequest -Uri http://localhost:8080/api/v1/users/login -Method POST -Body $adminLoginBody -ContentType 'application/json'
    $adminLoginData = $adminLoginResponse.Content | ConvertFrom-Json
    $adminToken = $adminLoginData.data.token
    Write-Host "  ✅ 管理员登录成功" -ForegroundColor Green
    
    $adminHeaders = @{
        "Authorization" = "Bearer $adminToken"
        "Content-Type" = "application/json"
    }
    
    # 获取待审核宠物
    try {
        $pendingResponse = Invoke-WebRequest -Uri "http://localhost:8080/api/v1/pets/pending?page=1&page_size=10" -Method GET -Headers $adminHeaders
        $pendingData = $pendingResponse.Content | ConvertFrom-Json
        Write-Host "  ✅ 待审核宠物: $($pendingData.data.pagination.total) 个" -ForegroundColor Green
        
        # 如果有待审核的宠物，测试审核功能
        if ($petId) {
            Write-Host "  测试审核通过..." -ForegroundColor Gray
            try {
                $approveResponse = Invoke-WebRequest -Uri "http://localhost:8080/api/v1/pets/$petId/approve" -Method PUT -Headers $adminHeaders
                Write-Host "  ✅ 审核通过成功！" -ForegroundColor Green
            } catch {
                Write-Host "  ⚠️  审核操作: $($_.Exception.Message)" -ForegroundColor Yellow
            }
        }
    } catch {
        Write-Host "  ❌ 管理员功能失败: $($_.Exception.Message)" -ForegroundColor Red
    }
    
    # 获取统计信息
    try {
        $statsResponse = Invoke-WebRequest -Uri "http://localhost:8080/api/v1/pets/statistics" -Method GET -Headers $adminHeaders
        $statsData = $statsResponse.Content | ConvertFrom-Json
        Write-Host "  ✅ 统计信息:" -ForegroundColor Green
        Write-Host "     总数: $($statsData.data.total)" -ForegroundColor Cyan
        Write-Host "     待审核: $($statsData.data.pending)" -ForegroundColor Cyan
        Write-Host "     可领养: $($statsData.data.available)" -ForegroundColor Cyan
        Write-Host "     已领养: $($statsData.data.adopted)" -ForegroundColor Cyan
    } catch {
        Write-Host "  ❌ 统计信息获取失败" -ForegroundColor Red
    }
} catch {
    Write-Host "  ❌ 管理员登录失败: $($_.Exception.Message)" -ForegroundColor Red
}

# 总结
Write-Host "`n========================================" -ForegroundColor Cyan
Write-Host "   测试完成！" -ForegroundColor Cyan
Write-Host "========================================" -ForegroundColor Cyan
