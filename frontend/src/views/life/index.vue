<template>
  <div class="life-page">
    <!-- 召唤师信息卡片 -->
    <SummonerHeader />

    <!-- 筛选头部 -->
    <MatchFilterHeader :matches="matchList" @filterChange="onFilterChange" />

    <!-- 可滚动区域 -->
    <div class="match-scroll-container">
      <MatchList :matches="filteredMatches" />
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref, computed, onMounted } from "vue";
import { getRecentMatches } from "@/api/lcu";
import SummonerHeader from "./components/SummonerHeader.vue";
import MatchList from "./components/MatchList.vue";
import MatchFilterHeader from "./components/MatchFilterHeader.vue";

const matchList = ref(
  [] as ReturnType<typeof getRecentMatches> extends Promise<infer R>
    ? R
    : never,
);

onMounted(async () => {
  try {
    matchList.value = await getRecentMatches();
  } catch (e) {
    console.error(e);
  }
});

// 当前选中的筛选模式
const selectedMode = ref("全部");

// 筛选后的比赛列表
const filteredMatches = computed(() => {
  if (selectedMode.value === "全部") return matchList.value;
  return matchList.value.filter((m) => m.mode === selectedMode.value);
});

// 响应筛选模式变化
function onFilterChange(mode: string) {
  selectedMode.value = mode;
}
</script>

<style scoped lang="less">
.life-page {
  display: flex;
  flex-direction: column;
  height: calc(100vh - 32px); // 16px padding * 2
  //height: 100vh; // 填满视口
  padding: 16px;
  box-sizing: border-box;
  background: var(--color-bg-1);
  color: var(--color-text);
  overflow: hidden;

  // 召唤师信息和筛选组件区域撑开内容
  > :not(.match-scroll-container) {
    flex-shrink: 0;
  }

  .match-scroll-container {
    flex: 1; // 占满剩余空间
    overflow-y: auto;
    min-height: 0; // 防止内容撑高
  }
}
</style>
