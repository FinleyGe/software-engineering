import { createRouter, createWebHistory } from "vue-router";
import type { RouteRecordRaw } from "vue-router";

const routes = <RouteRecordRaw[]>[{
  path: "/",
  name: "Home",
  component: () => import("@views/Home.vue"),
  meta:{
    title:"医院管理系统 | 首页"
  }
},
{
  path: "/call",
  name: "Call",
  component: () => import("@views/Call.vue"),
  meta:{
    title:"医院管理系统 | 呼叫"
  },
},{
  path: "/patient",
  name: "Patient",
  component: () => import("@views/PatientHome.vue"),
  meta:{
    title:"医院管理系统 | 患者详情"
  }
},{
  path: "/doctor",
  name: "Doctor",
  children: [
    {
      path: "patient",
      name: "DoctorPatient",
      component: () => import("@views/doctor/Patient.vue"),
      meta: {
        title:"医院管理系统 | 患者管理"
      }
    },
    {
      path: "patient/:id",
      name: "DoctorPatientDetail",
      component: () => import("@views/PatientDetail.vue"),
      props: true,
      meta: {
        title: "医院管理系统 | 患者详情"
      }
    }, {
      path: "log",
      name: "DoctorLog",
      component: () => import("@views/doctor/Log.vue"),
      meta: {
        title: "医院管理系统 | 日志管理"
      }
    }
  ]
},
{
  path: "/monitor/:id",
  name: "AdminVitalSign",
  component: () => import("@views/Monitor.vue"),
  props: true,
  meta: {
    title: "医院管理系统 | 生命体征监测"
  }
},
{
  path: "/admin",
  children: [
    {
      path: "patient",
      name: "AdminPatient",
      component: () => import("@views/manage/Patient.vue"),
      meta:{
        title:"医院管理系统 | 患者管理"
      }
    },
    {
      path: "doctor",
      name: "AdminDoctor",
      component: () => import("@views/manage/Doctor.vue"),
      meta: {
        title: "医院管理系统 | 医生管理"
      }
    },
    {
      path: "department",
      name: "AdminDepartment",
      component: () => import("@views/manage/Department.vue"),
      meta: {
        title: "医院管理系统 | 科室管理"
      }
    },
    {
      path: "bed",
      name: "AdminBed",
      component: () => import("@views/manage/Bed.vue"),
      meta: {
        title: "医院管理系统 | 床位管理"
      }
    },
    {
      path: "patient/:id",
      name: "AdminPatientDetail",
      component: () => import("@views/PatientDetail.vue"),
      props: true,
      meta: {
        title: "医院管理系统 | 患者详情"
      }
    },
  ]
}
];

const router =  createRouter({
  history: createWebHistory(),
  routes
});

router.beforeEach((to) => {
  document.title = to.meta.title as string;
}
);
export default router;
