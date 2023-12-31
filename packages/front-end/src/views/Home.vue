<script setup lang="ts">
import { useUserStore } from "@/store/user";
const userStore = useUserStore();
</script>
<template>
  <h1 class="text-2xl m-2">
    <template v-if="userStore.isLogin">
      欢迎使用医院管理系统
    </template>
    <template v-else>
      请先登录
    </template>
  </h1>
</template>

<style scoped lang="scss"></style>
