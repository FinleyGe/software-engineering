<script setup lang="ts">
import {AdminAddDoctor, AdminDeleteDoctor,AdminGetAllDoctors,AdminGetDepartments, AdminUpdateDoctor } from "@/api/admin";
import type {DataTableColumns, SelectOption} from "naive-ui";
import { Department, Doctor } from "@/type";
import {NButton, NDataTable, NDivider, NInput, NSelect, NSpace, useMessage} from "naive-ui";
import {h, onMounted, ref} from "vue";
import SelectableCell from "@components/SelectableCell.vue";

const pagetitle = "医生管理";
const doctors = ref<Doctor[]>([]);
const doctorName = ref<string>("");
const doctorDepartment = ref<Department>({id: 0, name: ""});
const message = useMessage();
const DepartmentList = ref<SelectOption[]>([]);

const cellEdit = ref<Array<boolean>>([]);
const isAdd = ref<boolean>(false);

function handleRefresh  ()  {
  AdminGetAllDoctors().then((res) => {
    doctors.value = res.data.data;
  });
  cellEdit.value = new Array(doctors.value.length);
  cellEdit.value.map(() => {false;});

  AdminGetDepartments().then((res) => {
    DepartmentList.value = res.data.data.map((item: Department) => {
      return {
        label: `${item.name} (${item.id})`,
        value: item.id,
      };
    });
  });
};

function handleDelete(index: number) {
  AdminDeleteDoctor(doctors.value[index].id).then((res) => {
    if (res.data.msg === "OK") {
      handleRefresh();
      message.success("删除成功");
    } else {
      message.error("删除失败");
    }
  });
}

function handleAddDoctor  () {
  if (doctorName.value === "") {
    message.error("医生名称不能为空");
    return;
  }
  AdminAddDoctor(doctorName.value, doctorDepartment.value.id).then((res) => {
    if (res.data.msg === "OK") {
      handleRefresh();
      message.success("添加成功");
    } else {
      message.error("添加失败");
    }
  });
};

function handleEdit(index: number, department_id: number) {
  if (department_id === doctors.value[index].department_id)  {
    message.info("未做修改");
    return;
  }
  AdminUpdateDoctor(
    doctors.value[index].id,
    department_id,
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

const columns = <DataTableColumns<Doctor>>[
  {
    title: "医生编号",
    key: "id",
  }, {
    title: "医生名称",
    key: "name",
    // render: (_, index: number) => {
    //   return h(
    //     EditableCell,
    //     {
    //       value: doctors.value[index].name,
    //       onSubmit: (v:string)=>{handleEdit(index, v, doctors.value[index].department_id);},
    //     }
    //   );
    // },
  },{
    title: "所属科室",
    key: "department",
    render: (_, index: number) => {
      return h(
        SelectableCell,
        {
          value: doctors.value[index].department_id,
          options:
          DepartmentList.value.map((item: SelectOption) => {
            return {
              label: item.label,
              value: item.value as number,
            };
          }),
          onSubmit: (v:number)=>{handleEdit(index, v);},
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

const data = doctors;

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
          <n-button
            type="info"
            @click="isAdd = !isAdd"
          >
            添加医生
          </n-button>
        </n-space>
      </n-space>
    </header>
    <main>
      <n-space v-if="isAdd">
        <n-space vertical>
          <n-input
            v-model:value="doctorName"
            placeholder="医生名称"
          />
          <n-select
            v-model:value="doctorDepartment.id"
            :options="DepartmentList"
            placeholder="所属科室"
          />
        </n-space>
        <n-button
          type="success"
          @click="handleAddDoctor"
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
