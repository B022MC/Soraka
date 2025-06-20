<template>
  <div class="conn-status">
    <a-tag :color="online ? 'green' : 'red'">
      {{ online ? "已连接LCU" : "未连接" }}
    </a-tag>
  </div>
</template>

<script setup lang="ts">
import { onMounted, onUnmounted, computed } from "vue";
import { LcuService } from "/#/Soraka/service";
import { useAppStore } from "@/store";

const appStore = useAppStore();
const online = computed(() => appStore.lcuOnline);

let timer: number | undefined;

const update = async () => {
  try {
    appStore.lcuOnline = await LcuService.CheckLogin();
  } catch {
    appStore.lcuOnline = false;
  }
};

onMounted(() => {
  update();
  timer = window.setInterval(update, 5000);
});

onUnmounted(() => {
  if (timer) window.clearInterval(timer);
});
</script>

<style scoped>
.conn-status {
  padding: 8px;
}
</style>
