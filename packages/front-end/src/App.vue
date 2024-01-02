<script setup lang="ts">
import AdminNavigator from "@components/AdminNavigator.vue";
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
  }
});

</script>

<template>
  <n-config-provider>
    <n-message-provider>
      <Login
        v-show="isLoginShow"
        @close="isLoginShow = false"
      />
      <header
        class="drop-shadow-xl h-fit bg-gray-100 p-3"
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
          <template v-else-if="userStore.role == 'admin'">
            <AdminNavigator />
          </template>
        </n-space>
      </header>
      <main class="bg-gray-50 m-3 p-3">
        <RouterView />
      </main>
    </n-message-provider>
  </n-config-provider>
</template>

<style scoped></style>
