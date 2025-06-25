<template>
  <div class="filter-header">
    <div class="summary">
      <span>近期对局（最近 {{ total }} 场）</span>
      <span class="summary-win">胜：{{ winCount }}</span>
      <span class="summary-lose">负：{{ loseCount }}</span>
      <span>KDA：{{ kdaStr }}</span>
    </div>

    <div class="avatars">
      <img v-for="(hero, index) in recentChampions" :key="index" :src="hero" class="champion-icon" />
    </div>

    <div class="tools">
      <a-button size="small" type="outline">最近队友</a-button>
      <a-dropdown>
        <a-button size="small">全部</a-button>
        <template #content>
          <a-doption v-for="mode in filterModes" :key="mode" @click="() => emit('filterChange', mode)">
            {{ mode }}
          </a-doption>
        </template>
      </a-dropdown>
    </div>
  </div>
</template>

<script setup lang="ts">
import { computed } from 'vue'

const props = defineProps<{
  matches: {
    result: string
    kills: number
    deaths: number
    assists: number
    champion: string
    mode: string
  }[]
}>()

const emit = defineEmits<{
  (e: 'filterChange', mode: string): void
}>()

const total = computed(() => props.matches.length)
const winCount = computed(() => props.matches.filter(m => m.result === 'win').length)
const loseCount = computed(() => props.matches.filter(m => m.result === 'lose').length)

const kdaStr = computed(() => {
  const k = props.matches.reduce((sum, m) => sum + m.kills, 0)
  const d = props.matches.reduce((sum, m) => sum + m.deaths, 0)
  const a = props.matches.reduce((sum, m) => sum + m.assists, 0)
  const ratio = d === 0 ? (k + a) : ((k + a) / d).toFixed(1)
  return `${k} / ${d} / ${a} (${ratio})`
})

const recentChampions = computed(() => {
  const seen = new Set()
  return props.matches
      .map(m => m.champion)
      .filter(c => {
        if (seen.has(c)) return false
        seen.add(c)
        return true
      })
      .slice(0, 10)
})

const filterModes = ['全部', '匹配模式', '极地大乱斗', '单 / 双排', '灵活排位']
</script>

<style scoped lang="less">
.filter-header {
  width: 92%;
  padding: 8px 12px;
  margin: 0 auto 12px;

  background: color-mix(in srgb, var(--color-bg-2), transparent 30%);
  backdrop-filter: blur(10px);
  -webkit-backdrop-filter: blur(10px);
  border-radius: 12px;
  box-shadow: 0 2px 8px var(--color-shadow);
  border: 1px solid var(--color-border);

  display: flex;
  flex-wrap: wrap;
  align-items: center;
  justify-content: space-between;
  row-gap: 8px;

  .summary {
    display: flex;
    gap: 12px;
    font-size: 13px;
    color: var(--color-text-1);

    .summary-win {
      color: var(--color-success);
      font-weight: 500;
    }

    .summary-lose {
      color: var(--color-danger);
      font-weight: 500;
    }
  }

  .avatars {
    display: flex;
    gap: 4px;
    flex-wrap: wrap;
    padding-top: 4px;

    .champion-icon {
      width: 28px;
      height: 28px;
      border-radius: 50%;
      border: 1px solid var(--color-border);
      background-color: var(--color-fill-2);
    }
  }

  .tools {
    display: flex;
    gap: 8px;
    padding-top: 4px;

    :deep(.arco-btn) {
      font-size: 12px;
    }

    :deep(.arco-dropdown) {
      font-size: 12px;
    }
  }
}


</style>
