<script setup lang="ts">
import {
  DoctorGetPatients,
  DoctorAddPatient,
  DoctorGetBeds,
  DoctorDeletePatient
} from "@/api/doctor";
import {AddPatientRequest} from "@/type/request";
import { Bed, Patient } from "@/type";
import { useUserStore } from "@/store";
// import EditableCell from "@components/EditableCell.vue";
import { dateFormat } from "@/util/date";
import {
  NButton,
  NDataTable,
  NDivider,
  NInput,
  NSelect,
  NSpace,
  NDatePicker,
  useMessage,
} from "naive-ui";
import {h, onMounted, reactive, ref, watch} from "vue";
import { useRouter } from "vue-router";
import type {DataTableColumns, SelectOption} from "naive-ui";

const pagetitle = "患者管理";
const patients = ref<Patient[]>([]);
const message = useMessage();
const cellEdit = ref<Array<boolean>>([]);
const isAdd = ref<boolean>(false);
const userStore = useUserStore();
const birth = ref<number>(0);
const beds = ref<Bed[]>([]);

const patient = reactive<AddPatientRequest>({
  name: "",
  state: "",
  birth: "",
  bed_id: beds.value[0]?.id,
  phone: "",
  gender: "男",
});

watch(birth, (newVal) => {
  patient.birth = dateFormat(new Date(newVal));
});

const genderOptions = ref<SelectOption[]>([
  { label: "男", value: "男" },
  { label: "女", value: "女" },
]);

const bedOptions = ref<SelectOption[]>([]);
function handleLog() {
  router.push({
    path: "/doctor/log",
  });
}
function handleRefresh (){
  if (userStore.role === "doctor") {
    DoctorGetPatients().then((res) => {
      patients.value = res.data.data.filter((patient) => {
        return patient.state !== "out";
      });
    });
    cellEdit.value = new Array(patients.value.length);
    cellEdit.value.map(() => {false;});
    DoctorGetBeds().then((res) => {
      beds.value = res.data.data;
      bedOptions.value = beds.value.filter((bed) => {
        return bed.occupied === false;
      }).
        map((bed) => {
          return {
            label: bed.bed_number,
            value: bed.id,
          };
        });
    });
    return;
  } else {
    message.error("您没有权限");
  }
};

function handleAddPatient  () {
  DoctorAddPatient(patient).then((res) => {
    if (res.data.msg === "OK") {
      message.success("添加成功");
      handleRefresh();
    } else {
      message.error("添加失败");
    }
  }).catch((err) => {
    message.error("添加失败" + err);
  });
  handleRefresh();
};

const router = useRouter();

function handleCheck(index: number) {
  router.push({
    path: "/doctor/patient/" + patients.value[index].id,
  });
};

function handleDelete(index: number) {
  DoctorDeletePatient(patients.value[index].id).then((res) => {
    if (res.data.msg === "OK") {
      message.success("出院成功");
      handleRefresh();
    } else {
      message.error("请求失败");
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
  }, {
    title: "床号",
    key: "bed.number",
  } , {
    title: "状态",
    key: "state",
  }
  , {
    title:"操作",
    key: "action",
    render: (_, index: number) => {
      return h(
        NSpace, {}, {
          default: () => [
            h(
              NButton, {
                type: "error",
                onClick: () => {
                  handleDelete(index);
                },
              }, {
                default: () => "出院",
              }
            ),
            h(
              NButton, {
                type: "info",
                onClick: () => {
                  handleCheck(index);
                },
              }, {
                default: () => "查看详细",
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
          <n-button
            type="info"
            @click="handleLog"
          >
            写病例
          </n-button>
        </n-space>
      </n-space>
    </header>
    <main>
      <n-space v-if="isAdd">
        <n-space vertical>
          <n-input
            v-model:value="patient.name"
            placeholder="患者名称"
          />
          <n-input
            v-model:value="patient.phone"
            placeholder="患者电话"
          />
          <n-date-picker
            v-model:value="birth"
            type="date"
            placeholder="患者生日"
          />
          <n-select
            v-model:value="patient.gender"
            :options="genderOptions"
            placeholder="患者性别"
          />
          <n-input
            v-model:value="patient.state"
            placeholder="患者状态"
          />
          <n-select
            v-model:value="patient.bed_id"
            :options="bedOptions"
            placeholder="患者床号"
          />
          <n-button
            type="success"
            @click="handleAddPatient"
          >
            添加
          </n-button>
        </n-space>
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
