<script setup lang="ts">
import PatientDetail from "@views/PatientDetail.vue";
import { useUserStore } from "@/store";

const userStore = useUserStore();
</script>
<template>
  {{ userStore }}
  <PatientDetail :id="userStore.id" />
</template>

<style scoped lang="scss">

</style>
