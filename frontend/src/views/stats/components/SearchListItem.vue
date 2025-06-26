<template>
  <a-card class="match-item" :class="{ win: match.result === 'win', lose: match.result === 'lose' }" :bordered="false">
    <div class="match-content">
      <div class="left">
        <div class="avatar-wrapper">
          <img :src="match.champion" class="champion-avatar" />
          <span class="level">{{ match.level }}</span>
        </div>
        <div class="info-wrapper">
          <span class="result" :class="match.result">{{ match.result === 'win' ? '胜利' : '失败' }}</span>
          <span class="mode">{{ match.mode }}</span>
          <div class="spells">
            <img v-for="(spell, idx) in match.spells" :key="idx" :src="spell" class="spell-icon" />
          </div>
          <div class="runes">
            <img v-for="(rune, idx) in match.runes" :key="idx" :src="rune" class="rune-icon" />
          </div>
        </div>
      </div>
      <div class="middle">
        <div class="kda-line">
          <strong>{{ match.kills }}</strong> /
          <strong class="deaths">{{ match.deaths }}</strong> /
          <strong>{{ match.assists }}</strong>
        </div>
        <span class="cs">{{ match.cs }} 补刀</span>
        <div class="items">
          <img v-for="(item, idx) in match.items" :key="idx" :src="item" class="item-icon" />
        </div>
        <span class="gold">{{ match.gold.toLocaleString() }} 金币</span>
        <span class="damage">{{ match.damage.toLocaleString() }} 伤害</span>
      </div>
      <div class="right">
        <div class="team-stats">
          <span>塔 {{ match.towers }}</span>
          <span>水晶 {{ match.inhibitor }}</span>
          <span>男爵 {{ match.baron }}</span>
          <span>巨龙 {{ match.dragon }}</span>
          <span>先锋 {{ match.herald }}</span>
          <span>巢虫 {{ match.vilemaw }}</span>
        </div>
        <div class="meta">
          <span class="map">{{ match.map }}</span>
          <span class="time">{{ match.time }}</span>
        </div>
      </div>
    </div>
  </a-card>
</template>

<script setup lang="ts">
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

defineProps<{ match: MatchDetail }>()
</script>

<style scoped lang="less">
.match-item {
  border-radius: 6px;
  padding: 12px;
  box-shadow: 0 0 4px var(--color-shadow);
  transition: transform 0.2s;
  &:hover {
    transform: translateY(-2px);
  }
  &.win {
    background-color: var(--color-success-light-1);
    border: 1px solid var(--color-success-light-3);
  }
  &.lose {
    background-color: var(--color-danger-light-1);
    border: 1px solid var(--color-danger-light-3);
  }
  .match-content {
    display: flex;
    justify-content: space-between;
    flex-wrap: wrap;
    gap: 12px;
  }
  .left {
    display: flex;
    gap: 12px;
    .avatar-wrapper {
      position: relative;
      .champion-avatar {
        width: 48px;
        height: 48px;
        border-radius: 8px;
      }
      .level {
        position: absolute;
        bottom: -8px;
        right: -8px;
        background: var(--color-bg-1);
        font-size: 12px;
        width: 20px;
        height: 20px;
        line-height: 20px;
        text-align: center;
        border-radius: 50%;
        box-shadow: 0 0 2px var(--color-shadow);
      }
    }
    .info-wrapper {
      display: flex;
      flex-direction: column;
      justify-content: center;
      .result {
        font-weight: bold;
        font-size: 14px;
        &.win {
          color: var(--color-success);
        }
        &.lose {
          color: var(--color-danger);
        }
      }
      .mode {
        font-size: 12px;
        color: var(--color-text-2);
        margin: 2px 0;
      }
      .spells,
      .runes {
        display: flex;
        gap: 4px;
        margin-top: 2px;
        .spell-icon,
        .rune-icon {
          width: 20px;
          height: 20px;
          border-radius: 3px;
        }
      }
    }
  }
  .middle {
    display: flex;
    align-items: center;
    flex-wrap: wrap;
    gap: 12px;
    .kda-line {
      font-size: 13px;
      .deaths {
        color: var(--color-danger);
        font-weight: bold;
      }
    }
    .cs,
    .gold,
    .damage {
      font-size: 13px;
      color: var(--color-text-2);
    }
    .items {
      display: flex;
      gap: 4px;
      .item-icon {
        width: 26px;
        height: 26px;
        border-radius: 4px;
        background-color: var(--color-fill-2);
      }
    }
  }
  .right {
    text-align: right;
    min-width: 120px;
    display: flex;
    flex-direction: column;
    justify-content: space-between;
    .team-stats {
      display: flex;
      flex-direction: column;
      font-size: 12px;
      color: var(--color-text-2);
    }
    .meta {
      font-size: 12px;
      color: var(--color-text-2);
      margin-top: 4px;
      display: flex;
      flex-direction: column;
      align-items: flex-end;
    }
  }
}
</style>
