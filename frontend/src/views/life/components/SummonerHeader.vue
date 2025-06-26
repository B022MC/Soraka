<template>
  <a-card class="summoner-header" :bordered="false">
    <!-- 顶部操作按钮 -->
    <div class="actions">
      <a-button type="outline" size="small" @click="onRefresh">刷新</a-button>
      <a-button type="text" size="small" @click="onShowHistory">历史战绩</a-button>
    </div>

    <!-- 进度环 + 头像 -->
    <div class="level-avatar">
      <a-progress
          type="circle"
          :percent="expPercent/100"
          :width="88"
          strokeColor="#00b42a"
          :format="() => ''"
          class="avatar-progress"
      />
      <img
          :src="userStore.avatar || defaultAvatar"
          class="avatar-img"
      />
      <div class="level-text">{{ userStore.level ?? '--' }}</div>
    </div>

    <!-- 名字 -->
    <div class="name-line">
      <span class="nickname">{{ userStore.nickname || '未知召唤师' }}</span>
      <icon-font
          v-if="userStore.privacy === 'PRIVATE'"
          type="icon-Lock"
          class="lock-icon"
      />
    </div>
    <div class="tag-line">#{{ userStore.tag || '00000' }}</div>
    <div class="rank-table-section">
      <a-table
          :columns="columns"
          :data="rankData"
          :pagination="false"
          size="small"
          row-key="type"
          class="rank-table"
          :bordered="true"
      >
      </a-table>
    </div>
  </a-card>
</template>

<script setup lang="ts">
import { computed, watch, ref, onMounted } from 'vue';
import { useUserStore } from '@/store';
import { GetRankSummary } from "/#/Soraka/internal/wails/lcu/lcuapiwails";

const userStore = useUserStore();
const defaultAvatar = new URL('@/assets/logo.png', import.meta.url).href;

const expPercent = computed(() => {
  if (
      userStore.percentCompleteForNextLevel !== undefined &&
      userStore.percentCompleteForNextLevel !== null
  ) {
    return Math.min(Math.max(userStore.percentCompleteForNextLevel, 0), 100);
  }

  const current = userStore.xpSinceLastLevel ?? 0;
  const next = userStore.xpUntilNextLevel ?? 0;
  const total = current + next;

  if (total > 0) {
    return Math.min(Math.max((current / total) * 100, 0), 100);
  }

  return 0;
});

const columns = [
  { title: "类型", dataIndex: "type", align: "center" },
  { title: "总场次", dataIndex: "total", align: "center" },
  { title: "胜率", dataIndex: "rate", align: "center" },
  { title: "胜场", dataIndex: "wins", align: "center" },
  { title: "负场", dataIndex: "loses", align: "center" },
  { title: "段位", dataIndex: "rank", align: "center" },
  { title: "胜点", dataIndex: "lp", align: "center" },
  { title: "赛季最高", dataIndex: "seasonMax", align: "center" },
  { title: "上赛季结算", dataIndex: "lastSeason", align: "center" },
];


// 用 ref 存 rank 数据
const rankData = ref<any[]>([]);

// 请求排位数据
const loadRankData = async () => {
  try {
    console.log("开始获取排位数据...");
    const res = await GetRankSummary();
    console.log("获取排位数据成功:", res);
    rankData.value = res;
  } catch (err) {
    console.error("获取排位数据失败:", err);
    rankData.value = [];  // 清空或保留之前数据视需求而定
  }
};

// 页面初始化加载
onMounted(() => {
  loadRankData();
});

// 点击刷新时重新加载
const onRefresh = () => {
  console.log('点击刷新召唤师信息');
  loadRankData();
};

const onShowHistory = () => {
  console.log('点击查看历史战绩');
};

// 监听隐私变化
watch(
    () => userStore.privacy,
    (val) => {
      console.log('privacy 变化了:', val);
    },
    { immediate: true }
);

// 监听经验相关变化
watch(
    () => [userStore.xpSinceLastLevel, userStore.xpUntilNextLevel, userStore.percentCompleteForNextLevel],
    () => {
      console.log('经验相关变化了:', {
        xpSinceLastLevel: userStore.xpSinceLastLevel,
        xpUntilNextLevel: userStore.xpUntilNextLevel,
        percentCompleteForNextLevel: userStore.percentCompleteForNextLevel,
        percent: expPercent.value
      });
    },
    { immediate: true }
);
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

  .level-avatar {
    position: relative;
    width: 88px;
    height: 88px;
    margin: 0 auto;

    .avatar-progress {
      display: block;
    }

    .avatar-img {
      position: absolute;
      top: 50%;
      left: 50%;
      width: 75px;
      height: 75px;
      border-radius: 50%;
      object-fit: cover;
      transform: translate(-50%, -50%);
      box-shadow: 0 0 4px rgba(0, 0, 0, 0.2);
    }

    .level-text {
      margin-top: 4px;
      text-align: center;
      font-weight: bold;
      font-size: 14px;
      color: #00b42a;
    }
  }

  .name-line {
    display: flex;
    justify-content: center;
    align-items: center;
    gap: 4px;
    font-weight: bold;
    font-size: 16px;
    margin-top: 20px;

    .lock-icon {
      font-size: 14px;
      color: #aaa;
    }
  }

  .tag-line {
    font-size: 12px;
    color: #888;
    text-align: center;
    margin-top: 2px;
  }
}
.rank-table {
  margin-top: 30px;

  :deep(.arco-table-th) {
    font-weight: 600;
    text-align: center;
    background: var(--table-header-bg);
    transition: background 0.3s ease;
  }

  :deep(.arco-table-td) {
    text-align: center;
    padding: 6px 8px;
  }
}
</style>
