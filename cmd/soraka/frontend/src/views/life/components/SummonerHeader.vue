<template>
  <a-card class="summoner-header" :bordered="false">
    <!-- 顶部操作按钮 -->
    <div class="actions">
      <a-button type="outline" size="small" @click="refresh">刷新</a-button>
      <a-button type="text" size="small">历史战绩</a-button>
    </div>

    <!-- 居中头像 + 等级环 -->
    <div class="top-section">
      <div class="level-avatar">
        <a-progress
            type="circle"
            :percent="expPercent"
            :width="88"
            strokeColor="#00b42a"
            :format="() => ''"
            class="avatar-progress"
        >
          <a-avatar :size="64">
            <img :src="userStore.avatar || defaultAvatar" />
          </a-avatar>
        </a-progress>
        <div class="level-text">{{ userStore.level ?? '--' }}</div>
      </div>

      <div class="name-line">
        <span class="nickname">{{ userStore.nickname || "未知召唤师" }}</span>
        <icon-lock class="lock-icon" />
      </div>
      <div class="tag-line">#{{ userStore.tag || "00000" }}</div>
    </div>

    <!-- 段位数据表格 -->
    <div class="rank-table-section">
      <a-table
          :columns="columns"
          :data="rankData"
          :pagination="false"
          size="small"
          class="rank-table"
      />
    </div>
  </a-card>
</template>

<script setup lang="ts">
import { computed } from "vue";
import { useUserStore } from "@/store";
import { IconLock } from "@arco-design/web-vue/es/icon";

const userStore = useUserStore();

const defaultAvatar = new URL("@/assets/logo.png", import.meta.url).href;

// 模拟经验百分比计算（真实情况你应从 LCU 获取经验）
const expPercent = computed(() => {
  const current = userStore.xpSinceLastLevel || 0;
  const next = userStore.xpUntilNextLevel || 1;
  return Math.floor((current / (current + next)) * 100);
});

// 表格字段
const columns = [
  { title: "类型", dataIndex: "type" },
  { title: "总场次", dataIndex: "total" },
  { title: "胜率", dataIndex: "rate" },
  { title: "胜场", dataIndex: "wins" },
  { title: "负场", dataIndex: "loses" },
  { title: "段位", dataIndex: "rank" },
  { title: "胜点", dataIndex: "lp" },
  { title: "赛季最高", dataIndex: "seasonMax" },
  { title: "上赛季结算", dataIndex: "lastSeason" },
];

// 静态演示用的数据
const rankData = [
  {
    type: "单 / 双排",
    total: 9,
    rate: "22%",
    wins: 2,
    loses: 7,
    rank: "华贵铂金 IV",
    lp: 78,
    seasonMax: "华贵铂金 IV",
    lastSeason: "--",
  },
  {
    type: "灵活排位",
    total: 36,
    rate: "41%",
    wins: 15,
    loses: 21,
    rank: "流光翡翠 IV",
    lp: 98,
    seasonMax: "流光翡翠 II",
    lastSeason: "流光翡翠 IV",
  },
];

const refresh = () => {
  console.log("点击刷新召唤师信息");
};
</script>

<style scoped lang="less">
.summoner-header {
  padding: 16px;
  position: relative;

  .actions {
    position: absolute;
    top: 12px;
    right: 16px;
    display: flex;
    gap: 8px;
  }

  .top-section {
    display: flex;
    flex-direction: column;
    align-items: center;
    margin-bottom: 16px;

    .level-avatar {
      display: flex;
      flex-direction: column;
      align-items: center;
      margin-bottom: 6px;

      .avatar-progress {
        margin-bottom: 4px;
      }

      .level-text {
        margin-top: -8px;
        font-weight: 600;
        font-size: 16px;
        color: var(--color-success-6);
      }
    }

    .name-line {
      font-size: 20px;
      font-weight: 600;
      display: flex;
      align-items: center;
      gap: 6px;

      .lock-icon {
        font-size: 16px;
        color: var(--color-text-3);
      }
    }

    .tag-line {
      font-size: 13px;
      color: var(--color-text-2);
      margin-top: 4px;
    }
  }

  .rank-table-section {
    .rank-table {
      :deep(.arco-table-th), :deep(.arco-table-td) {
        padding: 6px 8px;
        font-size: 13px;
        white-space: nowrap;
      }
    }
  }
}
</style>
