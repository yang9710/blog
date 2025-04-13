# 博客系统前端

这是博客系统的前端项目，使用 Next.js 和 TypeScript 开发。

## 技术栈

- Next.js 14
- React 18
- TypeScript
- Tailwind CSS
- Axios

## 开发环境要求

- Node.js 18.x 或更高版本
- npm 或 yarn

## 安装依赖

```bash
npm install
# 或
yarn install
```

## 开发模式

```bash
npm run dev
# 或
yarn dev
```

开发服务器将在 [http://localhost:3000](http://localhost:3000) 启动。

## 构建生产版本

```bash
npm run build
# 或
yarn build
```

## 启动生产服务器

```bash
npm run start
# 或
yarn start
```

## 环境变量

创建 `.env.local` 文件并设置以下变量：

```
NEXT_PUBLIC_API_URL=http://localhost:8080/api/v1
```

## 项目结构

```
frontend/
├── components/     # React 组件
├── pages/         # Next.js 页面
├── styles/        # 样式文件
├── utils/         # 工具函数
└── types/         # TypeScript 类型定义
```

## 代码规范

- 使用 ESLint 进行代码检查
- 使用 Prettier 进行代码格式化
- 遵循 TypeScript 最佳实践

## 部署

项目可以部署到 Vercel 或其他支持 Next.js 的平台。

### Vercel 部署

1. 将代码推送到 GitHub 仓库
2. 在 Vercel 上导入项目
3. 配置环境变量
4. 部署

## 测试

```bash
npm run test
# 或
yarn test
```