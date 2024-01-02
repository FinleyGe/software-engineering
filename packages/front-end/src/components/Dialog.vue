<script setup lang="ts">
import {NButton} from "naive-ui";
const emits = defineEmits<{
  (e:"close"):void,
}>();

function close() {
  emits("close");
}

</script>
<template>
  <div
    class="fixed inset-0 z-40 flex items-center justify-center bg-gray-900 bg-opacity-50 backdrop-blur"
  >
    <div class="flex flex-col bg-white rounded-lg shadow-lg p-10 z-50">
      <header />

      <footer class="mt-2 flex flex-row justify-around">
        <n-button
          type="error"
          @click="close"
        >
          取消
        </n-button>
      </footer>
    </div>
  </div>
</template>

<style scoped lang="scss">

</style>
