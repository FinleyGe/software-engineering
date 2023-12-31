import { defineStore } from "pinia";
import { ref } from "vue";
import type { Role } from "@/type";
export const useUserStore = defineStore(
  "user",
  () => {
    const isLogin = ref<boolean>(false);
    const role = ref<Role>("guest");
    return { isLogin, role };
  },
);
