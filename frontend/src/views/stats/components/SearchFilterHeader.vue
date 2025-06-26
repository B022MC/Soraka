<template>
  <div class="filter-header">
    <a-input v-model="keyword" placeholder="召唤师昵称" allow-clear style="width: 160px" />
    <a-select v-model="mode" style="width: 120px" class="mode-select">
      <a-option v-for="m in modes" :key="m" :value="m">{{ m }}</a-option>
    </a-select>
    <a-button type="primary" size="small" @click="handleSearch">查询</a-button>
  </div>
</template>

<script setup lang="ts">
import { ref } from 'vue'

const emit = defineEmits<{ (e: 'search', payload: { keyword: string; mode: string }): void }>()

const keyword = ref('')
const mode = ref('全部')
const modes = ['全部', '排位', '匹配', '大乱斗']

const handleSearch = () => {
  emit('search', { keyword: keyword.value.trim(), mode: mode.value })
}
</script>

<style scoped lang="less">
.filter-header {
  display: flex;
  gap: 8px;
  padding: 8px 0;
  align-items: center;
}
.mode-select {
  min-width: 80px;
}
</style>
