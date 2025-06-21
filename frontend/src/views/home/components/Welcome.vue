<template>
  <a-card class="card" :bordered="false">
    <a-row
      wrap
      :gutter="[{ xs: 0, sm: 14, md: 14, lg: 14, xl: 14, xxl: 14 }, 16]"
      class="content"
    >
      <a-space size="medium">
        <a-avatar style="padding: 2px" :size="68">
          <img
              :src="userStore.avatar || defaultAvatar"
              @error="(e) => e.target.src = defaultAvatar"
          />        </a-avatar>
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
import { WailsAPI } from "/#/Soraka/service/lcu";
import { useUserStore } from "@/store";
import { goodTimeText } from "@/utils";

const userStore = useUserStore();

const sysTime = ref("");
const clientPath = ref("");
const lcuOnline = ref(false);
const lcuPort = ref("");
const lcuToken = ref("");

// const loadUserInfo = () => {
//   WailsAPI.GetCurrentSummoner()
//     .then((info: any) => {
//       if (!info) return;
//       userStore.setInfo({
//         nickname: info.displayName,
//         avatar: info.avatar,
//         region: info.region,
//         tag: info.tag,
//         rank: info.rank,
//         winRate: info.winRate,
//         wins: info.wins,
//         losses: info.losses,
//         totalGames: info.totalGames,
//         createtime: info.createtime,
//         level: info.level,
//         xpSinceLastLevel: info.xpSinceLastLevel,
//         xpUntilNextLevel: info.xpUntilNextLevel,
//       });
//     })
//     .catch((err) => {
//       console.error("[GetCurrentSummoner]", err);
//     });
// };
const defaultAvatar = new URL('@/assets/logo.png', import.meta.url).href;

onMounted(() => {
  // 监听系统时间事件
  Events.On("time", (time: any) => {
    sysTime.value = time.data as string;
  });
  // loadUserInfo
  // 获取客户端路径
  WailsAPI.GetClientPath()
    .then((p: string) => {
      clientPath.value = p;
    })
    .catch(() => {
      clientPath.value = "";
    });

  // 监听 lcuStatus 事件
  Events.On("lcuStatus", (d: any) => {
    const payload = Array.isArray(d.data) ? d.data[0] : d.data;
    console.log("[Event] lcuStatus 收到:", payload);
    if (typeof payload?.status === "boolean") {
      lcuOnline.value = payload.status;
    }
  });


  // 监听 lcuCreds 事件
  Events.On("lcuCreds", (d: any) => {
    const payload = Array.isArray(d.data) ? d.data[0] : d.data;
    console.log("[Event] lcuCreds 收到:", payload);
    if (typeof payload?.port === "number" && typeof payload?.token === "string") {
      lcuPort.value = payload.port.toString();
      lcuToken.value = payload.token;
    }
  });

  // 接收当前召唤师信息
  Events.On("summonerInfo", (d: any) => {
    console.log("[Event] summonerInfo 原始:", d);

    const info = Array.isArray(d.data) ? d.data[0] : d.data;
    console.log("[Event] 解析后的 info:", info, "类型是", typeof info);

    if (info && typeof info === "object") {
      console.log("[Event] summonerInfo 收到:", info);
      userStore.setInfo({
        nickname: info.displayName,
        avatar: info.avatar,
        region: info.region,
        tag: info.tag,
        rank: info.rank,
        winRate: info.winRate,
        wins: info.wins,
        losses: info.losses,
        totalGames: info.totalGames,
        createtime: info.createtime,
        server: info.server,
        level: info.level,
        xpSinceLastLevel: info.xpSinceLastLevel,
        xpUntilNextLevel: info.xpUntilNextLevel,
      });
    } else {
      console.warn("[Event] summonerInfo 无效数据:", info);
    }

    console.log("userStore.avatar:", userStore.avatar);
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
