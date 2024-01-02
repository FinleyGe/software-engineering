<script setup lang="ts">
import {NButton, NInput} from "naive-ui";
import {ref} from "vue";
const emits = defineEmits<{
  (e:"close"):void,
}>();

function close() {
  emits("close");
}

const departmentName = ref<string>("");

</script>
<template>
  <div
    class="fixed inset-0 z-40 flex items-center justify-center bg-gray-900 bg-opacity-50 backdrop-blur"
  >
    <div class="flex flex-col bg-white rounded-lg shadow-lg p-10 z-50">
      <header>
        <h1 class="text-xl">
          添加科室
        </h1>
      </header>

      <n-input
        v-model:value="departmentName"
        placeholder="科室名称"
      />
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
