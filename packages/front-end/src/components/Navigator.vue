<script setup lang="ts">

const props = defineProps<{
  navList: {
    name: string;
    path: string;
  }[];
}>();

</script>
<template>
  <div class="flex flex-row gap-3">
    <button
      v-for="item in props.navList"
      :key="item.name"
      class="text-l cursor-pointer bg-blue-100 hover:bg-blue-200 p-2 rounded-xl"
    >
      <router-link :to="item.path">
        {{ item.name }}
      </router-link>
    </button>
  </div>
</template>

<style scoped lang="scss">

</style>
