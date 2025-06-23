# Soraka 桌面客户端

本项目基于 [Wails3](https://wails.io/) 构建，用于演示桌面应用框架的搭建。

## 目录结构

```
.
├── build/                # 打包与任务脚本
├── config/               # 配置文件目录
├── frontend/             # 前端 Vue3 代码
├── service/              # Go 后端服务
├── main.go               # 应用入口
└── Taskfile.yml          # 开发辅助任务
```

`config` 目录中包含 `clientconfig.default.json`，为默认配置。用户修改后会生成 `clientconfig.json` 保存在本地，不会被纳入版本控制。

## 已实现功能

- 启动时自动从注册表搜索 Riot Client 路径，如未找到则尝试读取《英雄联盟》客户端路径，并发送到前端显示
- 客户端路径保存到 `config/clientconfig.json`，下次启动无需再次搜索
- 托盘菜单与基础窗口控制
- 最近比赛查询接口 `/v1/lcu/recentMatches`

## 开发与调试

```bash
# 安装依赖并启动前后端热重载
wails3 dev
```

构建发布版本：

```bash
wails3 build
```

### 前端构建说明

运行 `wails3 build` 或 `npm run build` 会在 `frontend/dist` 目录生成前端文件。**注意**：构建结果会在编译 Go 程序时被嵌入。如果在编译后才执行前端构建，需要重新执行 `wails3 build` 或 `go build`，否则运行时可能报告找不到 `index.html`。

如需打包安装包，可执行：

```bash
wails3 package
```

### 常见问题排查

Windows 环境可能会在启动时输出以下错误：

```
SetProcessDpiAwarenessContext failed 0: 18446744073709551612 Access is denied.
```

这是由于系统 DPI 设置权限不足导致，通常不会影响应用正常运行，可忽略。
