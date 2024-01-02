<script setup lang="ts">
import Navigator from "@components/Navigator.vue";

const NavList = [{
  "name": "患者管理",
  "path": "/admin/patient"
}, {
  "name": "医生管理",
  "path": "/admin/doctor"
}, {
  "name": "科室管理",
  "path": "/admin/department"
}, {
  "name": "床位管理",
  "path": "/admin/bed"
}] ;
</script>
<template>
  <Navigator :nav-list="NavList" />
</template>

<style scoped lang="scss">

</style>
