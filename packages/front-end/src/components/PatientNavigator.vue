<script setup lang="ts">
import Navigator from "@components/Navigator.vue";

const NavList = [{
  "name": "我的",
  "path": "/patient"
},
{
  "name": "呼叫",
  "path": "/call"
}
];

</script>
<template>
  <Navigator :nav-list="NavList" />
</template>

<style scoped lang="scss">

</style>
