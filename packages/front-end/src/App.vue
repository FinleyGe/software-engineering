<script setup lang="ts">
import { NButton, NMessageProvider, NSpace } from "naive-ui";
import { onMounted, ref } from "vue";
import AdminNavigator from "@components/AdminNavigator.vue";
import Login from "@components/Login.vue";
import { RouterView } from "vue-router";
import { useUserStore } from "@store/user";
import { useVueCookies } from "@/cookie";

const isLoginShow = ref<boolean>(false);
const userStore = useUserStore();
const cookies = useVueCookies();

onMounted(() => {
  if (cookies.get("token")) {
    if (!localStorage.getItem("role")) {
      return;
    } else {
      userStore.isLogin = true;
      userStore.role = localStorage.getItem("role") as string;
    }
  }
});

</script>

<template>
  <n-message-provider>
    <Login
      v-show="isLoginShow"
      @close="isLoginShow = false"
    />
    <div class="h-full bg-white dark:bg-slate-800">
      <n-space
        justify="space-between"
        align="center"
        style="margin: 0 1rem;"
      >
        <div class="text-3xl m-3 font-bold cursor-pointer">
          <router-link to="/">
            医院管理系统
          </router-link>
        </div>
        <template v-if="userStore.role == 'guest'">
          <n-button
            v-if="!userStore.isLogin"
            type="primary"
            class="bg-green-500 hover:bg-green-600"
            @click="isLoginShow = true"
          >
            登录
          </n-button>
        </template>
        <template v-else-if="userStore.role == 'admin'">
          <AdminNavigator />
        </template>
      </n-space>
      <main>
        <RouterView />
      </main>

      <footer class="text-xs italic">
        Software Engineering Project by ZJUTers
      </footer>
    </div>
  </n-message-provider>
</template>

<style scoped></style>
