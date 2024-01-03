<script setup lang="ts">
import {PatientCall} from "@/api/patient";
import {NButton,useMessage} from "naive-ui";

const message = useMessage();

function handleCall(){
  PatientCall().then(res => {
    if(res.data.msg === "OK"){
      message.success("呼叫成功");
    }
  }).catch(err => {
    message.error("呼叫失败");
  });
  // console.log("call");
}
</script>
<template>
  <div class="flex h-full justify-center">
    <n-button
      type="error"
      class="w-32 h-32 rounded-full"
      @click="handleCall"
    >
      Call
    </n-button>
  </div>
</template>

<style scoped lang="scss">

</style>
