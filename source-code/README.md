
## 开发指南

启动开发服务器：

```bash
npm run dev
# 或
yarn dev
# 或
pnpm dev
```

在浏览器打开 `http://localhost:3000` 查看效果。

你可以通过修改 `app/page.tsx` 开始编辑页面，保存后会自动热更新。

### 构建与启动

构建生产版本：

```bash
npm run build
# 或
yarn build
# 或
pnpm build
```

启动生产服务器：

```bash
npm run start
# 或
yarn start
# 或
pnpm start
```

### 命令说明

| 命令 | 用途 | 适用场景 | 说明 |
|------|------|----------|------|
| `dev` | 启动开发服务器 | 开发阶段 | 支持热更新 (Hot Reloading)，修改代码后立即生效。不要在生产环境使用。 |
| `build` | 构建生产版本 | 部署前 | 编译、压缩代码，生成静态页面和优化资源。生成的产物位于 `.next` 目录。 |
| `start` | 启动生产服务器 | 生产环境 | 运行经过 `build` 优化后的代码。**必须先执行 `build` 才能运行此命令**。 |

> **提示**：通常的部署流程是先执行 `npm run build`，然后执行 `npm run start`。

## 了解更多

- Next.js 文档：https://nextjs.org/docs
- 学习 Next.js：https://nextjs.org/learn
- Next.js 仓库：https://github.com/vercel/next.js

## 部署

推荐使用 [Vercel 平台](https://vercel.com/new?utm_medium=default-template&filter=next.js&utm_source=create-next-app&utm_campaign=create-next-app-readme) 部署。

更多部署细节参考官方文档：https://nextjs.org/docs/app/building-your-application/deploying
