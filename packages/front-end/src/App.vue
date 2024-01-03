<script setup lang="ts">
import AdminNavigator from "@components/AdminNavigator.vue";
import DoctorNavigator from "@components/DoctorNavigator.vue";
import PatientNavigator from "@components/PatientNavigator.vue";
import Login from "@components/Login.vue";
import { useUserStore } from "@store/user";
import { useVueCookies } from "@/cookie";
import { NButton, NConfigProvider, NMessageProvider, NSpace } from "naive-ui";
import { RouterView } from "vue-router";
import { onMounted, ref } from "vue";
import type { Role } from "@/type";

const isLoginShow = ref<boolean>(false);
const userStore = useUserStore();
const cookies = useVueCookies();

onMounted(() => {
  if (cookies.get("token")) {
    if (!localStorage.getItem("role")) {
      return;
    } else {
      userStore.isLogin = true;
      userStore.role = localStorage.getItem("role") as Role;
    }
    if (!localStorage.getItem("id")) {
      userStore.id = localStorage.getItem("id") as string;
    }
  }
});

</script>

<template>
  <n-config-provider
    style="min-height: 100vh;"
    class="
    from-cyan-100
    to-blue-100
    bg-gradient-to-l
    hover:bg-gradient-to-r"
  >
    <n-message-provider>
      <Login
        v-show="isLoginShow"
        @close="isLoginShow = false"
      />
      <header
        class="drop-shadow-xl h-fit bg-gray-100 p-3
        bg-gradient-to-r from-cyan-500 to-blue-50
        hover:bg-gradient-to-l
        ease-in-out
        transition
        "
      >
        <n-space
          justify="space-between"
          align="center"
          style="margin: 0 1rem;"
        >
          <router-link
            to="/"
            class="text-3xl m-3 font-bold text-gray-800"
            style="text-decoration: none;"
          >
            医院管理系统
          </router-link>
          <template v-if="userStore.role == 'guest'">
            <n-button
              v-if="!userStore.isLogin"
              type="primary"
              @click="isLoginShow = true"
            >
              登录
            </n-button>
          </template>
          <AdminNavigator v-else-if="userStore.role==='admin'" />
          <DoctorNavigator v-else-if="userStore.role==='doctor'" />
          <PatientNavigator v-else-if="userStore.role==='patient'" />
        </n-space>
      </header>
      <main class="p-10">
        <RouterView />
      </main>
    </n-message-provider>
  </n-config-provider>
</template>

<style scoped lang="scss"></style>
