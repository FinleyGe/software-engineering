<script setup lang="ts">
import Navigator from "@components/Navigator.vue";

const NavList = [{
  "name": "患者管理",
  "path": "/doctor/patient"
},
{
  "name": "病例管理",
  "path": "/doctor/log"
},{
  "name": "我的",
  "path": "/doctor"
}];

</script>
<template>
  <Navigator :nav-list="NavList" />
</template>

<style scoped lang="scss">

</style>
