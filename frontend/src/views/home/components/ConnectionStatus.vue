<template>
  <div class="conn-status">
    <template v-if="online">
      <a-tag color="green">已连接LCU</a-tag>
    </template>
    <template v-else>
      <span class="waiting">等待连接</span>
    </template>
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
.waiting {
  display: inline-block;
  position: relative;
  padding-right: 1.4em;
}
.waiting::after {
  content: "";
  position: absolute;
  top: 50%;
  right: 0;
  width: 0.8em;
  height: 0.8em;
  margin-top: -0.4em;
  border: 2px solid var(--color-neutral-6);
  border-top-color: transparent;
  border-radius: 50%;
  animation: spin 1s linear infinite;
}
@keyframes spin {
  to {
    transform: rotate(360deg);
  }
}
</style>
