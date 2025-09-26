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
# 1. 确保MySQL服务运行
# 2. 创建数据库
mysql -u root -p
CREATE DATABASE anti_fake_system;

# 3. 导入初始数据
mysql -u root -p anti_fake_system < backend/init_data.sql

# 4. 启动后端服务
cd backend
go run main.go
```

## 数据库配置
线上数据库配置:
- 主机: 8.138.244.31
- 端口: 3306
- 用户: fangwei
- 密码: NMDYyLjF52KS5S6N
- 数据库: fangwei

Redis配置:
- 主机: 8.138.244.31:6379
- 密码: 123456
- DB: 0

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