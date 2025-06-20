<template>
  <!-- 自定义标题栏 -->
  <div class="title-bar">
    <div class="title">🧊 Soraka 游戏大厅</div>
    <div class="buttons">
      <button class="window-button" @click="WindowMinimise">—</button>
      <button class="window-button" @click="WindowClose">×</button>
    </div>
  </div>

  <!-- 页面内容区域 -->
  <div class="app-content">
    <router-view />
  </div>
</template>

<script lang="ts" setup>
import { Events, Window } from "@wailsio/runtime";
import { useAppStore } from "@/store";

// pinia store
const appStore = useAppStore();

// 接收后端事件
Events.On("time", (time: any) => {
  appStore.sysTime = time.data;
});
// 接收客户端路径事件
Events.On("clientPath", (e: any) => {
  appStore.clientPath = e.data as string;
});

// 控制窗口
const WindowMinimise = () => {
  Window.Minimise();
};
const WindowClose = () => {
  Window.Hide();
};
</script>

<style lang="less">
/* 布局基础 */
html,
body,
#app {
  margin: 0;
  padding: 0;
  height: 100%;
  overflow: hidden;
  background: var(--color-menu-light-bg);
  color: var(--color-neutral-10);
  border-radius: 12px;
  box-shadow: 0 10px 40px rgba(0, 0, 0, 0.2);
}

/* 自定义标题栏 */
.title-bar {
  height: 40px;
  background: var(--color-menu-light-bg);
  display: flex;
  align-items: center;
  justify-content: space-between;
  padding: 0 12px;
  z-index: 999;
  --wails-draggable: drag;

  .title {
    font-size: 14px;
    font-weight: bold;
    user-select: none;
  }

  .buttons {
    display: flex;
    gap: 8px;

    .window-button {
      --wails-draggable: no-drag; // ✅ 阻止按钮拖动
      background: transparent;
      border: none;
      color: var(--color-neutral-10);
      width: 32px;
      height: 28px;
      font-size: 16px;
      border-radius: 4px;
      cursor: pointer;
      transition: background 0.2s;

      &:hover {
        background: var(--color-fill-2);
      }
    }
  }
}
</style>
