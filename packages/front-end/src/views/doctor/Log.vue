<script setup lang="ts">
import {DoctorAddLog,DoctorGetLogs, DoctorGetAllLogs, DoctorGetPatients} from "@/api/doctor";
import { Log } from "@/type/log";
import { Patient } from "@/type";
import { h,onMounted,ref } from "vue";
import {NSpace,NSelect, NButton, NDataTable, NInput} from "naive-ui";
import type {SelectOption,DataTableColumns} from "naive-ui";

const log = ref<Log>({
  patient_id: null,
  content: ""
});
const isCheck = ref<boolean>(false);
function handleCheck(id: number){
  log.value = data.value.find(item => item.id === id);
  isCheck.value = true;
}
const columns = ref<DataTableColumns> ([
  {
    title: "ID",
    key: "id",
  },
  {
    title: "患者",
    key: "patient_id",
  },
  {
    title: "操作",
    key: "action",
    render(row: Log){
      return h(
        NSpace,
        {
          style: {
            width: "100%"
          }
        },
        {
          default(){
            return [
              h(
                NButton,{
                  type: "info",
                  onClick(){
                    handleCheck(row.id);
                  }
                },
                "查看")
            ];
          }
        }
      );
    }
  }
]);
const isAdd = ref<boolean>(false);
const options = ref<SelectOption[]>([]);
function handleAdd(){
  DoctorAddLog(log.value.patient_id, log.value.content).then(res => {
    if(res.data.msg === "OK"){
      isAdd.value = false;
      handleRefresh();
    }
  });
}

const data = ref<Log[]>([]);
function handleRefresh(){
  DoctorGetAllLogs().then(res => {
    data.value = res.data.data;
  });

  DoctorGetPatients().then(res => {
    options.value = res.data.data.map((item: Patient) => {
      return {
        label: item.name,
        value: item.id
      };
    });
  });
}

onMounted(() => {
  handleRefresh();
});

</script>
<template>
  <div>
    <!-- {{ data }} -->
    <h1>日志管理</h1>
    <n-space>
      <n-button
        type="primary"
        @click="handleRefresh"
      >
        刷新
      </n-button>
      <n-button
        type="success"
        @click="isAdd = !isAdd"
      >
        添加
      </n-button>
    </n-space>
    <n-space
      v-if="isAdd"
      vertical
    >
      <n-input
        v-model:value="log.content"
        type="textarea"
        style="min-width: 800px;"
        placeholder="请输入病例内容"
      />
      <n-space>
        <n-select
          v-model:value="log.patient_id"
          placeholder="请选择患者"
          clearable
          style="min-width: 200px;"
          :options="options"
        />
        <n-button
          type="primary"
          @click="handleAdd"
        >
          添加
        </n-button>
      </n-space>
    </n-space>
    <div
      v-if="isCheck"
      class="text-xl m-2 p-2 flex flex-col bg-gray-100 rounded-xl"
    >
      {{ log.content }}
      <n-button
        type="primary"
        class="mt-2"
        @click="isCheck = false"
      >
        收起
      </n-button>
    </div>
    <n-data-table
      :columns="columns"
      :data="data"
    />
  </div>
</template>

<style scoped lang="scss">

</style>
