<script setup lang="ts">
import {NButton, NIcon, NSelect, NSpace} from "naive-ui";
import {Checkmark} from "@vicons/carbon";
import { ref } from "vue";
// const value = defineModel<string>("value");
// const isEdit = defineModel<boolean>("isedit", {default: false});
type Props = {
  value: number
  options: Array<{label: string, value: number}>
}
const props = defineProps<Props>();
const value = ref<number>(props.value);
const isEdit = ref<boolean>(false);
const emits = defineEmits<{
  (e:"submit", value: number): void
}>();

function submit() {
  emits("submit", value.value);
  isEdit.value = false;
}

</script>
<template>
  <div style="width:fit-content">
    <template v-if="isEdit">
      <n-space>
        <n-select
          v-if="isEdit"
          v-model:value="value"
          :options="options"
        />
        <n-button
          v-if="isEdit"
          type="primary"
          circle
          @click="submit()"
        >
          <template #icon>
            <n-icon>
              <Checkmark />
            </n-icon>
          </template>
        </n-button>
      </n-space>
    </template>
    <span
      v-else
      @click="isEdit = true"
    >
      {{ options.find((item) => item.value === value)?.label }}
    </span>
  </div>
</template>

<style scoped lang="scss">

</style>
