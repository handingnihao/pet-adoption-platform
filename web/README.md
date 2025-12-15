# 宠物领养平台前端

基于 React + TypeScript + Vite + TailwindCSS 构建的宠物领养平台前端。

## 技术栈

- **框架**: React 18 + TypeScript
- **构建工具**: Vite 6
- **样式**: TailwindCSS + shadcn/ui
- **状态管理**: Zustand
- **数据请求**: TanStack Query + Axios
- **路由**: React Router v6
- **图标**: Lucide React

## 快速开始

```bash
# 安装依赖
npm install

# 启动开发服务器
npm run dev

# 构建生产版本
npm run build

# 预览生产版本
npm run preview
```

## 项目结构

```
src/
├── components/          # 组件
│   ├── layout/         # 布局组件
│   └── ui/             # UI 基础组件
├── lib/                # 工具函数和 API
├── pages/              # 页面组件
├── router/             # 路由配置
├── store/              # 状态管理
└── App.tsx             # 应用入口
```

## 已完成页面

- [x] 首页 - 宠物展示、功能介绍
- [x] 登录页 - 用户登录

## 开发中页面

- [ ] 注册页
- [ ] 宠物列表页
- [ ] 宠物详情页
- [ ] 个人中心
- [ ] 社区动态
- [ ] 领养申请

## 环境变量

创建 `.env.development` 文件：

```
VITE_API_URL=http://localhost:8080/api/v1
```

## API 代理

开发环境下，`/api` 路径会自动代理到后端服务 `http://localhost:8080`
