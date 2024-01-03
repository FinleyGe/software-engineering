import { defineStore } from "pinia";
import { ref } from "vue";
import type { Role } from "@/type";

export const useUserStore = defineStore(
  "user",
  () => {
    const isLogin = ref<boolean>(false);
    const role = ref<Role>("guest");
    const id = ref<number>(0);
    function logout() {
      isLogin.value = false;
      role.value = "guest";
    }
    return { isLogin, role, logout, id};
  },
);
