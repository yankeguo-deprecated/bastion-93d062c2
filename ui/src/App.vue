<template>
  <div id="app">
    <el-row v-if="!hidesNavbar">
      <el-col :span="24">
        <router-view name="navbar"></router-view>
      </el-col>
    </el-row>
    <el-row v-if="!hidesSidebar">
      <el-col :md="{ span: 6}">
        <router-view name="sidebar"></router-view>
      </el-col>
      <el-col :md="{ span: 18}">
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

function shouldHidesSidebar ($route) {
  return !!$route.matched.some((r) => r.meta.hidesSidebar)
}

function shouldHidesNavbar ($route) {
  return !!$route.matched.some((r) => r.meta.hidesNavbar)
}

export default {
  name: 'app',
  data: function () {
    return {
      hidesSidebar: shouldHidesSidebar(this.$route),
      hidesNavbar: shouldHidesNavbar(this.$route)
    }
  },
  watch: {
    '$route': 'updateBars'
  },
  methods: {
    updateBars () {
      this.hidesSidebar = shouldHidesSidebar(this.$route)
      this.hidesNavbar = shouldHidesNavbar(this.$route)
    }
  }
}
</script>

<style>
li.el-menu-item > a {
  text-decoration: none;
}
</style>
