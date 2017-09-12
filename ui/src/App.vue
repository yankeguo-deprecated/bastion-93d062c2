<template>
  <div id="app">
    <el-row v-if="!hidesNavbar">
      <el-col :span="24">
        <router-view name="navbar"></router-view>
      </el-col>
    </el-row>
    <el-row v-if="!hidesSidebar">
      <el-col :md="{ span: 6}" :lg="{ span: 4 }">
        <router-view name="sidebar"></router-view>
      </el-col>
      <el-col :md="{ span: 18}" :lg="{ span: 20 }">
        <router-view></router-view>
      </el-col>
    </el-row>
    <el-row v-if="hidesSidebar">
      <el-col>
        <router-view></router-view>
      </el-col>
    </el-row>
  </div>
</template>

<script>

export default {
  name: 'app',
  data: function () {
    return {
      hidesSidebar: false,
      hidesNavbar: false
    }
  },
  created () {
    this.updateBars()
  },
  watch: {
    '$route': 'updateBars'
  },
  methods: {
    updateBars () {
      this.hidesSidebar = !!this.$route.matched.some((r) => r.meta.hidesSidebar)
      this.hidesNavbar = !!this.$route.matched.some((r) => r.meta.hidesNavbar)
    }
  }
}
</script>

<style>
li.el-menu-item > a {
  text-decoration: none;
}
</style>
