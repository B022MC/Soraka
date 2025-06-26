<template>
  <div class="search-list">
    <SearchListItem v-for="match in matches" :key="match.id" :match="match" />
    <div v-if="!matches.length" class="empty">暂无数据</div>
  </div>
</template>

<script setup lang="ts">
import SearchListItem from './SearchListItem.vue'

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

defineProps<{ matches: MatchDetail[] }>()
</script>

<style scoped lang="less">
.search-list {
  display: flex;
  flex-direction: column;
  gap: 12px;
}
.empty {
  text-align: center;
  color: var(--color-text-3);
  padding: 20px 0;
}
</style>
