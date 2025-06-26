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
          /></a-avatar>
        <div class="welcome">
          <p class="hello">{{ goodTimeText() }}！{{ userStore.nickname }}</p>
          <p>
            当前时间：<span class="datetime">
    {{ appStore.system.sysTime || "---" }}
  </span>
          </p>
          <p>
            客户端路径：{{ appStore.client.clientPath || "未找到" }}
            <span class="status">
    <a-tag v-if="appStore.lcu.online" color="green">已连接LCU</a-tag>
    <span v-else class="waiting">等待连接</span>
  </span>
          </p>
          <p v-if="appStore.lcu.port">
            端口：{{ appStore.lcu.port }} Token：{{ appStore.lcu.token }}
          </p>

        </div>
      </a-space>
    </a-row>
  </a-card>
</template>

<script lang="ts" setup>
import {ref, computed, watch, onMounted} from "vue";
import {useRouter, useRoute} from "vue-router";
import routerMap from "@/router/routerMap";
import {Events} from "@wailsio/runtime";
import {LcuApiWails} from "/#/Soraka/internal/wails/lcu";
import {useUserStore} from "@/store";
import {useAppStore} from "@/store";
import {goodTimeText} from "@/utils";

const router = useRouter();
const route = useRoute();
const current = ref(route.name);
const userStore = useUserStore();
const appStore = useAppStore();

const menulist = computed(
    () => routerMap.find((item) => item.path === "/")?.children || [],
);

const defaultAvatar = new URL('@/assets/logo.png', import.meta.url).href;

watch(
    () => route.name,
    (name) => (current.value = name),
);

const menuHandle = (key: string) => {
  current.value = key;
  const target = menulist.value.find((i) => i.name === key);
  if (target) {
    router.push({name: target.name});
  }
};

const handleSetting = () => {
  current.value = "setting";
  router.push({name: "setting"});
};

const handleTool = (type: string) => {
  console.log("工具点击", type);
};

onMounted(() => {
  Events.On("time", (time: any) => {
    const t = Array.isArray(time.data) ? time.data[0] : time.data;
    appStore.setSysTime(t);
  });

  LcuApiWails.GetClientPath()
      .then((p: string) => {
        appStore.setClientPath(p);
      })
      .catch(() => {
        appStore.setClientPath("");
      });

  Events.On("lcuStatus", (d: any) => {
    const payload = Array.isArray(d.data) ? d.data[0] : d.data;
    if (typeof payload?.status === "boolean") {
      appStore.setLcuStatus(payload.status);
    }
  });

  Events.On("lcuCreds", (d: any) => {
    const payload = Array.isArray(d.data) ? d.data[0] : d.data;
    if (typeof payload?.port === "number" && typeof payload?.token === "string") {
      appStore.setLcuCreds(payload.port.toString(), payload.token);
    }
  });

  Events.On("summonerInfo", (d: any) => {
    const info = Array.isArray(d.data) ? d.data[0] : d.data;
    if (info && typeof info === "object") {
      userStore.setInfo({
        accountId: info.accountId,
        summonerId: info.summonerId,
        puuid: info.puuid,
        nickname: info.gameName || info.displayName,
        avatar: info.avatarUrl,
        region: info.region || "",
        tag: info.tagLine,
        rank: info.rank || "",
        winRate: info.winRate || 0,
        wins: info.wins || 0,
        losses: info.losses || 0,
        totalGames: info.totalGames || 0,
        createtime: info.createtime || "",
        level: info.summonerLevel,
        xpSinceLastLevel: info.xpSinceLastLevel,
        xpUntilNextLevel: info.xpUntilNextLevel,
        percentCompleteForNextLevel: info.percentCompleteForNextLevel,
        privacy: info.privacy,
      });
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
