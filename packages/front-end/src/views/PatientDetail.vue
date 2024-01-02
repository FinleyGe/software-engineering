<script setup lang="ts">
import { useUserStore } from "@/store";
import {AdminGetPatient} from "@/api/admin";
import { useMessage ,NButton, NSpace} from "naive-ui";
import { useRouter } from "vue-router";
import { ref, onBeforeMount, computed } from "vue";
import type { PatientDetail } from "@/type";

const router = useRouter();
const message = useMessage();
const userStore = useUserStore();
const props = defineProps<{
  id: number;
}>();

const patient = ref<PatientDetail>();

function handleRefresh() {
  switch (userStore.role) {
  case "admin":
    AdminGetPatient(props.id)
      .then((res) => {
        patient.value = res.data.data;
      })
      .catch((err) => {
        message.error(err);
      });
    break;
  case "doctor":
    break;
  case "patient":
    router.push({
      path: "/patient"
    });
    break;
  }
}

function handleCheckMonitor() {
  router.push({
    path: "/admin/monitor/" + patient.value!.bed.id
  });
}

const items = computed(
  () => {
    return {
      "ID": patient.value!.id,
      "姓名": patient.value!.name,
      "医生姓名": patient.value!.doctor.name,
      "电话": patient.value!.phone,
      "状态": patient.value!.state,
      "入院时间": patient.value!.in_time,
      "出院时间": patient.value!.out_time === "0001-01-01 00:00:00"
        ? "未出院"
        : patient.value!.out_time,
      "病房号": patient.value!.bed.number
    };
  }
);

onBeforeMount(() => {
  handleRefresh();
});

</script>
<template>
  <h2 v-if="!patient">
    信息获取失败
  </h2>
  <div
    v-else
    class="p-4 grid grid-cols-12 gap-4"
  >
    <n-space
      vertical
      class="col-span-4"
    >
      <n-space
        v-for="v, k in items"
        :key="k"
        align="center"
        justify="space-between"
      >
        <span class="text-xl"> {{ k }}: </span>
        <span class="text-xl"> {{ v }} </span>
      </n-space>
    </n-space>
    <n-space
      vertical
      class="col-start-8 col-span-2"
    >
      <n-button
        type="primary"
        class="w-full"
        round
        @click="handleRefresh"
      >
        刷新
      </n-button>
      <n-button
        type="info"
        class="w-full"
        round
        @click="handleCheckMonitor"
      >
        查看监视信息
      </n-button>
      <n-button
        type="warning"
        class="w-full"
        round
        @click="handleRefresh"
      >
        修改
      </n-button>
    </n-space>
  </div>
</template>

<style scoped lang="scss">

</style>
