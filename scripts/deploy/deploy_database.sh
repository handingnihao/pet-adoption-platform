#!/bin/bash
# ============================================
# 爱心宠物领养平台 - 数据库一键部署脚本
# MySQL 8.4
# ============================================

set -e  # 遇到错误立即退出

# 颜色定义
RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
BLUE='\033[0;34m'
NC='\033[0m' # No Color

# 打印函数
print_info() {
    echo -e "${BLUE}ℹ️  $1${NC}"
}

print_success() {
    echo -e "${GREEN}✅ $1${NC}"
}

print_warning() {
    echo -e "${YELLOW}⚠️  $1${NC}"
}

print_error() {
    echo -e "${RED}❌ $1${NC}"
}

print_header() {
    echo ""
    echo "=================================================="
    echo -e "${GREEN}$1${NC}"
    echo "=================================================="
    echo ""
}

# 检查命令是否存在
check_command() {
    if ! command -v $1 &> /dev/null; then
        print_error "$1 未安装，请先安装"
        exit 1
    fi
}

# ============================================
# 主流程
# ============================================

print_header "🐾 爱心宠物领养平台 - 数据库部署工具"

# 检查 MySQL 是否安装
print_info "检查 MySQL 环境..."
check_command mysql

# 获取 MySQL 版本
MYSQL_VERSION=$(mysql --version | grep -oP '\d+\.\d+' | head -1)
print_success "MySQL 版本: $MYSQL_VERSION"

# 检查脚本文件
SCRIPT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")/../migration" && pwd)"
print_info "脚本目录: $SCRIPT_DIR"

if [ ! -f "$SCRIPT_DIR/01_create_user_and_database.sql" ]; then
    print_error "找不到SQL脚本文件，请检查目录结构"
    exit 1
fi

# 获取 root 密码
print_header "📝 请输入 MySQL Root 密码"
read -sp "Root密码: " ROOT_PASSWORD
echo ""

# 测试 root 连接
print_info "测试 MySQL 连接..."
if ! mysql -u root -p"$ROOT_PASSWORD" -e "SELECT 1;" &> /dev/null; then
    print_error "无法连接到 MySQL，请检查密码"
    exit 1
fi
print_success "MySQL 连接成功"

# 确认部署
print_header "⚠️  即将开始部署"
echo "这将执行以下操作："
echo "  1. 创建数据库 pet_adoption"
echo "  2. 创建用户 pet_admin"
echo "  3. 创建 10 个数据表"
echo "  4. 插入初始数据（包括管理员账户）"
echo ""
read -p "是否继续？(y/n): " CONFIRM

if [ "$CONFIRM" != "y" ] && [ "$CONFIRM" != "Y" ]; then
    print_warning "部署已取消"
    exit 0
fi

# ============================================
# 步骤1: 创建用户和数据库
# ============================================
print_header "步骤 1/3: 创建用户和数据库"

if mysql -u root -p"$ROOT_PASSWORD" < "$SCRIPT_DIR/01_create_user_and_database.sql"; then
    print_success "数据库和用户创建成功"
else
    print_error "创建数据库失败"
    exit 1
fi

# 获取应用用户密码
print_info "请输入 pet_admin 用户密码（默认: Pet@Admin#2024!Secure）"
read -sp "密码: " APP_PASSWORD
echo ""
if [ -z "$APP_PASSWORD" ]; then
    APP_PASSWORD="Pet@Admin#2024!Secure"
    print_warning "使用默认密码，请稍后修改！"
fi

# ============================================
# 步骤2: 创建数据表
# ============================================
print_header "步骤 2/3: 创建数据表"

if mysql -u pet_admin -p"$APP_PASSWORD" pet_adoption < "$SCRIPT_DIR/02_create_tables.sql"; then
    print_success "数据表创建成功"
else
    print_error "创建数据表失败"
    exit 1
fi

# ============================================
# 步骤3: 插入初始数据
# ============================================
print_header "步骤 3/3: 插入初始数据"

read -p "是否插入测试数据（管理员账户）？(y/n): " INSERT_DATA

if [ "$INSERT_DATA" == "y" ] || [ "$INSERT_DATA" == "Y" ]; then
    if mysql -u pet_admin -p"$APP_PASSWORD" pet_adoption < "$SCRIPT_DIR/03_init_data.sql"; then
        print_success "初始数据插入成功"
    else
        print_error "插入初始数据失败"
        exit 1
    fi
else
    print_info "跳过初始数据插入"
fi

# ============================================
# 验证部署
# ============================================
print_header "🔍 验证部署"

# 查看表列表
print_info "数据表列表:"
mysql -u pet_admin -p"$APP_PASSWORD" pet_adoption -e "SHOW TABLES;"

# 统计信息
print_info "数据统计:"
mysql -u pet_admin -p"$APP_PASSWORD" pet_adoption -e "
SELECT 
    (SELECT COUNT(*) FROM users) AS '用户数',
    (SELECT COUNT(*) FROM organizations) AS '机构数',
    (SELECT COUNT(*) FROM pets) AS '宠物数';
"

# ============================================
# 部署完成
# ============================================
print_header "🎉 部署完成！"

echo "数据库信息："
echo "  数据库名: pet_adoption"
echo "  用户名:   pet_admin"
echo "  字符集:   utf8mb4"
echo "  端口:     3306"
echo ""

if [ "$INSERT_DATA" == "y" ] || [ "$INSERT_DATA" == "Y" ]; then
    print_warning "默认管理员账户:"
    echo "  用户名: admin"
    echo "  密码:   Admin@123456"
    echo ""
    print_error "⚠️  请立即登录修改密码！"
    echo ""
fi

print_info "下一步操作："
echo "  1. 修改 config/config.yaml 中的数据库配置"
echo "  2. 修改默认密码（如果插入了测试数据）"
echo "  3. 配置防火墙规则"
echo "  4. 设置定时备份"
echo ""

print_success "可以启动应用了！"

# ============================================
# 生成配置示例
# ============================================
CONFIG_FILE="database_config.txt"
cat > "$CONFIG_FILE" << EOF
# ============================================
# 数据库配置信息
# 生成时间: $(date)
# ============================================

# 应用配置 (config/config.yaml)
database:
  mysql:
    host: "localhost"
    port: 3306
    username: "pet_admin"
    password: "$APP_PASSWORD"
    database: "pet_adoption"
    charset: "utf8mb4"
    max_idle_conns: 10
    max_open_conns: 100
    conn_max_lifetime: 3600

# 连接字符串
DSN: pet_admin:$APP_PASSWORD@tcp(localhost:3306)/pet_adoption?charset=utf8mb4&parseTime=True&loc=Local

# 管理员账户（如果已插入）
admin_username: admin
admin_password: Admin@123456

# ⚠️ 此文件包含敏感信息，请妥善保管！
EOF

print_success "配置信息已保存到: $CONFIG_FILE"
print_warning "此文件包含敏感信息，请妥善保管！"
