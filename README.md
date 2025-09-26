# 防伪系统 (Anti-Fake System)

## 项目概述
这是一个完整的防伪码管理系统，包含Vue 3前端和Go后端。

## 技术栈
### 前端
- Vue 3 + Vite
- Element Plus UI组件库
- ECharts数据可视化
- Axios HTTP客户端

### 后端
- Go + Gin Web框架
- GORM数据库ORM
- MySQL数据库
- Redis缓存
- JWT认证

## 运行状态

### 前端 ✅
- 已成功启动并运行在 http://localhost:5173
- 可以正常访问和预览

### 后端 ⚠️
- 代码编译成功，无语法错误
- 需要MySQL数据库服务才能启动

## 默认账户信息

### 平台管理员
- **用户名**: `admin`
- **密码**: `password`
- **邮箱**: `admin@antifake.com`
- **权限**: 平台管理员，可管理所有商户

### 商户管理员
- **用户名**: `merchant`
- **密码**: `password`
- **邮箱**: `merchant@test.com`
- **权限**: 商户管理员，可管理自己的商品和防伪码

## 启动步骤

### 1. 前端启动 (已完成)
```bash
cd frontend
npm install
npm run dev
```
访问: http://localhost:5173

### 2. 后端启动 (需要数据库)
```bash
# 1. 配置环境变量
cd backend
cp .env.example .env
# 编辑 .env 文件，设置你的数据库连接信息

# 2. 确保MySQL服务运行
# 3. 创建数据库
mysql -u root -p
CREATE DATABASE anti_fake_system;

# 4. 导入初始数据
mysql -u root -p anti_fake_system < init_data.sql

# 5. 安装Go依赖并启动后端服务
go mod tidy
go run main.go
```

## 环境配置

### 环境变量设置

项目使用环境变量来管理敏感配置信息，确保安全性。

#### 根目录环境变量
1. 复制环境变量模板文件：
```bash
cp .env.example .env
```

2. 编辑根目录的 `.env` 文件，配置你的数据库和Redis连接信息。

#### 后端环境变量
1. 复制后端环境变量模板文件：
```bash
cd backend
cp .env.example .env
```

2. 编辑后端的 `.env` 文件，配置具体的连接参数：
```bash
# 数据库配置
DB_HOST=your_database_host
DB_PORT=3306
DB_USER=your_database_user
DB_PASSWORD=your_database_password
DB_NAME=your_database_name

# Redis配置
REDIS_ADDR=your_redis_host:6379
REDIS_PASSWORD=your_redis_password
REDIS_DB=0

# JWT配置
JWT_SECRET=your_strong_jwt_secret_key
JWT_EXPIRE=24
```

**安全提示**: 
- `.env` 文件包含敏感信息，已被添加到 `.gitignore` 中，不会被提交到版本控制系统
- 请使用强密码和复杂的JWT密钥
- 生产环境中请更改所有默认密码

## API接口
- 健康检查: GET /health
- 平台登录: POST /api/auth/platform/login
- 商户登录: POST /api/auth/merchant/login
- 防伪码验证: POST /api/public/verify

## 功能特性
- 🔐 JWT认证系统
- 👥 多角色权限管理 (平台管理员/商户管理员)
- 🏷️ 防伪码生成和验证
- 📊 数据统计和报表
- 🛡️ 安全中间件和CORS支持
- 📱 响应式前端界面

## 项目结构
```
├── frontend/          # Vue 3前端
│   ├── src/
│   │   ├── components/
│   │   ├── views/
│   │   ├── router/
│   │   └── api/
├── backend/           # Go后端
│   ├── controllers/   # 控制器
│   ├── handlers/      # 处理器
│   ├── models/        # 数据模型
│   ├── middleware/    # 中间件
│   ├── services/      # 业务服务
│   ├── database/      # 数据库配置
│   └── config/        # 配置管理