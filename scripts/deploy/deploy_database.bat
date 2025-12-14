@echo off
chcp 65001 >nul
REM ============================================
REM 爱心宠物领养平台 - 数据库一键部署脚本 (Windows)
REM MySQL 8.4
REM ============================================

setlocal enabledelayedexpansion

echo.
echo ==================================================
echo 🐾 爱心宠物领养平台 - 数据库部署工具
echo ==================================================
echo.

REM 检查 MySQL 是否安装
where mysql >nul 2>nul
if %ERRORLEVEL% NEQ 0 (
    echo ❌ 错误: 未找到 mysql 命令，请确认 MySQL 已安装并添加到 PATH
    pause
    exit /b 1
)

echo ✅ MySQL 已安装
mysql --version
echo.

REM 设置脚本目录
set SCRIPT_DIR=%~dp0..\migration
echo 📁 脚本目录: %SCRIPT_DIR%
echo.

REM 检查SQL脚本文件
if not exist "%SCRIPT_DIR%\01_create_user_and_database.sql" (
    echo ❌ 错误: 找不到SQL脚本文件
    pause
    exit /b 1
)

REM 获取 root 密码
echo 📝 请输入 MySQL Root 密码:
set /p ROOT_PASSWORD="Root密码: "
echo.

REM 测试连接
echo 🔍 测试 MySQL 连接...
mysql -u root -p%ROOT_PASSWORD% -e "SELECT 1;" >nul 2>nul
if %ERRORLEVEL% NEQ 0 (
    echo ❌ 无法连接到 MySQL，请检查密码
    pause
    exit /b 1
)
echo ✅ MySQL 连接成功
echo.

REM 确认部署
echo ==================================================
echo ⚠️  即将开始部署
echo ==================================================
echo.
echo 这将执行以下操作：
echo   1. 创建数据库 pet_adoption
echo   2. 创建用户 pet_admin
echo   3. 创建 10 个数据表
echo   4. 插入初始数据（包括管理员账户）
echo.
set /p CONFIRM="是否继续？(Y/N): "

if /i not "%CONFIRM%"=="Y" (
    echo ⚠️  部署已取消
    pause
    exit /b 0
)

echo.
echo ==================================================
echo 步骤 1/3: 创建用户和数据库
echo ==================================================
echo.

mysql -u root -p%ROOT_PASSWORD% < "%SCRIPT_DIR%\01_create_user_and_database.sql"
if %ERRORLEVEL% NEQ 0 (
    echo ❌ 创建数据库失败
    pause
    exit /b 1
)
echo ✅ 数据库和用户创建成功
echo.

REM 获取应用用户密码
echo 📝 请输入 pet_admin 用户密码
echo    (直接回车使用默认密码: Pet@Admin#2024!Secure)
set /p APP_PASSWORD="密码: "
if "%APP_PASSWORD%"=="" (
    set APP_PASSWORD=Pet@Admin#2024!Secure
    echo ⚠️  使用默认密码，请稍后修改！
)
echo.

echo ==================================================
echo 步骤 2/3: 创建数据表
echo ==================================================
echo.

mysql -u pet_admin -p%APP_PASSWORD% pet_adoption < "%SCRIPT_DIR%\02_create_tables.sql"
if %ERRORLEVEL% NEQ 0 (
    echo ❌ 创建数据表失败
    pause
    exit /b 1
)
echo ✅ 数据表创建成功
echo.

echo ==================================================
echo 步骤 3/3: 插入初始数据
echo ==================================================
echo.

set /p INSERT_DATA="是否插入测试数据（管理员账户）？(Y/N): "

if /i "%INSERT_DATA%"=="Y" (
    mysql -u pet_admin -p%APP_PASSWORD% pet_adoption < "%SCRIPT_DIR%\03_init_data.sql"
    if %ERRORLEVEL% NEQ 0 (
        echo ❌ 插入初始数据失败
        pause
        exit /b 1
    )
    echo ✅ 初始数据插入成功
) else (
    echo ℹ️  跳过初始数据插入
)
echo.

echo ==================================================
echo 🔍 验证部署
echo ==================================================
echo.

echo 📊 数据表列表:
mysql -u pet_admin -p%APP_PASSWORD% pet_adoption -e "SHOW TABLES;"
echo.

echo 📊 数据统计:
mysql -u pet_admin -p%APP_PASSWORD% pet_adoption -e "SELECT (SELECT COUNT(*) FROM users) AS '用户数', (SELECT COUNT(*) FROM organizations) AS '机构数', (SELECT COUNT(*) FROM pets) AS '宠物数';"
echo.

echo ==================================================
echo 🎉 部署完成！
echo ==================================================
echo.

echo 📝 数据库信息：
echo   数据库名: pet_adoption
echo   用户名:   pet_admin
echo   字符集:   utf8mb4
echo   端口:     3306
echo.

if /i "%INSERT_DATA%"=="Y" (
    echo ⚠️  默认管理员账户:
    echo   用户名: admin
    echo   密码:   Admin@123456
    echo.
    echo ❌ 请立即登录修改密码！
    echo.
)

echo 📋 下一步操作：
echo   1. 修改 config\config.yaml 中的数据库配置
echo   2. 修改默认密码（如果插入了测试数据）
echo   3. 配置防火墙规则
echo   4. 设置定时备份
echo.

REM 生成配置文件
set CONFIG_FILE=database_config.txt
(
echo # ============================================
echo # 数据库配置信息
echo # 生成时间: %date% %time%
echo # ============================================
echo.
echo # 应用配置 ^(config/config.yaml^)
echo database:
echo   mysql:
echo     host: "localhost"
echo     port: 3306
echo     username: "pet_admin"
echo     password: "%APP_PASSWORD%"
echo     database: "pet_adoption"
echo     charset: "utf8mb4"
echo     max_idle_conns: 10
echo     max_open_conns: 100
echo     conn_max_lifetime: 3600
echo.
echo # 连接字符串
echo DSN: pet_admin:%APP_PASSWORD%@tcp^(localhost:3306^)/pet_adoption?charset=utf8mb4^&parseTime=True^&loc=Local
echo.
echo # 管理员账户^(如果已插入^)
echo admin_username: admin
echo admin_password: Admin@123456
echo.
echo # ⚠️ 此文件包含敏感信息，请妥善保管！
) > "%CONFIG_FILE%"

echo ✅ 配置信息已保存到: %CONFIG_FILE%
echo ⚠️  此文件包含敏感信息，请妥善保管！
echo.

echo 🚀 可以启动应用了！
echo.

pause
