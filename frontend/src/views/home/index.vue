<template>
  <div class="container">
    <div class="left-side">
      <div class="panel">
        <Welcome />
      </div>
      <!--      <div class="panel" style="margin-top: 10px;">-->
      <!--        <ContentChart />-->
      <!--      </div>-->
    </div>
    <div class="right-side">
      <a-grid :cols="24" :row-gap="16">
        <a-grid-item :span="24">
          <div class="panel moduler-wrap">
            <QuickOperation />
          </div>
        </a-grid-item>
        <a-grid-item class="panel" :span="24">
          <ConnectionStatus />
        </a-grid-item>
        <!--        <a-grid-item class="panel" :span="24">-->
        <!--          <CategoriesPercent />-->
        <!--        </a-grid-item>-->
      </a-grid>
    </div>
  </div>
</template>

<script setup>
import Welcome from "./components/Welcome.vue";
import QuickOperation from "./components/quick-operation.vue";
import ConnectionStatus from "./components/ConnectionStatus.vue";
</script>
<style scoped lang="less">
.container {
  padding: 15px;
  display: flex;
  .left-side {
    flex: 1;
    overflow: auto;
    .panel {
      background-color: var(--color-neutral-1);
      border-radius: 4px;
      overflow: auto;
    }
  }
  .right-side {
    width: 280px;
    margin-left: 16px;
    .panel {
      background-color: var(--color-neutral-1);
      border-radius: 4px;
      overflow: auto;
    }
  }
}
</style>
