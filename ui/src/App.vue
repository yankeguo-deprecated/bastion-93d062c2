<template>
  <b-container id="app">

    <!-- Navbar -->
    <b-row v-if="components.navbar">
      <b-col>
        <router-view name="navbar"></router-view>
      </b-col>
    </b-row>

    <!-- Sidebar with Content -->
    <b-row>
      <b-col v-if="components.sidebar" md="3">
        <router-view name="sidebar"></router-view>
      </b-col>
      <b-col>
        <router-view></router-view>
      </b-col>
    </b-row>
  </b-container>
</template>

<script>
export default {
  name: 'app',
  data: function () {
    return {
      components: {
        navbar: false,
        sidebar: false
      }
    }
  },
  created () {
    this.updateComponents()
    this.fetchCurrentUserIfNeeded()
  },
  watch: {
    '$route': ['updateComponents', 'fetchCurrentUserIfNeeded']
  },
  methods: {
    updateComponents () {
      this.components = {
        navbar: !this.$route.matched.some((r) => r.meta.hidesNavbar),
        sidebar: !this.$route.matched.some((r) => r.meta.hidesSidebar)
      }
    },
    fetchCurrentUserIfNeeded () {
      if (this.$store.getters.isSignedIn && !this.$store.getters.hasCurrentUser) {
        this.$api.fetchCurrentUser().then(({body}) => {
          const {user} = body
          this.$store.commit('setCurrentUser', user)
        })
      }
    }
  }
}
</script>
