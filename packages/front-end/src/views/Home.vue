<script setup lang="ts">
import { useUserStore } from "@/store/user";
import { useVueCookies } from "@/cookie";
import {useRouter} from "vue-router";
import {NButton} from "naive-ui";

const router = useRouter();
const userStore = useUserStore();
const cookie = useVueCookies();

function handleLogout() {
  cookie.remove("token");
  userStore.isLogin = false;
  userStore.role = "guest";
  router.push({
    path: "/"
  });
}

</script>
<template>
  <h1 class="text-2xl m-2">
    <template v-if="userStore.isLogin">
      <div class="flex flex-col">
        <h1 class="text-2xl flex justify-center">
          欢迎使用
        </h1>
        <n-button
          type="error"
          class="w-fit mx-auto"
          @click="handleLogout"
        >
          退出登录
        </n-button>
      </div>
    </template>
    <template v-else>
      <!-- <router-link>登录</router-link> -->
      <div
        class="
        flex justify-center items-center"
        style="min-height: calc(100vh - 130px);"
        bg-transparent
      >
        <h1 class="text-3xl flex justify-center">
          欢迎使用, 请登录
        </h1>
      </div>
    </template>
  </h1>
</template>

<style scoped lang="scss"></style>
