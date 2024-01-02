import { createRouter, createWebHistory } from "vue-router";
import type { RouteRecordRaw } from "vue-router";

const routes = <RouteRecordRaw[]>[
  {
    path: "/",
    name: "Home",
    component: () => import("@views/Home.vue"),
    meta:{
      title:"医院管理系统 | 首页"
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
      }
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
