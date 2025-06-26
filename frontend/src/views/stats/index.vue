<template>
  <div class="stats-page">
    <SearchFilterHeader @filter-change="" />
    <SearchList :matches="matches" />
  </div>
</template>

<script setup lang="ts">
import { ref, onMounted } from 'vue';
import SearchFilterHeader from './components/SearchFilterHeader.vue';
import SearchList from './components/SearchList.vue';

const matches = ref([]);


onMounted(() => {
  // 初始化加载战绩数据
});
</script>

<style scoped>
.stats-page {
  padding: 16px;
}
</style>
