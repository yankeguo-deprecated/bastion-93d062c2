<template>
  <el-menu theme="dark" mode="horizontal">
    <el-menu-item index="1"><router-link :to="{ name: 'dashboard' }">Bastion</router-link></el-menu-item>
  </el-menu>
</template>

<script>
export default {
  name: 'navbar',
  data () {
    return {}
  }
}
</script>
