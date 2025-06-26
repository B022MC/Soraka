# Soraka 桌面客户端

本项目基于 [Wails3](https://wails.io/)，采用 Go 后端与 Vue3 前端，提供桌面应用开发示例。

## 目录结构

```text
.
├── build/         # 打包脚本和平台任务
├── config/        # 默认配置文件
├── const/         # 常量及资源
├── frontend/      # Vue3 前端源码
├── global/        # 全局变量与版本信息
├── internal/      # 后端业务代码
│   ├── app/       # 应用初始化
│   ├── biz/       # 业务逻辑
│   ├── conf/      # 配置加载
│   ├── dal/       # 数据访问层
│   ├── proxy/     # 代理中间件
│   ├── router/    # HTTP 路由
│   ├── server/    # 服务启动入口
│   ├── service/   # 向前端暴露的服务
│   └── wails/     # Wails 绑定代码
├── openwindow/    # 新窗口示例
├── pkg/           # 通用工具库
├── utils/         # 辅助函数
├── main.go        # 程序入口
└── Taskfile.yml   # 开发辅助任务
```

`config` 目录存放默认配置文件，运行时修改后的配置不会纳入版本控制。

## 已实现功能

- 启动时自动搜索 Riot Client 路径并保存到本地
- 托盘菜单与基础窗口控制
- 提供最近比赛等查询接口
- 通过 `/lcu/proxy/*` 路由与客户端交互

## 开发与调试

```bash
# 安装依赖并启动前后端热重载
wails3 dev
```

构建发布版本：

```bash
wails3 build
```

若使用 `go build`，需加入 `production` 构建标签以嵌入前端资源：

```bash
go build -tags production ./...
```

通过 `Taskfile` 构建：

```bash
task build PRODUCTION=true
```

### 前端构建说明

运行 `wails3 build` 或 `npm run build` 会在 `frontend/dist` 生成前端文件。**注意**：若在编译 Go 程序后执行前端构建，需要重新运行 `wails3 build` 或 `go build`，否则可能无法找到 `index.html`。

如需打包安装包，可执行：

```bash
wails3 package
```

### 常见问题排查

Windows 环境启动时可能出现以下错误：

```
SetProcessDpiAwarenessContext failed 0: 18446744073709551612 Access is denied.
```

通常是系统 DPI 权限不足，不影响程序运行，可忽略。
