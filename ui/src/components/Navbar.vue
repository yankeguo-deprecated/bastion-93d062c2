<template>
  <el-menu theme="dark" mode="horizontal">
    <el-menu-item index="1"><router-link :to="{ name: 'dashboard' }">Bastion</router-link></el-menu-item>
    <el-menu-item index="2"><router-link :to="{ name: 'profile' }">当前用户：{{ currentUser.nickname }}</router-link></el-menu-item>
  </el-menu>
</template>

<script>
export default {
  name: 'navbar',
  computed: {
    currentUser () {
      return this.$store.state.user.currentUser || { nickname: '正在加载' }
    }
  },
  data () {
    return {}
  }
}
</script>
