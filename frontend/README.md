# 防伪码系统 - 前端

基于 Vue 3 + Element Plus 构建的防伪码系统前端项目。

## 功能特性

- 🚀 **现代化技术栈**: Vue 3 + Vite + Element Plus
- 📱 **响应式设计**: 支持PC端和移动端
- 🔐 **权限管理**: 平台管理员和商户两级权限体系
- 🎨 **美观界面**: 基于Element Plus的现代化UI设计
- 📊 **数据可视化**: 丰富的统计图表和报表
- 🔍 **防伪验证**: 独立的防伪码验证页面

## 技术栈

- **框架**: Vue 3 + Composition API
- **构建工具**: Vite
- **UI组件库**: Element Plus
- **路由管理**: Vue Router 4
- **状态管理**: Pinia
- **HTTP客户端**: Axios
- **图标库**: Element Plus Icons

## 项目结构

```
frontend/
├── src/
│   ├── components/          # 公共组件
│   │   └── Layout.vue       # 布局组件
│   ├── views/               # 页面组件
│   │   ├── Login.vue        # 登录页面
│   │   ├── Verify.vue       # 防伪验证页面
│   │   ├── platform/        # 平台管理页面
│   │   └── merchant/        # 商户管理页面
│   ├── stores/              # 状态管理
│   │   └── auth.js          # 认证状态
│   ├── services/            # API服务
│   │   └── api.js           # HTTP客户端配置
│   ├── router/              # 路由配置
│   │   └── index.js         # 路由定义
│   ├── App.vue              # 根组件
│   └── main.js              # 入口文件
├── public/                  # 静态资源
├── index.html               # HTML模板
├── vite.config.js           # Vite配置
├── package.json             # 项目配置
└── README.md               # 项目说明
```

## 快速开始

### 环境要求

- Node.js >= 16.0.0
- npm >= 7.0.0

### 安装依赖

```bash
cd frontend
npm install
```

### 开发环境运行

```bash
npm run dev
```

访问 http://localhost:3000

### 生产环境构建

```bash
npm run build
```

构建产物位于 `dist` 目录。

### 预览生产构建

```bash
npm run preview
```

## 页面说明

### 公共页面

- **登录页面** (`/login`): 平台管理员和商户登录入口
- **防伪验证页面** (`/verify`): 独立的防伪码验证入口

### 平台管理页面

- **控制台** (`/platform/dashboard`): 平台数据概览和统计
- **商户管理** (`/platform/merchants`): 商户信息管理
- **防伪码管理** (`/platform/codes`): 全局防伪码管理
- **规则设置** (`/platform/rules`): 系统规则配置
- **统计分析** (`/platform/statistics`): 数据统计和报表

### 商户管理页面

- **控制台** (`/merchant/dashboard`): 商户数据概览
- **防伪码管理** (`/merchant/codes`): 商户防伪码管理
- **生成防伪码** (`/merchant/codes/generate`): 批量生成防伪码
- **验证记录** (`/merchant/verifies`): 防伪码验证历史
- **商户信息** (`/merchant/profile`): 商户基本信息

## API配置

在 `.env.development` 和 `.env.production` 文件中配置API基础URL：

```env
VITE_API_BASE_URL=http://localhost:8080/api
```

## 权限系统

系统采用两级权限管理：

### 平台管理员
- 管理所有商户
- 查看全局统计数据
- 配置系统规则
- 管理所有防伪码

### 商户用户
- 管理自己的防伪码
- 查看自己的统计数据
- 生成防伪码批次
- 查看验证记录

## 浏览器支持

- Chrome >= 88
- Firefox >= 78
- Safari >= 14
- Edge >= 88

## 开发指南

### 添加新页面

1. 在 `src/views` 目录下创建Vue组件
2. 在 `src/router/index.js` 中添加路由配置
3. 如果需要权限控制，设置路由的 `meta` 属性

### API调用

使用 `src/services/api.js` 封装的axios实例：

```javascript
import api from '@/services/api'

// GET请求
const response = await api.get('/endpoint')

// POST请求
const response = await api.post('/endpoint', data)
```

### 状态管理

使用Pinia进行状态管理：

```javascript
import { useAuthStore } from '@/stores/auth'

const authStore = useAuthStore()
const { isAuthenticated, userInfo } = authStore
```

## 部署说明

### 构建优化

- 使用Vite进行代码分割和压缩
- 自动生成preload和prefetch链接
- CSS代码分割和压缩

### 服务器配置

确保服务器配置支持SPA路由：

```nginx
location / {
  try_files $uri $uri/ /index.html;
}
```

## 许可证

MIT License