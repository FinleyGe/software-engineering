<script setup lang="ts">
import {NButton, NIcon, NInput, NSpace} from "naive-ui";
import {Checkmark} from "@vicons/carbon";
import { ref } from "vue";
// const value = defineModel<string>("value");
// const isEdit = defineModel<boolean>("isedit", {default: false});
type Props = {
  value: string
}
const props = defineProps<Props>();
const value = ref<string>(props.value);
const isEdit = ref<boolean>(false);
const emits = defineEmits<{
  (e:"submit", value: string): void
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
        <n-input
          v-if="isEdit"
          v-model:value="value"
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
      {{ value }}
    </span>
  </div>
</template>

<style scoped lang="scss">

</style>
