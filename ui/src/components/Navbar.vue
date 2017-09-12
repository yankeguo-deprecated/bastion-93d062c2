<template>
  <b-navbar toggleable="md" type="dark" variant="info">
    <b-nav-toggle target="nav_collapse"></b-nav-toggle>
    <router-link class="navbar-brand" :to="{ name: 'dashboard' }">Bastion</router-link>
    <b-collapse is-nav id="nav_collapse">
      <b-nav is-nav-bar>
        <b-nav-item :to="{name:'dashboard'}">仪表盘</b-nav-item>
      </b-nav>

      <b-nav is-nav-bar class="ml-auto">
        <b-nav-item-dropdown right>
          <template slot="button-content">
            <em>{{ currentUser.nickname }}</em>
          </template>
          <b-dropdown-item :to="{name:'profile'}">个人信息</b-dropdown-item>
          <b-dropdown-divider></b-dropdown-divider>
          <b-dropdown-item @click="signout">退出登录</b-dropdown-item>
        </b-nav-item-dropdown>
      </b-nav>
    </b-collapse>
  </b-navbar>
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
  },
  methods: {
    signout () {
      this.$api.destroyToken({ id: 'current' }).then(() => {
        this.$store.commit('setToken', null)
        this.$store.commit('setCurrentUser', null)
        this.$router.push({ name: 'index' })
      })
    }
  }
}
</script>
