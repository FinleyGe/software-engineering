<script setup lang="ts">
import {NButton, NSpace} from "naive-ui";

const props = defineProps<{
  navList: {
    name: string;
    path: string;
  }[];
}>();

</script>
<template>
  <n-space>
    <n-button
      v-for="item in props.navList"
      :key="item.name"
      type="primary"
      @click="$router.push(item.path)"
    >
      <!-- <router-link -->
      <!--   :to="item.path" -->
      <!--   style="text-decoration: none;" -->
      <!-- > -->
      {{ item.name }}
      <!-- </router-link> -->
    </n-button>
  </n-space>
</template>

<style scoped lang="scss">

</style>
