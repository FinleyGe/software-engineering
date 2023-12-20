# Software Engineering Curriculum Design

这是《软件工程》课程设计项目。医院患者监护系统

## 开发说明
使用 monorepo 架构

`packages/` 中的不同的repo
- `server` 后端
- `front-end` 前端

共享代码流程：

1. fork本项目
2. `git clone your_forked_repo.git`
3. `cd path/to/repo && pnpm i`
4. 码
5. PR

### 后端

后端使用 Golang 开发，使用 Gin + Gorm
- [Go](https://go.dev/)
- [Gin](https://pkg.go.dev/github.com/gin-gonic/gin)
- [Gorm](https://gorm.io/zh_CN/docs/index.html)

### 前端

前端使用 Vue3 开发, 使用 Vite 脚手架，使用 Vue-router, pinia, Tailwind.css

- [vue3](https://vuejs.org/)
- [vue-router](https://router.vuejs.org/zh/)
- [pinia](https://pinia.vuejs.org/)
- [tailwind.css](https://tailwindcss.com/)

通过 Eslint 进行代码规范
- [eslint](https://eslint.org/)

规则集: [eslint-config-vue](https://www.npmjs.com/package/@finley_ge/eslint-config-vue)

### 架构 & 其他说明

1. 使用 pnpm monorepo 管理项目
2. 注意git commit 规范格式(husky + commitlint进行强制规范)
    example:
    ```
    docs(README): Edit something
    fix(front-end): Fix some bug
    ```
3. 写完你的功能提 PR
