📖 项目定位
本项目是“返乡青年助力乡村数字素养提升”倡议的官方数字化基础设施。针对乡村教育中数字资源匮乏、AI工具应用门槛高、支教活动短期化等痛点，提供一套可持续、可复制、易上手的在线教育支持平台。

它不仅是一个资源库，更是连接返乡大学生、乡村教师与乡村学生的数字桥梁。

🛠 功能模块
基于 src/views 下的工程架构，平台划分为以下三大核心专区：

1. 资源库 (Resource Hub)
AI 课程创作 (CourseCreat.vue)：提供标准化的 AI 启蒙教案与多媒体课件。

数字教学 PPT (PPT.vue)：针对不同学段设计的数字素养提升课件库。

2. 教师支持专区 (Teacher Support)
AI 备课指南 (AIPrepare.vue)：手把手教乡村教师如何利用 AI 提高备课效率，实现减负增效。

AI 批改示例 (AIRevise.vue)：展示 AI 在作业批改、作文评价中的应用场景。

3. 互动学习区 (Interact Zone)
作品展示墙 (WorkShow.vue)：学生 AI 创作成果（如 AI 画作、AI 故事）的集中展示空间。

每日任务 (EverydayWork.vue)：轻量化的数字素养训练营，通过打卡任务保持学习热度。

🚀 技术架构
项目采用现代前端主流技术栈，确保了系统的响应速度与维护性：

框架：Vue 3 (Composition API + <script setup>)

构建工具：Vite 7 (极速冷启动与热更新)

UI 组件库：Ant Design Vue 4.2 (专业级 UI 交互)

路由管理：Vue Router 4 (实现 SPA 单页应用导航)

开发语言：TypeScript (增强代码健壮性与类型安全)

📁 目录结构说明
Plaintext
frontend/
├── src/
│   ├── components/       # 通用组件 (如自定义 Wrapper 布局)
│   ├── router/           # 路由配置 (控制页面跳转逻辑)
│   ├── views/            # 业务页面
│   │   ├── Resource/     # 课件与创作模块
│   │   ├── Teacher/      # 教师赋能工具
│   │   └── Interact/     # 师生互动模块
│   ├── App.vue           # 根组件 (采用 Home 布局导航)
│   └── main.ts           # 应用入口
├── package.json          # 项目依赖与运行指令
└── tsconfig.json         # TypeScript 配置
💻 本地开发
环境准备
Node.js (建议 v20+)

npm 或 pnpm

安装启动
克隆仓库

Bash
git clone [repository-url]
安装依赖

Bash
npm install
启动开发服务器

Bash
npm run dev
项目构建

Bash
npm run build
📅 项目规划与愿景
Phase 1: 完成基础功能开发，实现在甘肃等试点地区的资源落地。

Phase 2: 接入轻量化 AI 接口，提供在线 AI 辅助对话与创作功能。

Phase 3: 完善数据反馈模块，通过学情数据分析优化乡村教育方案。

📄 开源协议
本项目遵循 MIT License。 Copyright (c) 2026 LiuYuheng.