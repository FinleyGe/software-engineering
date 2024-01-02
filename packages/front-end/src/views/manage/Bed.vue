<script setup lang="ts">
import { AdminAddBed, AdminDeleteBed,AdminGetBeds,  } from "@/api/admin";
import {NButton, NDataTable, NDivider, NInput, NSpace, useMessage} from "naive-ui";
import {h, onMounted, ref} from "vue";
import type { Bed } from "@/type";
import type {DataTableColumns} from "naive-ui";

const pagetitle = "床位管理";
const beds = ref<Bed[]>([]);
const bedName = ref<string>("");
const message = useMessage();

const cellEdit = ref<Array<boolean>>([]);

function handleRefresh  ()  {
  AdminGetBeds().then((res) => {
    beds.value = res.data.data;
  });
  cellEdit.value = new Array(beds.value.length);
  cellEdit.value.map(() => {false;});
};

function handleDelete(index: number) {
  AdminDeleteBed(beds.value[index].id).then((res) => {
    if (res.data.msg === "OK") {
      handleRefresh();
      message.success("删除成功");
    } else {
      message.error("删除失败");
    }
  });
}

const handleAddBed = () => {
  if (bedName.value === "") {
    message.error("床位名称不能为空");
    return;
  }
  AdminAddBed(bedName.value).then((res) => {
    if (res.data.msg === "OK") {
      handleRefresh();
      message.success("添加成功");
    } else {
      message.error("添加失败");
    }
  });
};

const columns = <DataTableColumns<Bed>>[
  {
    title: "床位ID",
    key: "id",
  },{
    title: "床位号码",
    key: "bed_number",
  },{
    title: "是否被占用",
    key: "occupied",
    render(row) {
      return h(
        "span",
        {
          style: {
            color: row.occupied ? "red" : "green",
          },
        },
        {
          default: () => (row.occupied ? "是" : "否"),
        }
      );
    },
  },
  {
    title: "操作",
    key: "action",
    render(_, index) {
      return h(
        NButton,
        {
          type: "error",
          onClick: () => {
            handleDelete(index);
          },
        },
        {
          default: () => "删除",
        }
      );
    },
  }
];

onMounted(() => {
  handleRefresh();
});

const data = beds;

</script>
<template>
  <div>
    <header class="flex flex-row items-center">
      <n-space
        justify="center"
        align="center"
      >
        <h1 class="text-xl">
          {{ pagetitle }}
        </h1>
        <n-space>
          <n-button
            type="primary"
            @click="handleRefresh"
          >
            刷新
          </n-button>
        </n-space>
      </n-space>
    </header>
    <main>
      <n-space>
        <n-input
          v-model:value="bedName"
          placeholder="床位名称"
        />
        <n-button
          type="primary"
          @click="handleAddBed"
        >
          添加
        </n-button>
      </n-space>
      <n-divider />
      <n-data-table
        :columns="columns"
        :data="data"
        style="max-height: 80svh"
      />
    </main>
  </div>
</template>

<style scoped lang="scss">

</style>
