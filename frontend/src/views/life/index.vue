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
import { ref, computed } from 'vue'
import SummonerHeader from './components/SummonerHeader.vue'
import MatchList from './components/MatchList.vue'
import MatchFilterHeader from './components/MatchFilterHeader.vue'

// mock 数据
const matchList = ref([
  {
    id: 1,
    result: 'lose' as const,
    mode: '无限乱斗',
    kills: 13,
    deaths: 5,
    assists: 4,
    cs: 145,
    gold: 18341,
    time: '2025/06/17 23:48',
    level: 23,
    champion: 'https://ddragon.leagueoflegends.com/cdn/14.11.1/img/champion/Blitzcrank.png',
    spells: [
      'https://ddragon.leagueoflegends.com/cdn/14.11.1/img/spell/SummonerFlash.png',
      'https://ddragon.leagueoflegends.com/cdn/14.11.1/img/spell/SummonerSmite.png'
    ],
    items: [
      'https://ddragon.leagueoflegends.com/cdn/14.11.1/img/item/6692.png',
      'https://ddragon.leagueoflegends.com/cdn/14.11.1/img/item/3142.png',
      'https://ddragon.leagueoflegends.com/cdn/14.11.1/img/item/6676.png',
      'https://ddragon.leagueoflegends.com/cdn/14.11.1/img/item/3156.png',
      'https://ddragon.leagueoflegends.com/cdn/14.11.1/img/item/3036.png',
      'https://ddragon.leagueoflegends.com/cdn/14.11.1/img/item/3364.png'
    ],
    map: '召唤师峡谷'
  },
  {
    id: 2,
    result: 'lose' as const,
    mode: '无限乱斗',
    kills: 11,
    deaths: 12,
    assists: 16,
    cs: 78,
    gold: 16018,
    time: '2025/06/17 23:22',
    level: 24,
    champion: 'https://ddragon.leagueoflegends.com/cdn/14.11.1/img/champion/Amumu.png',
    spells: [
      'https://ddragon.leagueoflegends.com/cdn/14.11.1/img/spell/SummonerFlash.png',
      'https://ddragon.leagueoflegends.com/cdn/14.11.1/img/spell/SummonerIgnite.png'
    ],
    items: [
      'https://ddragon.leagueoflegends.com/cdn/14.11.1/img/item/6655.png',
      'https://ddragon.leagueoflegends.com/cdn/14.11.1/img/item/3020.png',
      'https://ddragon.leagueoflegends.com/cdn/14.11.1/img/item/4645.png',
      'https://ddragon.leagueoflegends.com/cdn/14.11.1/img/item/3157.png',
      'https://ddragon.leagueoflegends.com/cdn/14.11.1/img/item/3916.png',
      'https://ddragon.leagueoflegends.com/cdn/14.11.1/img/item/3363.png'
    ],
    map: '召唤师峡谷'
  },
  {
    id: 3,
    result: 'win' as const,
    mode: '匹配模式',
    kills: 8,
    deaths: 2,
    assists: 10,
    cs: 130,
    gold: 15000,
    time: '2025/06/16 20:00',
    level: 21,
    champion: 'https://ddragon.leagueoflegends.com/cdn/14.11.1/img/champion/Ahri.png',
    spells: [
      'https://ddragon.leagueoflegends.com/cdn/14.11.1/img/spell/SummonerFlash.png',
      'https://ddragon.leagueoflegends.com/cdn/14.11.1/img/spell/SummonerHeal.png'
    ],
    items: [
      'https://ddragon.leagueoflegends.com/cdn/14.11.1/img/item/6655.png',
      'https://ddragon.leagueoflegends.com/cdn/14.11.1/img/item/3020.png',
      'https://ddragon.leagueoflegends.com/cdn/14.11.1/img/item/3165.png',
      'https://ddragon.leagueoflegends.com/cdn/14.11.1/img/item/3089.png',
      'https://ddragon.leagueoflegends.com/cdn/14.11.1/img/item/1058.png',
      'https://ddragon.leagueoflegends.com/cdn/14.11.1/img/item/3363.png'
    ],
    map: '召唤师峡谷'
  }
])

// 当前选中的筛选模式
const selectedMode = ref('全部')

// 筛选后的比赛列表
const filteredMatches = computed(() => {
  if (selectedMode.value === '全部') return matchList.value
  return matchList.value.filter(m => m.mode === selectedMode.value)
})

// 响应筛选模式变化
function onFilterChange(mode: string) {
  selectedMode.value = mode
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
