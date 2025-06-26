<template>
  <a-card class="match-wrapper" :bordered="false">
    <div class="match-scroll-list">
      <a-card
          v-for="match in matches"
          :key="match.id"
          class="match-item"
          :class="{ win: match.result === 'win', lose: match.result === 'lose' }"
          :bordered="false"
      >
        <div class="match-content">
          <!-- 左侧：头像、等级、胜负、模式、召唤师技能 -->
          <div class="left">
            <div class="left-top">
              <div class="avatar-wrapper">
                <img :src="match.champion" class="champion-avatar" />
                <span class="level">{{ match.level }}</span>
              </div>
              <div class="info-wrapper">
                <span class="result" :class="match.result">{{ match.result === 'win' ? '胜利' : '失败' }}</span>
                <span class="mode">{{ match.mode_detail }}</span>
                <div class="spells">
                  <img v-for="(spell, idx) in match.spells" :key="idx" :src="spell" class="spell-icon" />
                </div>
              </div>
            </div>
          </div>

          <!-- 中间：KDA、补兵、装备、金币 -->
          <div class="middle">
            <div class="middle-row">
              <div class="kda-line">
                <strong>{{ match.kills }}</strong>
                /<strong class="deaths">{{ match.deaths }}</strong>
                /<strong>{{ match.assists }}</strong>
              </div>
              <span class="cs">{{ match.cs }} 补刀</span>
              <div class="items">
                <img v-for="(item, idx) in match.items" :key="idx" :src="item" class="item-icon" />
              </div>
              <span class="gold">{{ match.gold.toLocaleString() }} 金币</span>
            </div>
          </div>


          <!-- 右侧：地图与时间 -->
          <div class="right">
            <div class="meta">
              <span class="map">{{ match.map }}</span>
              <span class="time">{{ match.time }}</span>
            </div>
          </div>
        </div>
      </a-card>
    </div>
  </a-card>
</template>

<script setup lang="ts">
defineProps<{
  matches: {
    id: string | number
    result: string
    level: number
    spells: string[]
    mode: string
    mode_detail:string
    kills: number
    deaths: number
    assists: number
    cs: number
    gold: number
    items: string[]
    map: string
    time: string
    champion: string
  }[]
}>()
</script>

<style scoped lang="less">
.match-wrapper {
  width: 95%;
  max-height: 90%;
  overflow-y: auto;
  margin: 0 auto;    // ✅ 居中显示
  background: color-mix(in srgb, var(--color-bg-2), transparent 30%);
  backdrop-filter: blur(10px);
  -webkit-backdrop-filter: blur(10px);
  border-radius: 12px;
  box-shadow: 0 2px 8px var(--color-shadow);
  border: 1px solid var(--color-border);
  // 鼠标悬停才显示滚动条
  &::-webkit-scrollbar {
    width: 6px;
    height: 6px;
  }
  &::-webkit-scrollbar-thumb {
    background-color: var(--color-fill-4); // Arco Design 次填充色
    border-radius: 4px;
    border: 2px solid transparent;
    background-clip: content-box;
  }
  // 鼠标不悬停时隐藏滚动条
  &:not(:hover) {
    scrollbar-width: none;

    &::-webkit-scrollbar {
      display: none;
    }
  }
  .match-scroll-list {
    display: flex;
    flex-direction: column;
    gap: 12px;
    padding: 12px 16px 8px 16px;
    scrollbar-width: none;
    -ms-overflow-style: none;

    &::-webkit-scrollbar {
      display: none;
    }
  }

  .match-item {
    max-width: 960px;
    width: 100%;
    margin: 0 auto;
    border-radius: 6px;
    padding: 12px;
    transition: transform 0.2s;
    box-shadow: 0 0 4px var(--color-shadow);

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
      align-items: center;
      justify-content: space-between;
      flex-wrap: wrap;
    }

    .left {
      display: flex;
      flex-direction: column;
      justify-content: center;
      width: 160px;

      .left-top {
        display: flex;
      }

      .avatar-wrapper {
        position: relative;
        margin-right: 12px;

        .champion-avatar {
          width: 48px;
          height: 48px;
          border-radius: 8px;
        }

        .level {
          position: absolute;
          bottom: -8px;
          right: -8px;
          font-size: 12px;
          color: var(--color-text-1);
          background: var(--color-bg-1);
          border-radius: 50%;
          width: 20px;
          height: 20px;
          line-height: 20px;
          text-align: center;
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

        .spells {
          display: flex;
          gap: 4px;

          .spell-icon {
            width: 20px;
            height: 20px;
            border-radius: 3px;
          }
        }
      }
    }

    .middle {
      flex: 1;
      padding-left: 16px;

      .middle-row {
        display: flex;
        align-items: center;
        flex-wrap: wrap;
        gap: 12px;
      }

      .kda-line {
        font-size: 13px;
        color: var(--color-text-2);

        .deaths {
          color: var(--color-danger);
          font-weight: bold;
        }
      }

      .cs,
      .gold {
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
      min-width: 110px;

      .meta {
        font-size: 12px;
        color: var(--color-text-2);
        display: flex;
        flex-direction: column;
        align-items: flex-end;
      }
    }
  }
}

</style>
