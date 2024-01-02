<script setup lang="ts">
import {NButton, NInput,NSpace} from "naive-ui";
import { Login } from "@/api";
import { Role } from "@/type/user";
import { ref } from "vue";
import {useMessage} from "naive-ui";
import { useUserStore } from "@/store/user";
const message = useMessage();
const userStore = useUserStore();
const emits = defineEmits<{
  (e:"close"):void,
}>();

const id = ref<string>("");
const password = ref<string>("");
const role = ref<Role>("doctor");

function handleLogin(){
  console.log("login");
  Login(
    Number(id.value),
    password.value,
    role.value
  ).then((res) => {
    if (res.data.msg === "OK") {
      message.success("登录成功");
      userStore.isLogin = true;
      userStore.role = role.value;
      localStorage.setItem("role", role.value);
    }
    emits("close");
  }).catch((err) => {
    message.error(err);
  });
  emits("close");
}

</script>
<template>
  <div
    class="fixed inset-0 z-40 flex items-center justify-center bg-gray-900 bg-opacity-50 backdrop-blur"
  >
    <div class="flex flex-col bg-white rounded-lg shadow-lg p-10 z-50">
      <header>
        <h1 class="text-xl font-bold w-fit m-auto">
          登录
        </h1>
      </header>

      <n-space vertical>
        <n-space vertical>
          <n-input
            v-model:value="id"
            type="text"
            placeholder="请输入帐号"
          />
          <n-input
            v-model:value="password"
            type="password"
            placeholder="请输入密码"
          />
        </n-space>
        <div class="flex flex-row items-center">
          <label class="text-sm font-bold">
            身份
          </label>
          <input
            v-model="role"
            type="radio"
            name="identity"
            value="doctor"
            class="border border-gray-300 rounded-lg p-2"
          >
          <label class="text-sm font-bold">
            医生
          </label>
          <input
            v-model="role"
            type="radio"
            name="identity"
            value="patient"
            class="border border-gray-300 rounded-lg p-2"
          >
          <label class="text-sm font-bold">
            患者
          </label>
          <input
            v-model="role"
            type="radio"
            name="identity"
            value="admin"
            class="border border-gray-300 rounded-lg p-2"
          >
          <label class="text-sm font-bold">
            管理员
          </label>
        </div>
      </n-space>
      <footer class="mt-2 flex flex-row justify-around">
        <n-button
          type="primary"
          @click="handleLogin"
        >
          登录
        </n-button>
        <n-button
          type="error"
          @click="$emit('close')"
        >
          取消
        </n-button>
      </footer>
    </div>
  </div>
</template>

<style scoped lang="scss"></style>
