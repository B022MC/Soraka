<template>
  <div class="conn-status">
    <a-tag :color="online ? 'green' : 'red'">
      {{ online ? "已连接LCU" : "未连接" }}
    </a-tag>
  </div>
</template>

<script setup lang="ts">
import { computed } from "vue";
import { useAppStore } from "@/store";

const appStore = useAppStore();
const online = computed(() => appStore.lcuOnline);
</script>

<style scoped>
.conn-status {
  padding: 8px;
}
</style>
