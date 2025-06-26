<template>
  <div class="pagination">
    <a-pagination
      :current="page"
      :total="total"
      :page-size="pageSize"
      @change="p => emit('change', p)"
      size="small"
      show-jumper
    />
  </div>
</template>

<script setup lang="ts">
const props = defineProps<{ page: number; total: number; pageSize?: number }>()
const emit = defineEmits<{ (e: 'change', page: number): void }>()

const pageSize = props.pageSize ?? 10
</script>

<style scoped lang="less">
.pagination {
  display: flex;
  justify-content: center;
  padding: 12px 0;
}
</style>
