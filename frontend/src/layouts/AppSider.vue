<template>
  <a-layout id="app-layout-sider">
    <a-layout-sider
        v-model:collapsed="collapsed"
        :width="220"
        :collapsed-width="64"
        collapsible
        theme="light"
        class="layout-sider"
        :class="{ 'collapsed-sider': collapsed }"
    >
      <div class="sider-menu">
        <!-- 顶部 logo 与标题 -->
        <div class="sider-header">
          <a-avatar :size="32" class="logo-avatar">
            <img src="@/assets/logo.png" />
          </a-avatar>
          <span class="logo-title" :class="{ hidden: collapsed }">Soraka</span>
        </div>

        <!-- 主菜单 -->
        <a-menu
            class="menu-item"
            theme="light"
            mode="inline"
            :selectedKeys="[current]"
            @menu-item-click="menuHandle"
        >
          <template v-for="menuInfo in menulist" :key="menuInfo.name">
            <a-menu-item v-if="!menuInfo.meta?.hideInMenu" :key="menuInfo.name">
              <div class="menu-item-inner">
                <icon-font :type="menuInfo.meta.icon" class="menu-icon" />
                <span class="menu-text" :class="{ hidden: collapsed }">{{ menuInfo.meta.title }}</span>
              </div>
            </a-menu-item>
          </template>
        </a-menu>

        <a-menu
            class="footer-tools"
            theme="light"
            mode="inline"
            :selectedKeys="current === 'setting' ? ['setting'] : []"
        >
          <a-menu-item key="opgg" @click="handleTool('opgg')">
            <div class="menu-item-inner">
              <img src="@/assets/images/opgg.svg" class="menu-icon" />
              <span class="menu-text" :class="{ hidden: collapsed }">OP.GG</span>
            </div>
          </a-menu-item>
          <a-menu-item key="fix" @click="handleTool('fix')">
            <div class="menu-item-inner">
              <icon-font type="icon-ArrowCircle" class="menu-icon" />
              <span class="menu-text" :class="{ hidden: collapsed }">修复无限加载</span>
            </div>
          </a-menu-item>
          <a-menu-item key="notice" @click="handleTool('notice')">
            <div class="menu-item-inner">
              <icon-font type="icon-Alert" class="menu-icon" />
              <span class="menu-text" :class="{ hidden: collapsed }">公告</span>
            </div>
          </a-menu-item>
          <a-menu-item key="setting" @click="handleSetting">
            <div class="menu-item-inner">
              <icon-font type="icon-Setting" class="menu-icon" />
              <span class="menu-text" :class="{ hidden: collapsed }">设置</span>
            </div>
          </a-menu-item>
        </a-menu>

        <!-- 用户信息 -->
        <div class="footer" :class="{ factive: current === 'setting' }" @click="handleSetting">
          <a-avatar :size="32" :src="userStore.avatar || '@/assets/logo.png'" />
          <div class="footer-text" :class="{ hidden: collapsed }">
            <div class="footer-name">{{ userStore.nickname || '未登录' }}</div>
            <div class="footer-rank">{{ userStore.region }}</div>
          </div>
        </div>
      </div>
    </a-layout-sider>

    <a-layout>
      <a-layout-content class="layout-content">
        <router-view />
      </a-layout-content>
    </a-layout>
  </a-layout>
</template>

<script lang="ts" setup>
import { ref, computed, watch, onMounted } from 'vue'
import { useRouter, useRoute } from 'vue-router'
import routerMap from '@/router/routerMap'
import { Events } from '@wailsio/runtime'
import { LcuService } from '/#/Soraka/service'
import { useUserStore } from '@/store'

const router = useRouter()
const route = useRoute()
const current = ref(route.name)
const collapsed = ref(false)
const userStore = useUserStore()

const menulist = computed(() => routerMap.find(item => item.path === '/')?.children || [])

watch(() => route.name, name => (current.value = name))

const menuHandle = key => {
  current.value = key
  const target = menulist.value.find(i => i.name === key)
  if (target) {
    router.push({ name: target.name, params: target.params })
  }
}

const handleSetting = () => {
  current.value = 'setting'
  router.push({ name: 'setting' })
}

const handleTool = type => {
  console.log('工具点击', type)
}

const loadUserInfo = () => {
  LcuService.GetCurrentUserInfo().then(info => {
    if (!info) return
    userStore.setInfo({
      nickname: info.displayName,
      avatar: `https://127.0.0.1:${info.port}/lol-game-data/assets/v1/profile-icons/${info.profileIconId}.jpg`,
      region: info.region
    })
  }).catch(err => {
    console.error('[GetCurrentUserInfo]', err)
  })
}

onMounted(() => {
  LcuService.CheckLogin().then(([ok]) => {
    if (ok) {
      loadUserInfo()
    }
  })
  Events.On('lcuStatus', (d: any) => {
    const status = Array.isArray(d.data) ? d.data[0] : d.data
    if (status) {
      loadUserInfo()
    }
  })
})
</script>

<style scoped lang="less">
#app-layout-sider {
  height: 100vh;

  .layout-sider {
    border-right: 1px solid var(--color-border);
    background: var(--color-menu-light-bg);
    transition: all 0.3s ease;
    will-change: width;
    z-index: 9;
    display: flex;
    flex-direction: column;
  }

  .sider-menu {
    height: 100%;
    display: flex;
    flex-direction: column;
    justify-content: space-between;
    width: 100%;
    overflow: hidden;

    .sider-header {
      height: 60px;
      display: flex;
      align-items: center;
      padding: 0 16px;
      background: var(--color-menu-light-bg);

      .logo-avatar {
        flex-shrink: 0;
      }

      .logo-title {
        font-size: 14px;
        font-weight: bold;
        margin-left: 12px;
        color: var(--color-text);
        transition: opacity 0.3s, width 0.3s;
        white-space: nowrap;
      }

      .hidden {
        opacity: 0;
        width: 0;
        overflow: hidden;
      }
    }

    .menu-item {
      flex: 1;
      overflow-y: auto;

      scrollbar-width: none;
      -ms-overflow-style: none;
      &::-webkit-scrollbar {
        display: none;
      }

      :deep(.arco-menu-item) {
        padding: 0;
        height: 48px;
        display: flex;
        align-items: center;
        color: var(--color-text);

        .menu-item-inner {
          display: flex;
          align-items: center;
          width: 100%;
          height: 100%;
          padding: 0 12px;

          .menu-icon {
            font-size: 18px;
            width: 24px;
            height: 48px;
            line-height: 48px;
            text-align: center;
            flex-shrink: 0;
          }

          .menu-text {
            font-size: 14px;
            margin-left: 12px;
            transition: opacity 0.3s, width 0.3s;
            white-space: nowrap;
            color: var(--color-text);
          }

          .hidden {
            opacity: 0;
            width: 0;
            overflow: hidden;
          }
        }

        &:hover {
          background-color: var(--color-fill-2);
        }

        &.arco-menu-selected {
          background-color: var(--color-fill-2);
          color: var(--color-primary);
        }
      }
    }

    .footer-tools {
      overflow-y: auto;

      scrollbar-width: none;
      -ms-overflow-style: none;
      &::-webkit-scrollbar {
        display: none;
      }

      :deep(.arco-menu-item) {
        padding: 0;
        height: 48px;
        display: flex;
        align-items: center;
        color: var(--color-text);

        .menu-item-inner {
          display: flex;
          align-items: center;
          width: 100%;
          height: 100%;
          padding: 0 12px;

          .menu-icon {
            font-size: 18px;
            width: 24px;
            height: 48px;
            line-height: 48px;
            text-align: center;
            flex-shrink: 0;
          }

          .menu-text {
            font-size: 14px;
            margin-left: 12px;
            transition: opacity 0.3s, width 0.3s;
            white-space: nowrap;
            color: var(--color-text);
          }

          .hidden {
            opacity: 0;
            width: 0;
            overflow: hidden;
          }
        }

        &:hover {
          background-color: var(--color-fill-2);
        }

        &.factive {
          background-color: var(--color-fill-2);
        }
      }
    }

    .footer {
      display: flex;
      align-items: center;
      height: 60px;
      padding: 0 16px;
      border-top: 1px solid var(--color-border);
      cursor: pointer;
      transition: background 0.3s;

      &:hover {
        background-color: var(--color-fill-2);
      }

      .arco-avatar {
        flex-shrink: 0;
        width: 32px;
        height: 32px;
        border-radius: 50%;
        object-fit: cover;
      }

      .footer-text {
        margin-left: 12px;
        transition: opacity 0.3s, width 0.3s;
        white-space: nowrap;

        .footer-name {
          font-weight: 500;
          font-size: 13px;
          color: var(--color-text);
        }

        .footer-rank {
          font-size: 12px;
          color: var(--color-text-2);
        }

        &.hidden {
          opacity: 0;
          width: 0;
          overflow: hidden;
        }
      }

      &.factive {
        background-color: var(--color-fill-2);
      }
    }
  }

  .layout-content {
    background: var(--color-bg-1);
    color: var(--color-text);
    padding: 16px;
  }

  .collapsed-sider {
    :deep(.arco-menu) {
      width: 100%;
    }

    :deep(.arco-menu-inner) {
      display: flex !important;
      flex-direction: column;
      align-items: center;
      justify-content: flex-start;
      padding: 4px 4px !important;
      width: 100%;
      box-sizing: border-box;
    }

    :deep(.arco-menu-item) {
      width: 100%;
      justify-content: center !important;
      padding: 0 !important;
      overflow: hidden;

      .arco-menu-item-inner {
        display: flex !important;
        justify-content: center !important;
        align-items: center !important;
        padding: 0 !important;
        width: 100%;
        height: 100%;
      }
    }

    .menu-item-inner {
      display: flex;
      justify-content: center;
      align-items: center;
      width: 100%;
      height: 100%;
      padding: 0;
    }

    .menu-icon {
      margin: 0 auto;
      width: 24px;
      height: 24px;
      display: flex;
      align-items: center;
      justify-content: center;
    }

    .menu-text,
    .footer-text {
      display: none !important;
    }

    .footer {
      justify-content: center;
      padding: 0;
    }
  }
}



</style>
