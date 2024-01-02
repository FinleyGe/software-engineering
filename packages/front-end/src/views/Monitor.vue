<script setup lang="ts">
import { useUserStore } from "@/store";
import { AdminGetVitalSigns } from "@/api/admin";
import { dateFormat } from "@/util/date";
import { useMessage, NButton, NSpace, NDatePicker, NDataTable } from "naive-ui";
// import { useRouter } from "vue-router";
import { ref, onBeforeMount, computed } from "vue";
import  VChart  from "vue-echarts";
import { use } from "echarts/core";
import { LinesChart,LineChart } from "echarts/charts";
import { GridComponent ,
  TitleComponent,
  TooltipComponent,
  LegendComponent
} from "echarts/components";
import { CanvasRenderer } from "echarts/renderers";
import type { VitalSign } from "@/type";
import type { DataTableColumns } from "naive-ui";
import type { ComposeOption } from "echarts/core";
import type { LineSeriesOption } from "echarts/charts";
import type {
  TitleComponentOption,
  TooltipComponentOption,
  LegendComponentOption
} from "echarts/components";

use([
  TitleComponent,
  TooltipComponent,
  LegendComponent,
  // LinesChart,
  LineChart,
  CanvasRenderer,
  GridComponent
]);

type EChartsOption = ComposeOption<
  | TitleComponentOption
  | TooltipComponentOption
  | LegendComponentOption
  | LineSeriesOption
>;

const message = useMessage();
// const router = useRouter();

const props = defineProps<{
  id: number; // bed_id
}>();

const timeRange = ref<[number, number]>([
  Date.now() - 1000 * 60 * 60 * 24,
  Date.now(),
]);

const userStore = useUserStore();

const vitalsigns = ref<VitalSign[]>();

const option = ref<EChartsOption>({
  title: {
    text: "监测数据",
    left: "center"
  },
  tooltip: {
    trigger: "axis"
  },
  legend: {
    data:[
      "体温",
      "心率",
      "呼吸频率",
      "血压高压",
      "血压低压",
      "血氧"
    ],
    top: "bottom"
  },
  xAxis:{
    type: "category",
    boundaryGap: false,
    data: computed(
      ()=>{return vitalsigns.value?.map((item) => item.time);})
  },
  yAxis: {
    type: "value",
  },
  // series: [
  //   {
  //     data: vitalsigns.value?.map((item) => item.temperature),
  //     name: "体温",
  //     type: "line"
  //   },{
  //     data: vitalsigns.value?.map((item) => item.heart_rate),
  //     name: "心率",
  //     type: "line"
  //   },{
  //     data: vitalsigns.value?.map((item) => item.breathing_rate),
  //     name: "呼吸频率",
  //     type: "line"
  //   },{
  //     data: vitalsigns.value?.map((item) => item.blood_pressure.split("/")[0]),
  //     name: "血压高压",
  //     type: "line"
  //   },{
  //     data: vitalsigns.value?.map((item) => item.blood_pressure.split("/")[1]),
  //     name: "血压低压",
  //     type: "line"
  //   },{
  //     data: vitalsigns.value?.map((item) => item.blood_oxygen),
  //     name: "血氧",
  //     type: "line"
  //   }]
});

function setSeriesData() {
  console.log(vitalsigns.value);
  option.value.series =
    [{
      data: vitalsigns.value?.map((item) => item.temperature),
      name: "体温",
      type: "line"
    },{
      data: vitalsigns.value?.map((item) => item.heart_rate),
      name: "心率",
      type: "line"
    },{
      data: vitalsigns.value?.map((item) => item.breathing_rate),
      name: "呼吸频率",
      type: "line"
    },{
      data: vitalsigns.value?.map((item) => item.blood_pressure.split("/")[0]),
      name: "血压高压",
      type: "line"
    },{
      data: vitalsigns.value?.map((item) => item.blood_pressure.split("/")[1]),
      name: "血压低压",
      type: "line"
    },{
      data: vitalsigns.value?.map((item) => item.blood_oxygen),
      name: "血氧",
      type: "line"
    }];
}

function handleRefresh() {
  switch (userStore.role) {
  case "admin":
    AdminGetVitalSigns(
      props.id,
      dateFormat(new Date(timeRange.value[0])),
      dateFormat(new Date(timeRange.value[1]))
    ).then((res) => {
      if (res.data.data.length === 0) {
        message.error("暂无数据");
        return;
      }
      vitalsigns.value = res.data.data;
      setSeriesData();
    });
    break;
  case "doctor":
    break;
  case "patient":
    // router.push({ path: "/patient"
    // });
    break;
  }
}

const columns = <DataTableColumns>[
  {
    title: "时间",
    key: "time",
  }, {
    title: "体温",
    key: "temperature",
  }, {
    title: "心率",
    key: "heart_rate",
  }, {
    title: "呼吸频率",
    key: "breathing_rate",
  }, {
    title: "血压",
    key: "blood_pressure",
  }, {
    title: "血氧",
    key: "blood_oxygen",
  }
];

onBeforeMount(() => {
  handleRefresh();
});

</script>
<template>
  {{ option }}
  <div>
    <n-space>
      <n-button
        type="primary"
        @click="handleRefresh"
      >
        刷新
      </n-button>
    </n-space>
    <n-date-picker
      v-model:value="timeRange"
      type="datetimerange"
      placeholder="选择结束时间"
    />
    <n-data-table
      :columns="columns"
      :data="vitalsigns"
      style="max-height: 100svh"
    />
  </div>
  <VChart
    v-if="vitalsigns"
    :option="option"
    class="vchart"
  />
</template>

<style scoped lang="scss">
.vchart {
  height: 500px;
}
</style>
