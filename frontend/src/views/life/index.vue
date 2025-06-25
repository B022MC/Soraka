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
import { LcuApiWails } from "/#/Soraka/internal/wails/lcu";
import SummonerHeader from "./components/SummonerHeader.vue";
import MatchList from "./components/MatchList.vue";
import MatchFilterHeader from "./components/MatchFilterHeader.vue";
import {ListRecentMatches} from "/#/Soraka/internal/wails/lcu/lcuapiwails";

interface MatchBrief {
  id: number;
  result: string;
  mode: string;
  kills: number;
  deaths: number;
  assists: number;
  cs: number;
  gold: number;
  time: string;
  level: number;
  champion: string;
  spells: string[];
  items: string[];
  map: string;
}
function mapResult(result: string): "win" | "lose" {
  return result === "胜利" ? "win" : "lose";
}
const matchList = ref<MatchBrief[]>([]);

const selectedMode = ref("全部");

const filteredMatches = computed(() => {
  if (selectedMode.value === "全部") return matchList.value;
  return matchList.value.filter((m) => m.mode === selectedMode.value);
});

function onFilterChange(mode: string) {
  selectedMode.value = mode;
}

onMounted(async () => {
  try {
    const p = ListRecentMatches(20);
    const data = await p;
    matchList.value = (data ?? [])
        .filter((m): m is MatchBrief => m !== null)
        .map(m => ({
          ...m,
          result: mapResult(m.result)
        }));
  } catch (e) {
    console.error("获取比赛数据失败", e);
  }
});
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
