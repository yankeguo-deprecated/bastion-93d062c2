<template>
  <b-container id="app">
    <!-- Navbar -->
    <b-row>
      <b-col>
        <router-view name="navbar"></router-view>
      </b-col>
    </b-row>

    <!-- Sidebar with Content -->
    <b-row>
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
    return {}
  },
  created () {
    this.fetchCurrentUserIfNeeded()
  },
  watch: {
    '$route': ['fetchCurrentUserIfNeeded']
  },
  methods: {
    fetchCurrentUserIfNeeded () {
      if (this.$store.getters.isSignedIn && !this.$store.getters.hasCurrentUser) {
        this.$api.fetchUser({id: 'current'}).then(({body}) => {
          const {user} = body
          this.$store.commit('setCurrentUser', user)
        })
      }
    }
  }
}
</script>

<style>
a {
  color: #17a2b8;
}
</style>
