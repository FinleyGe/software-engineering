<script setup lang="ts">
import { AdminAddDepartment, AdminDeleteDepartment,AdminGetDepartments, AdminUpdateDepartment } from "@/api/admin";
import {NButton, NDataTable, NDivider, NInput, NSpace, useMessage} from "naive-ui";
import {h, onMounted, ref} from "vue";
import type {DataTableColumns} from "naive-ui";
import { Department } from "@/type";
import EditableCell from "@components/EditableCell.vue";

const pagetitle = "科室管理";
const departments = ref<Department[]>([]);
const departmentName = ref<string>("");
const message = useMessage();

const cellEdit = ref<Array<boolean>>([]);

function handleRefresh  ()  {
  AdminGetDepartments().then((res) => {
    departments.value = res.data.data;
  });
  cellEdit.value = new Array(departments.value.length);
  cellEdit.value.map(() => {false;});
};

function handleDelete(index: number) {
  AdminDeleteDepartment(departments.value[index].id).then((res) => {
    if (res.data.msg === "OK") {
      handleRefresh();
      message.success("删除成功");
    } else {
      message.error("删除失败");
    }
  });
}

const handleAddDepartment = () => {
  if (departmentName.value === "") {
    message.error("科室名称不能为空");
    return;
  }
  AdminAddDepartment(departmentName.value).then((res) => {
    if (res.data.msg === "OK") {
      handleRefresh();
      message.success("添加成功");
    } else {
      message.error("添加失败");
    }
  });
};

function handleEdit(index: number, value: string){
  if (value === departments.value[index].name)  {
    message.info("未做修改");
    return;
  }
  AdminUpdateDepartment(
    departments.value[index].id,
    value,
  ).then(
    (res) => {
      if(res.data.msg === "OK"){
        message.success("修改成功");
        handleRefresh();
      }
    }
  ).catch(
    (err) => {
      message.error(err);
    }
  );
}
const columns = <DataTableColumns<Department>>[
  {
    title: "科室编号",
    key: "id",
  }, {
    title: "科室名称",
    key: "name",
    render: (_, index: number) => {
      return h(
        EditableCell,
        {
          value: departments.value[index].name,
          onSubmit: (v:string)=>{handleEdit(index, v);},
        }
      );
    },
  }, {
    title:"删除",
    key: "delete",
    render: (_, index: number) => {
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

const data = departments;

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
          v-model:value="departmentName"
          placeholder="科室名称"
        />
        <n-button
          type="primary"
          @click="handleAddDepartment"
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
