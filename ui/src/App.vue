<template>
  <b-container fluid id="app">

    <!-- Navbar -->
    <b-row v-if="!hidesNavbar">
      <b-col>
        <router-view name="navbar"></router-view>
      </b-col>
    </b-row>

    <!-- Sidebar with Content -->
    <b-row>
      <b-col v-if="!hidesSidebar" md="2">
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
      hidesSidebar: false,
      hidesNavbar: false
    }
  },
  created () {
    this.updateBars()
    this.fetchCurrentUserIfNeeded()
  },
  watch: {
    '$route': ['updateBars', 'fetchCurrentUserIfNeeded']
  },
  methods: {
    updateBars () {
      this.hidesSidebar = !!this.$route.matched.some((r) => r.meta.hidesSidebar)
      this.hidesNavbar = !!this.$route.matched.some((r) => r.meta.hidesNavbar)
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
