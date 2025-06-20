<template>
  <div class="conn-status">
    <a-tag :color="online ? 'green' : 'red'">
      {{ online ? '已连接LCU' : '未连接' }}
    </a-tag>
  </div>
</template>

<script setup lang="ts">
import { ref, onMounted, onUnmounted } from 'vue'
import { LcuService } from '/#/Soraka/service'

const online = ref(false)
let timer: number | undefined

const check = async () => {
  try {
    online.value = await LcuService.CheckLogin()
  } catch {
    online.value = false
  }
}

onMounted(() => {
  check()
  timer = window.setInterval(check, 5000)
})

onUnmounted(() => {
  if (timer) window.clearInterval(timer)
})
</script>

<style scoped>
.conn-status {
  padding: 8px;
}
</style>
