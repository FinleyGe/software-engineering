<script setup lang="ts">
import {
  AdminGetAllPatients,
  AdminDeletePatient,
} from "@/api/admin";
import { Department, Patient } from "@/type";
import {
  NButton,
  NDataTable,
  NDivider,
  NInput,
  NSelect,
  NSpace,
  useMessage,
  NButtonGroup,
} from "naive-ui";
import {h, onMounted, ref} from "vue";
import { useRouter } from "vue-router";
import type {DataTableColumns, SelectOption} from "naive-ui";

const pagetitle = "患者管理";
const patients = ref<Patient[]>([]);
const patientName = ref<string>("");
const patientDepartment = ref<Department>({id: 0, name: ""});
const message = useMessage();
const DepartmentList = ref<SelectOption[]>([]);

const cellEdit = ref<Array<boolean>>([]);
const isAdd = ref<boolean>(false);

function handleRefresh  ()  {
  AdminGetAllPatients().then((res) => {
    patients.value = res.data.data;
  });
  cellEdit.value = new Array(patients.value.length);
  cellEdit.value.map(() => {false;});
};

function handleAddPatient  () {
  if (patientName.value === "") {
    message.error("患者名称不能为空");
    return;
  }
  // AdminAddPatient(patientName.value, doctorDepartment.value.id).then((res) => {
  //   if (res.data.msg === "OK") {
  //     handleRefresh();
  //     message.success("添加成功");
  //   } else {
  //     message.error("添加失败");
  //   }
  // });
};
const router = useRouter();
function handleCheck(index: number) {
  router.push({
    path: "/admin/patient/" + patients.value[index].id,
  });
};

// function handleEdit(index: number, department_id: number) {
//   if (department_id === patients.value[index].department_id)  {
//     message.info("未做修改");
//     return;
//   }
//   AdminUpdatePatient(
//     patients.value[index].id,
//     department_id,
//   ).then(
//     (res) => {
//       if(res.data.msg === "OK"){
//         message.success("修改成功");
//         handleRefresh();
//       }
//     }
//   ).catch(
//     (err) => {
//       message.error(err);
//     }
//   );
// }

function handleDelete(index: number) {
  AdminDeletePatient(patients.value[index].id).then((res) => {
    if (res.data.msg === "OK") {
      message.success("删除成功");
      handleRefresh();
    } else {
      message.error("删除失败");
    }
  });
}

const columns = <DataTableColumns<Patient>>[
  {
    title: "患者编号",
    key: "id",
  } , {
    title: "患者姓名",
    key: "name",
  } , {
    title: "医生名称",
    key: "doctor.name",
  } , {
    title: "床号",
    key: "bed.number",
  } , {
    title: "状态",
    key: "state",
  } , {
    title:"操作",
    key: "action",
    render: (_, index: number) => {
      return h(
        // NButton, {
        //   type: "error",
        //   onClick: () => {
        //     handleDelete(index);
        //   },
        // } , {
        //   default: () => "删除",
        //}
        NButtonGroup, {
          size: "small",
        }, {
          default: () => [
            h(
              NButton, {
                type: "error",
                onClick: () => {
                  handleDelete(index);
                },
              }, {
                default: () => "删除",
              }
            ),
            h(
              NButton, {
                type: "warning",
                onClick: () => {
                  handleCheck(index);
                },
              }, {
                default: () => "查看",
              }
            ),
          ],
        }
      );
    },
  }
];

onMounted(() => {
  handleRefresh();
});

const data = patients;

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
            添加患者
          </n-button>
        </n-space>
      </n-space>
    </header>
    <main>
      <n-space v-if="isAdd">
        <n-space vertical>
          <n-input
            v-model:value="patientName"
            placeholder="患者名称"
          />
          <n-select
            v-model:value="patientDepartment.id"
            :options="DepartmentList"
            placeholder="所属科室"
          />
        </n-space>
        <n-button
          type="success"
          @click="handleAddPatient"
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
