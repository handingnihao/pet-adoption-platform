# Docker 部署指南

## 快速开始

### 1. 环境要求

- Docker 20.10+
- Docker Compose 2.0+

### 2. 配置文件

在启动前，请确保 `config/config.yaml` 中的数据库配置与 docker-compose.yml 一致：

```yaml
database:
  mysql:
    host: mysql  # Docker 服务名
    port: 3306
    username: pet_admin
    password: pet_admin_password
    database: pet_adoption
  redis:
    host: redis  # Docker 服务名
    port: 6379
    password: redis_password
    db: 0
```

### 3. 启动服务

```bash
# 在项目根目录执行
cd docker

# 构建并启动所有服务
docker-compose up -d

# 查看服务状态
docker-compose ps

# 查看日志
docker-compose logs -f

# 查看特定服务日志
docker-compose logs -f backend
docker-compose logs -f frontend
```

### 4. 停止服务

```bash
# 停止所有服务
docker-compose down

# 停止并删除数据卷（慎用！会删除数据库数据）
docker-compose down -v
```

## 服务说明

### MySQL (端口 3306)
- 容器名: `pet-adoption-mysql`
- 数据库: `pet_adoption`
- 用户名: `pet_admin`
- 密码: `pet_admin_password`
- Root密码: `123456`

### Redis (端口 6379)
- 容器名: `pet-adoption-redis`
- 密码: `redis_password`

### 后端服务 (端口 8080)
- 容器名: `pet-adoption-backend`
- API地址: http://localhost:8080/api/v1

### 前端服务 (端口 80)
- 容器名: `pet-adoption-frontend`
- 访问地址: http://localhost

## 常用命令

### 重启服务
```bash
docker-compose restart backend
docker-compose restart frontend
```

### 进入容器
```bash
# 进入后端容器
docker exec -it pet-adoption-backend sh

# 进入MySQL容器
docker exec -it pet-adoption-mysql mysql -upet_admin -ppet_admin_password pet_adoption

# 进入Redis容器
docker exec -it pet-adoption-redis redis-cli -a redis_password
```

### 查看资源使用
```bash
docker stats
```

### 清理未使用的资源
```bash
# 清理未使用的镜像
docker image prune -a

# 清理未使用的容器
docker container prune

# 清理未使用的卷
docker volume prune
```

## 数据持久化

数据卷：
- `mysql_data`: MySQL 数据
- `redis_data`: Redis 数据

数据卷位置查看：
```bash
docker volume inspect docker_mysql_data
docker volume inspect docker_redis_data
```

## 健康检查

所有服务都配置了健康检查：

```bash
# 查看健康状态
docker-compose ps
```

健康状态说明：
- `healthy`: 服务正常
- `unhealthy`: 服务异常
- `starting`: 服务启动中

## 故障排查

### 1. 服务启动失败

```bash
# 查看详细日志
docker-compose logs backend

# 检查配置文件
docker exec -it pet-adoption-backend cat /root/config/config.yaml
```

### 2. 数据库连接失败

```bash
# 检查 MySQL 是否健康
docker-compose ps mysql

# 测试数据库连接
docker exec -it pet-adoption-mysql mysql -upet_admin -ppet_admin_password -e "SELECT 1"
```

### 3. 前端无法访问后端

- 检查 Nginx 配置中的 `proxy_pass` 是否指向 `backend:8080`
- 确保前后端在同一 Docker 网络中

### 4. 端口冲突

如果本地已有服务占用端口，修改 `docker-compose.yml` 中的端口映射：

```yaml
ports:
  - "8081:8080"  # 将宿主机端口改为 8081
```

## 生产环境建议

1. **修改默认密码**: 修改 MySQL、Redis 的默认密码
2. **使用环境变量**: 敏感信息使用 `.env` 文件管理
3. **配置 HTTPS**: 使用 Let's Encrypt 配置 SSL 证书
4. **限制资源**: 为容器设置 CPU 和内存限制
5. **日志管理**: 配置日志轮转，避免日志文件过大
6. **备份策略**: 定期备份 MySQL 数据卷

## 更新部署

```bash
# 拉取最新代码
git pull

# 重新构建并启动
docker-compose up -d --build

# 查看更新状态
docker-compose ps
```
