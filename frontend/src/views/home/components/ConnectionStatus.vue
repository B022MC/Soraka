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
import { ref, onMounted, onBeforeUnmount } from "vue";
import { LcuService } from "/#/Soraka/service";
import { SSEClient } from "@/utils/sse";

const online = ref(false);
let sse: SSEClient | null = null;

onMounted(() => {
  LcuService.CheckLogin().then((ok) => {
    online.value = ok;
  });
  sse = new SSEClient("http://localhost:8233/events");
  sse.connect({
    lcuStatus: (d) => {
      online.value = !!d;
    },
  });
});

onBeforeUnmount(() => {
  sse?.close();
});
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
