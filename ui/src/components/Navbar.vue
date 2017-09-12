<template>
  <b-navbar toggleable="md" type="dark" variant="info">
    <b-nav-toggle target="nav_collapse"></b-nav-toggle>
    <router-link class="navbar-brand" :to="{ name: 'dashboard' }">Bastion</router-link>
    <b-collapse is-nav id="nav_collapse">
      <b-nav is-nav-bar>
        <b-nav-item :to="{name: 'dashboard'}">仪表盘</b-nav-item>
      </b-nav>

      <!-- Right aligned nav items -->
      <b-nav is-nav-bar class="ml-auto">
        <!--
        <b-nav-form>
          <b-form-input size="sm" class="mr-sm-2" type="text" placeholder="Search"/>
          <b-button size="sm" class="my-2 my-sm-0" type="submit">Search</b-button>
        </b-nav-form>
        <b-nav-item-dropdown text="Lang" right>
          <b-dropdown-item href="#">EN</b-dropdown-item>
          <b-dropdown-item href="#">ES</b-dropdown-item>
          <b-dropdown-item href="#">RU</b-dropdown-item>
          <b-dropdown-item href="#">FA</b-dropdown-item>
        </b-nav-item-dropdown>
      -->
        <b-nav-item-dropdown right>
          <!-- Using button-content slot -->
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
