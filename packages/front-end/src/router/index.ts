import { createRouter, createWebHistory } from "vue-router";
import type { RouteRecordRaw } from "vue-router";

const routes = <RouteRecordRaw[]>[
  {
    path: "/",
    name: "Home",
    component: () => import("@views/Home.vue")
  }
];

export default createRouter({
  history: createWebHistory(),
  routes
});
