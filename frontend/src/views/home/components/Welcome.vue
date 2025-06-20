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
              sysTime ? sysTime[0] : "---"
            }}</span>
          </p>
          <p>
            客户端路径：{{ clientPath || "未找到" }}
            <span class="status">
              <a-tag v-if="lcuOnline" color="green">已连接LCU</a-tag>
              <span v-else class="waiting">等待连接</span>
            </span>
          </p>
          <p v-if="lcuPort">端口：{{ lcuPort }} Token：{{ lcuToken }}</p>
        </div>
      </a-space>
    </a-row>
  </a-card>
</template>

<script lang="ts" setup>
import { ref, onMounted } from "vue";
import { Events } from "@wailsio/runtime";
import type { AuthInfo } from "@/api/lcu";
import { WailsAPI } from "/#/Soraka/service/lcu";
import { useUserStore } from "@/store";
import { goodTimeText } from "@/utils";

const userStore = useUserStore();

const sysTime = ref("");
const clientPath = ref("");
const lcuOnline = ref(false);
const lcuPort = ref("");
const lcuToken = ref("");
const authInfo = ref<AuthInfo | null>(null);

onMounted(() => {
  // 初始化获取 LCU 凭证
  WailsAPI.GetAuthInfo().then((info: AuthInfo) => {
    authInfo.value = info
    lcuOnline.value = true
    lcuPort.value = String(info.port)
    lcuToken.value = info.token
  }).catch(() => {
    lcuOnline.value = false
  })

  // 获取当前召唤师信息
  WailsAPI.GetCurrentSummoner().then((info: any) => {
    if (info && authInfo.value) {
      userStore.setInfo({
        nickname: info.displayName,
        avatar: `https://127.0.0.1:${authInfo.value.port}/lol-game-data/assets/v1/profile-icons/${info.profileIconId}.jpg`,
        region: info.region
      })
    }
  })

  // 监听系统时间事件
  Events.On("time", (time: any) => {
    // console.log("[Event] time 收到:", time);
    sysTime.value = time.data as string;
  });

  // 监听 clientPath 事件
  Events.On("clientPath", (p: any) => {
    console.log("[Event] clientPath 收到:", p);
    clientPath.value = p.data as string;
  });

  // 监听 lcuStatus 事件
  Events.On("lcuStatus", (d: any) => {
    console.log("[Event] lcuStatus 收到:", d);
    const status = Array.isArray(d.data) ? d.data[0] : d.data;
    if (typeof status === 'boolean') {
      lcuOnline.value = status;
    } else {
      console.warn("[Event] lcuStatus 数据格式异常:", d.data);
      lcuOnline.value = false;  // 或保留上次值
    }
    console.log("[状态] 更新 lcuOnline:", lcuOnline.value);
  });


  // 监听 lcuCreds 事件
  Events.On("lcuCreds", (d: any) => {
    console.log("[Event] lcuCreds 收到:", d);
    const creds = Array.isArray(d.data) ? d.data[0] : d.data;
    if (creds && typeof creds === 'object') {
      lcuPort.value = creds.port ?? "";
      lcuToken.value = creds.token ?? "";
      console.log("[状态] 更新 lcuPort:", lcuPort.value, "lcuToken:", lcuToken.value);
    } else {
      console.warn("[Event] lcuCreds 数据格式异常:", d.data);
      lcuPort.value = "";
      lcuToken.value = "";
    }
  });


});
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
