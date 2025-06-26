<template>
  <div class="stats-page">
    <SearchFilterHeader @search="onSearch" />
    <SearchList :matches="pageMatches" />
    <SearchPagination :page="page" :total="total" @change="onPageChange" />
  </div>
</template>

<script setup lang="ts">
import { ref, computed } from 'vue'
import SearchFilterHeader from './components/SearchFilterHeader.vue'
import SearchList from './components/SearchList.vue'
import SearchPagination from './components/SearchPagination.vue'

interface MatchDetail {
  id: number | string
  result: string
  mode: string
  time: string
  champion: string
  spells: string[]
  runes: string[]
  items: string[]
  kills: number
  deaths: number
  assists: number
  cs: number
  gold: number
  damage: number
  level: number
  towers: number
  inhibitor: number
  baron: number
  dragon: number
  herald: number
  vilemaw: number
  map: string
}

const matches = ref<MatchDetail[]>([])
const page = ref(1)
const pageSize = 5
const total = ref(0)

const pageMatches = computed(() => {
  const start = (page.value - 1) * pageSize
  return matches.value.slice(start, start + pageSize)
})

const sampleImg = new URL('@/assets/logo.png', import.meta.url).href
const mockData: MatchDetail[] = Array.from({ length: 8 }).map((_, idx) => ({
  id: idx + 1,
  result: idx % 2 === 0 ? 'win' : 'lose',
  mode: '匹配模式',
  time: '2023/01/01 12:00',
  champion: sampleImg,
  spells: [sampleImg, sampleImg],
  runes: [sampleImg, sampleImg],
  items: [sampleImg, sampleImg, sampleImg, sampleImg, sampleImg, sampleImg],
  kills: 5,
  deaths: 3,
  assists: 8,
  cs: 120,
  gold: 12345,
  damage: 23456,
  level: 15,
  towers: 8,
  inhibitor: 2,
  baron: 1,
  dragon: 2,
  herald: 1,
  vilemaw: 0,
  map: '召唤师峡谷'
}))

function onSearch() {
  matches.value = mockData
  total.value = mockData.length
  page.value = 1
}

function onPageChange(p: number) {
  page.value = p
}
</script>

<style scoped lang="less">
.stats-page {
  padding: 16px;
  display: flex;
  flex-direction: column;
  gap: 12px;
  height: calc(100vh - 32px);
  overflow: hidden;
  > *:not(:last-child) {
    flex-shrink: 0;
  }
  .search-list {
    flex: 1;
    overflow-y: auto;
  }
}
</style>
