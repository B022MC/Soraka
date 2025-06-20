<template>
  <a-card class="card" :bordered="false">
    <a-row
      wrap
      :gutter="[{ xs: 0, sm: 14, md: 14, lg: 14, xl: 14, xxl: 14 }, 16]"
      class="content"
    >
      <a-space size="medium">
        <a-avatar style="padding: 2px" :size="68">
          <img src="@/assets/logo.png" />
        </a-avatar>
        <div class="welcome">
          <p class="hello">{{ goodTimeText() }}！{{ userStore.nickname }}</p>
          <p>
            当前时间：<span class="datetime">{{
              appStore.sysTime ? appStore.sysTime[0] : "---"
            }}</span>
          </p>
          <p>
            客户端路径：{{ appStore.clientPath || "未找到" }}
            <span class="status">
              <a-tag v-if="appStore.lcuOnline" color="green">已连接LCU</a-tag>
              <span v-else class="waiting">等待连接</span>
            </span>
          </p>
          <p v-if="appStore.lcuPort">
            端口：{{ appStore.lcuPort }} Token：{{ appStore.lcuToken }}
          </p>
        </div>
      </a-space>
    </a-row>
  </a-card>
</template>

<script lang="ts" setup>
// import { ref } from 'vue';
import { useUserStore, useAppStore } from "@/store";
import { goodTimeText } from "@/utils";
const appStore = useAppStore();
// import {Events} from "@wailsio/runtime";
// const innerText=ref("")
// Events.On('time', (time) => {
//     innerText.value = time.data;
// });
const userStore = useUserStore();
</script>

<style scoped lang="less">
:deep(.arco-statistic-title) {
  margin-bottom: 0;
}

.card {
  background: transparent;
  .content {
    padding: 8px 20px;
    .welcome {
      margin: 8px 0;
      color: var(--color-neutral-6);
      .hello {
        font-size: 1.25rem;
        color: var(--color-neutral-8);
        margin-bottom: 10px;
      }
      .status {
        margin-left: 8px;
      }
      .waiting {
        display: inline-block;
        position: relative;
        padding-right: 1.4em;
      }
      .waiting::after {
        content: "";
        position: absolute;
        top: 50%;
        right: 0;
        width: 0.8em;
        height: 0.8em;
        margin-top: -0.4em;
        border: 2px solid var(--color-neutral-6);
        border-top-color: transparent;
        border-radius: 50%;
        animation: spin 1s linear infinite;
      }
      @keyframes spin {
        to {
          transform: rotate(360deg);
        }
      }
    }
  }
}
</style>
